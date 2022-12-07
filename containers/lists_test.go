package containers

import (
	"testing"
)

func BenchmarkLinkedList(b *testing.B) {
	list := NewLinkedList()
	for i := 0; i < b.N; i++ {
		list.push(data{})
		for j := i; j >= 0; j-- {
			list.pop()
		}
		for j := i; j >= 0; j-- {
			list.push(data{})
		}
	}
}

func BenchmarkLinkedWithSyncPool(b *testing.B) {
	list := NewLinkedListWithSyncPool()
	for i := 0; i < b.N; i++ {
		list.push(data{})
		for j := i; j >= 0; j-- {
			list.pop()
		}
		for j := i; j >= 0; j-- {
			list.push(data{})
		}
	}
}

func BenchmarkLinkedWithFreeList(b *testing.B) {
	list := NewLinkedListWithFreeList()
	for i := 0; i < b.N; i++ {
		list.push(data{})
		for j := i; j >= 0; j-- {
			list.pop()
		}
		for j := i; j >= 0; j-- {
			list.push(data{})
		}
	}
}
