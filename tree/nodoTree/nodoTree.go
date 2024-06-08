package nodotree

type Node[T any] struct {
	right       *Node[T]
	left        *Node[T]
	data        *T
	heightRight int
	heightLeft  int
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

func (node *Node[T]) SetHeightRight(height int) {
	node.heightRight = height
}

func (node *Node[T]) SetHeightLeft(height int) {
	node.heightLeft = height
}

func (node *Node[T]) GetHeightRight() int {
	return node.heightRight
}

func (node *Node[T]) GetHeightLeft() int {
	return node.heightLeft
}

func New[T any](data *T) *Node[T] {
	newNode := Node[T]{
		right: nil,
		left:  nil,
		data:  data,
	}

	return &newNode
}
