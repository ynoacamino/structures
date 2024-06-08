package linkedlist

type LinkedList[T any] struct {
	head    *Node[T]
	length  int
	equeals func(a, b T) bool
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) Add(data T) {
	node := NewNode(data)

	if list.head == nil {
		list.head = node
		return
	}

	currentNode := list.head

	for currentNode.GetNext() != nil {
		currentNode = currentNode.GetNext()
	}

	currentNode.SetNext(node)

	list.length = list.length + 1
}

func (list *LinkedList[T]) Get(n int) *T {
	if n < 0 || n > list.length {
		return nil
	}

	currentNode := list.head

	for i := 0; i < n; i++ {
		currentNode = currentNode.GetNext()
	}

	return currentNode.GetData()
}

func (list *LinkedList[T]) Contains(data T) bool {
	currentNode := list.head

	for currentNode != nil {
		if list.equeals(*currentNode.GetData(), data) {
			return true
		}

		currentNode = currentNode.GetNext()
	}

	return false
}

func NewLinkedList[T any](equals func(a, b T) bool) *LinkedList[T] {
	list := LinkedList[T]{
		head:    nil,
		length:  0,
		equeals: equals,
	}

	return &list
}
