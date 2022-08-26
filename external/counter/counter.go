package counter

import (
	"strconv"
	"sync/atomic"
)

// Counter просто потокобезопасный счетчик
type Counter struct {
	val uint64
}

func (c *Counter) Inc() uint64 {
	return atomic.AddUint64(&c.val, 1)
}

func (c *Counter) Get() uint64 {
	return atomic.LoadUint64(&c.val)
}

func (c *Counter) String() string {
	return strconv.FormatUint(c.Get(), 10)
}
