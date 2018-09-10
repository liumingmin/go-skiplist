package skiplist

import (
	"testing"
	"unsafe"
)

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 1)
		newNode(nil, 2)
		newNode(nil, 2)
		newNode(nil, 2)
		newNode(nil, 2)
		newNode(nil, 3)
		newNode(nil, 3)
		newNode(nil, 4)
	}
}

func BenchmarkAllocFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 1)
		newNodeFast(nil, 2)
		newNodeFast(nil, 2)
		newNodeFast(nil, 2)
		newNodeFast(nil, 2)
		newNodeFast(nil, 3)
		newNodeFast(nil, 3)
		newNodeFast(nil, 4)
	}
}

func BenchmarkAllocUnsafeOffsetof(b *testing.B) {
	var a uintptr
	for i := 0; i < b.N; i++ {
		a += unsafe.Offsetof(l1n{}.array)
	}
	b.StopTimer()
	println(a)
}