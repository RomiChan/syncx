// Copyright 2022 The syncx Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncx

import (
	"testing"
)

func TestPointer(t *testing.T) {
	var v Pointer[int]
	if v.Load() != nil {
		t.Fatal("initial Value is not nil")
	}
	v.StoreValue(42)
	if xx, ok := v.LoadValue(); !ok || xx != 42 {
		t.Fatalf("wrong value: got %+v, want 42", xx)
	}
	v.StoreValue(84)
	if xx, ok := v.LoadValue(); !ok || xx != 84 {
		t.Fatalf("wrong value: got %+v, want 84", xx)
	}
}
