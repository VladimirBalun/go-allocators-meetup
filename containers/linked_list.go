package containers

type LinkedList struct {
	head *node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) push(value data) {
	l.head = &node{value: value, next: l.head}
}

func (l *LinkedList) pop() {
	if l.head != nil {
		l.head = l.head.next
	}
}
