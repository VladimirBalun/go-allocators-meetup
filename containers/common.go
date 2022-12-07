package containers

// ------------------------------
//         Data
// ------------------------------

type data struct {
	id      int64
	name    string
	surname string
	address string
	status  string
	active  bool
}

// ------------------------------
//         Node
// ------------------------------

type node struct {
	value data
	next  *node
}

// ------------------------------
//         Free List
// ------------------------------

type freeList struct {
	head *node
}

func (l *freeList) put(n *node) {
	n.next = l.head
	l.head = n
}

func (l *freeList) get() *node {
	if l.head == nil {
		return &node{}
	} else {
		n := l.head
		l.head = l.head.next
		return n
	}
}

// ------------------------------
//              Pool
// ------------------------------

type pool struct {
	nodes []node
	head  *node
}

func newPool(size int) pool {
	var head *node
	nodes := make([]node, size)
	for i := 0; i < len(nodes); i++ {
		nodes[i].next = head
		head = &nodes[i]
	}

	return pool{
		nodes: nodes,
		head:  head,
	}
}

func (l *pool) put(n *node) {
	n.next = l.head
	l.head = n
}

func (l *pool) get() *node {
	if l.head != nil {
		n := l.head
		l.head = l.head.next
		return n
	} else {
		return nil
	}
}
