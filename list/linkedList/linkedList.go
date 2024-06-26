package linkedlist

import "errors"

type LinkedList[T any] struct {
	head    *Node[T]
	length  int
	equeals func(a, b T) bool
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) Add(data *T) {
	node := NewNode(*data)

	if list.head == nil {
		list.head = node
		list.length = list.length + 1
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

func (list *LinkedList[T]) Contains(data *T) bool {
	currentNode := list.head

	for currentNode != nil {
		if list.equeals(*currentNode.GetData(), *data) {
			return true
		}

		currentNode = currentNode.GetNext()
	}

	return false
}

func (list *LinkedList[T]) IndexOf(data *T) int {
	currentNode := list.head

	for i := 0; i < list.length && currentNode != nil; i++ {
		if list.equeals(*currentNode.GetData(), *data) {
			return i
		}
		currentNode = currentNode.GetNext()
	}
	return -1
}

func (list *LinkedList[T]) Set(n int, data *T) (*T, error) {
	if n < 0 || n > list.length {
		return nil, errors.New("index out of limits")
	}

	node := NewNode[T](*data)

	currentNode := list.head

	if n == 0 {
		node.SetNext(currentNode.GetNext())

		list.head = node
		return currentNode.GetData(), nil
	}

	count := 0
	for {
		if count+1 == n {
			saveNode := currentNode.GetNext()
			node.SetNext(currentNode.GetNext().GetNext())
			currentNode.SetNext(node)

			return saveNode.GetData(), nil
		}
		count = count + 1
	}
}

func NewLinkedList[T any](equals func(a, b T) bool) *LinkedList[T] {
	list := LinkedList[T]{
		head:    nil,
		length:  0,
		equeals: equals,
	}

	return &list
}
