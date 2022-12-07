package containers

type LinkedListWithPool struct {
	head      *node
	allocator pool
}

func NewLinkedListWithPool(size int) LinkedListWithPool {
	return LinkedListWithPool{
		allocator: newPool(size),
	}
}

func (l *LinkedListWithPool) push(value data) {
	n := l.allocator.get()
	n.value = value
	n.next = l.head
	l.head = n
}

func (l *LinkedListWithPool) pop() {
	if l.head != nil {
		n := l.head
		l.head = l.head.next
		l.allocator.put(n)
	}
}
