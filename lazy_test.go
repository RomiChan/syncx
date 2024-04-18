// Copyright 2022 The syncx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncx

import (
	"testing"
	"time"
)

var one = Lazy[int]{
	Init: func() int {
		time.Sleep(time.Second) // simulate slow work
		return 1
	},
}

func run(t *testing.T, c chan bool) {
	if v := one.Get(); v != 1 {
		t.Errorf("lazy failed inside run: %d is not 1", v)
	}
	c <- true
}

func TestLazy(t *testing.T) {
	c := make(chan bool)
	const N = 20
	for i := 0; i < N; i++ {
		go run(t, c)
	}
	for i := 0; i < N; i++ {
		<-c
	}
	if one.val != 1 {
		t.Errorf("lazy failed outside run: %d is not 1", one.val)
	}
}

func BenchmarkLazy(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			one.Get()
		}
	})
}
