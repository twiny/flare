package flare

import (
	"context"
	"testing"
	"time"
)

func TestNewNotifier(t *testing.T) {
	n := NewNotifier()
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	select {
	case <-n.Hold():
		t.Fatal("Notifier should not be signaled yet")
	default:
	}

	n.Signal()

	select {
	case <-n.Hold():
	default:
		t.Fatal("Notifier should be signaled")
	}
}
func TestNewNotifierWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure context cancel function is called to avoid leaking resources

	n, _ := NewNotifierWithCancel(ctx)
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	select {
	case <-n.Hold():
		t.Fatal("Notifier should not be signaled yet")
	default:
	}

	cancel()

	// Wait for the notifier to be signaled, with a timeout to prevent hanging indefinitely.
	select {
	case <-n.Hold():
	case <-time.After(1 * time.Second):
		t.Fatal("Notifier should be signaled after context cancel")
	}
}
func TestNewNotifierWithCancel_SignalBeforeContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n, _ := NewNotifierWithCancel(ctx)
	if n == nil {
		t.Fatal("expected Notifier, got nil")
	}

	n.Signal()

	select {
	case <-n.Hold():
	default:
		t.Fatal("Notifier should be signaled after Signal call")
	}

	cancel()

	// After canceling the context, the notifier should still be in the signaled state.
	select {
	case <-n.Hold():
	default:
		t.Fatal("Notifier should remain signaled after context cancel")
	}
}
