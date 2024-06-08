package nodotree

type Node[T any] struct {
	right   *Node[T]
	left    *Node[T]
	data    *T
	balance int
}

func (node *Node[T]) SetData(data *T) {
	node.data = data
}

func (node *Node[T]) GetData() *T {
	return node.data
}

func (node *Node[T]) SetLeft(n *Node[T]) {
	node.left = n
}

func (node *Node[T]) SetRight(n *Node[T]) {
	node.right = n
}

func (node *Node[T]) GetLeft() *Node[T] {
	return node.left
}

func (node *Node[T]) GetRight() *Node[T] {
	return node.right
}

func (node *Node[T]) SetBalance(n int) {
	node.balance = n
}

func (node *Node[T]) GetBalance() int {
	return node.balance
}

func New[T any](data *T) *Node[T] {
	newNode := Node[T]{
		right: nil,
		left:  nil,
		data:  data,
	}

	return &newNode
}
