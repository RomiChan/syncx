// Copyright 2022 The syncx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncx

import (
	"sync/atomic"
	"unsafe"
)

// A Pointer provides an atomic load and store of a pointer value.
// The zero value for a Pointer returns nil from Load.
// Once Store has been called, a Pointer must not be copied.
type Pointer[T any] struct {
	value unsafe.Pointer // *T
}

// Load returns the value set by the most recent Store.
func (v *Pointer[T]) Load() (val *T) {
	value := atomic.LoadPointer(&v.value)
	return (*T)(value)
}

// LoadValue returns dereferenced value set by the most recent Store.
// If v is zero value, ok is false.
func (v *Pointer[T]) LoadValue() (val T, ok bool) {
	ptr := v.Load()
	if ptr == nil {
		return
	}
	return *ptr, true
}

// Store sets the pointer value to v.
func (v *Pointer[T]) Store(val *T) {
	atomic.StorePointer(&v.value, unsafe.Pointer(val))
}

// StoreValue sets the value to Pointer v.
func (v *Pointer[T]) StoreValue(val T) {
	atomic.StorePointer(&v.value, unsafe.Pointer(&val))
}

// Swap stores new into Pointer and returns the previous value. It returns nil if
// the Pointer is empty.
func (v *Pointer[T]) Swap(new *T) (old *T) {
	return (*T)(atomic.SwapPointer(&v.value, unsafe.Pointer(new)))
}

// CompareAndSwap executes the compare-and-swap operation for the Pointer.
func (v *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return atomic.CompareAndSwapPointer(&v.value, unsafe.Pointer(old), unsafe.Pointer(new))
}
