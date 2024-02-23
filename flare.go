package flare

import (
	"context"
	"sync"
)

type (
	// Notifier interface defines the methods for signaling and waiting on notifications.
	Notifier interface {
		// Signal triggers the notifier, indicating an event has occurred.
		Signal()

		// Hold returns a channel that is closed when Signal is called, allowing callers to wait for a signal.
		Hold() <-chan struct{}
	}

	notifier struct {
		ch   chan struct{}
		once sync.Once
	}
)

func NewNotifier() Notifier {
	return &notifier{ch: make(chan struct{})}
}
func NewNotifierWithCancel(parent context.Context) (Notifier, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	n := NewNotifier()

	go func() {
		<-ctx.Done()
		n.Signal()
	}()

	return n, cancel
}

func (n *notifier) Signal() {
	n.once.Do(func() {
		close(n.ch)
	})
}
func (n *notifier) Hold() <-chan struct{} {
	return n.ch
}
