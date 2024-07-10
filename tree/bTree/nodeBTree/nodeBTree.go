package nodebtree

type NodeBTree[T any] struct {
	Values []*T
	Childs []*NodeBTree[T]

	NumValues int
	NumChilds int

	Grade int

	ComparteTo func(a, b *T) int
}

func NewNodeBTree[T any](grade int, comparteTo func(a, b *T) int) *NodeBTree[T] {
	return &NodeBTree[T]{
		Values:     make([]*T, grade),
		Childs:     make([]*NodeBTree[T], grade),
		NumValues:  0,
		NumChilds:  0,
		Grade:      grade,
		ComparteTo: comparteTo,
	}
}

func (n *NodeBTree[T]) Search(data *T) (*NodeBTree[T], *T) {
	var i int

	for i = 0; i < n.GetMaxValues(); i++ {
		slot := n.Values[i]

		if slot == nil {
			break
		}

		if n.ComparteTo(slot, data) == 0 {
			return n, slot
		}

		if n.ComparteTo(slot, data) > 0 {
			n := n.Childs[i]

			if n == nil {
				return nil, nil
			}

			return n.Search(data)
		}
	}

	return n.Childs[i].Search(data)
}

// el que sube es t - 1 [0, 1, 2, 3]
func (n *NodeBTree[T]) SplitChild(currentNode *NodeBTree[T]) *NodeBTree[T] {
	return nil
}

func (n *NodeBTree[T]) GetMaxValues() int {
	return n.Grade - 1
}

func (n *NodeBTree[T]) GetMaxChilds() int {
	return n.Grade
}

func (n *NodeBTree[T]) IsInvalid() bool {
	return n.NumValues == n.Grade
}

func (n *NodeBTree[T]) AddKey(data *T) *T {
	var i int

	for i = 0; i < n.GetMaxChilds(); i++ {
		slot := n.Values[i]

		if slot == nil {
			break
		}

		if n.ComparteTo(data, slot) < 1 { // revisar esto
			break
		}

		i = i + 1
	}

	n.Values = n.push(n.Values, data, i)

	return data
}

func (n *NodeBTree[T]) push(s []*T, data *T, index int) []*T {
	var sCopy []*T = s

	for i := index; i < n.GetMaxValues(); i++ {
		sCopy[i+1] = s[i]
	}

	sCopy[index] = data

	return sCopy
}

func (n *NodeBTree[T]) IsSheet() bool {
	return n.NumChilds == 0
}
