package containers

type LinkedListWithFreeList struct {
	head      *node
	allocator freeList
}

func NewLinkedListWithFreeList() LinkedListWithFreeList {
	return LinkedListWithFreeList{}
}

func (l *LinkedListWithFreeList) push(value data) {
	n := l.allocator.get()
	n.value = value
	n.next = l.head
	l.head = n
}

func (l *LinkedListWithFreeList) pop() {
	if l.head != nil {
		n := l.head
		l.head = l.head.next
		l.allocator.put(n)
	}
}
