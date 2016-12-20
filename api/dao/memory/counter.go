package memory

import "sync/atomic"

type count64 uint64

func (c *count64) increment() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *count64) get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}