// Copyright 2022 The syncx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncx

import (
	"runtime"
	"sync/atomic"
)

type Lazy[T any] struct {
	done   atomic.Uint32
	locked atomic.Uint32
	val    T

	Init func() T
}

func (l *Lazy[T]) Get() T {
	if l.done.Load() == 1 {
		return l.val
	}
	// Slow path (outlined so that the fast path can be inlined)
	return l.slowGet()
}

func (l *Lazy[T]) slowGet() T {
	// use a spinlock
	for !l.locked.CompareAndSwap(0, 1) {
		runtime.Gosched()
	}
	defer l.locked.Store(0)

	if l.done.Load() == 0 {
		defer l.done.Store(1)
		l.val = l.Init()
	}
	return l.val
}
