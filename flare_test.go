package flare

import (
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	n := New()
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	select {
	case <-n.Done():
		t.Fatal("Notifier should not be done yet")
	default:
	}

	n.Cancel()

	select {
	case <-n.Done():
	default:
		t.Fatal("Notifier should be done")
	}
}

func TestNewWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Just to be safe

	n := NewWithContext(ctx)
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	select {
	case <-n.Done():
		t.Fatal("Notifier should not be done yet")
	default:
	}

	cancel()

	time.Sleep(50 * time.Millisecond) // Giving some time for the goroutine to trigger Cancel

	select {
	case <-n.Done():
	default:
		t.Fatal("Notifier should be done")
	}
}

func TestDoubleCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	n := NewWithContext(ctx)
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	select {
	case <-n.Done():
		t.Fatal("Notifier should not be done yet")
	default:
	}

	n.Cancel()

	select {
	case <-n.Done():
	default:
		t.Fatal("Notifier should be done after manual cancel")
	}

	cancel() // This should not cause any panic or issue, as the notifier's Cancel should be idempotent

	time.Sleep(100 * time.Millisecond) // Giving some time to ensure any potential panic would have occurred
}
