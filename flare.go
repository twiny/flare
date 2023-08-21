package flare

import (
	"context"
	"sync"
)

type (
	Notifier interface {
		Cancel()
		Done() <-chan struct{}
	}

	notifier struct {
		ch   chan struct{}
		once sync.Once
	}
)

func New() Notifier {
	return &notifier{ch: make(chan struct{})}
}
func NewWithContext(ctx context.Context) Notifier {
	n := New()
	go func() {
		<-ctx.Done()
		n.Cancel()
	}()
	return n
}

func (n *notifier) Cancel() {
	n.once.Do(func() {
		close(n.ch)
	})
}

func (n *notifier) Done() <-chan struct{} {
	return n.ch
}
