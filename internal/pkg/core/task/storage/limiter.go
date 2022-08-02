package storage

import (
	"context"
	"time"
)

type workerLimiter struct {
	poolCh  chan struct{}
	timeout time.Duration
}

func newWorkerLimiter(poolSize uint, timeout time.Duration) workerLimiter {
	return workerLimiter{
		poolCh:  make(chan struct{}, poolSize),
		timeout: timeout,
	}
}

func (l *workerLimiter) start(ctx context.Context) error {
	ctx, cl := context.WithTimeout(ctx, l.timeout)
	defer cl()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case l.poolCh <- struct{}{}:
		return nil
	}
}

func (l *workerLimiter) end() {
	<-l.poolCh
}
