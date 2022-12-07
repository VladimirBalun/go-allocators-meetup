package containers

import "sync"

type LinkedListWithSyncPool struct {
	head *node
	pool sync.Pool
}

func NewLinkedListWithSyncPool() LinkedListWithSyncPool {
	return LinkedListWithSyncPool{
		pool: sync.Pool{New: func() interface{} { return &node{} }},
	}
}

func (l *LinkedListWithSyncPool) push(value data) {
	n := l.pool.Get().(*node)
	n.value = value
	n.next = l.head
	l.head = n
}

func (l *LinkedListWithSyncPool) pop() {
	if l.head != nil {
		n := l.head
		l.head = l.head.next
		l.pool.Put(n)
	}
}
