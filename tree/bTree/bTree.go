package btree

import (
	"math"
	n "structures/tree/bTree/nodeBTree"
)

type BTree[T any] struct {
	root *n.NodeBTree[T]
	t    int

	up bool

	ComparteTo func(a, b *T) int
}

func NewBTree[T any](t int, comparteTo func(a, b *T) int) *BTree[T] {
	return &BTree[T]{
		root:       nil,
		t:          t,
		ComparteTo: comparteTo,
	}
}

func (bTree *BTree[T]) IsEmpty() bool {
	return bTree.root == nil
}

func (bTree *BTree[T]) Search(data *T) *T {
	_, res := bTree.root.Search(data)

	return res
}

func (bTree *BTree[T]) Insert(node *n.NodeBTree[T], value *T) *n.NodeBTree[T] {
	if node.IsSheet() {
		bTree.insertInOrder(node, value, nil)

	} else {
		index := bTree.searchChild(node, value)

		new := bTree.Insert(node.Childs[index], value)

		if new != nil {
			pivote := new.Values[0]

			bTree.insertInOrder(node, pivote, new)
			// quitar la clave 0 de nuevo
		}
	}

	if node.IsInvalid() {
		return bTree.split(node)
	} else {
		return nil
	}

}

func (bTree *BTree[T]) insertInOrder(node *n.NodeBTree[T], value *T, rigthChild *n.NodeBTree[T]) {
	pivote := 0

	for node.ComparteTo(node.Values[pivote], value) > 0 {
		pivote = pivote + 1
	}

	for i := node.NumValues - 1; i >= pivote; i-- {
		node.Values[i+1] = node.Values[i]
	}

	for i := node.NumChilds - 1; i > pivote; i-- {
		node.Childs[i+1] = node.Childs[i]
	}

	node.Values[pivote] = value
	node.Childs[pivote+1] = rigthChild

	node.NumValues = node.NumValues + 1

	if !node.IsSheet() {
		node.NumChilds = node.NumChilds + 1
	}
}

func (bTree *BTree[T]) searchChild(node *n.NodeBTree[T], value *T) int {
	pibote := 0

	for pibote, currentValue := range node.Values {
		if currentValue == nil {
			return pibote
		}

		if bTree.ComparteTo(currentValue, value) > 0 {
			return pibote
		}
	}

	return pibote + 1
}

func (bTree BTree[T]) split(node *n.NodeBTree[T]) *n.NodeBTree[T] {
	newNode := n.NewNodeBTree(node.Grade, node.ComparteTo)

	l := int(math.Ceil(float64(node.Grade)/2 - 1))

	i := l
	j := 0
	for {
		if i >= node.GetMaxValues() {
			break
		}

		newNode.Values[j] = node.Values[i]
		node.Values[i] = nil
		// creo que tenemos que reducir el numero interno de valores

		if !node.IsSheet() {
			newNode.Childs[j+1] = node.Childs[i+1]
			node.Childs[i+1] = nil
			// creo que aqui se modifica el numero interno de hijos
		}

		i = i + 1
		j = j + 1
	}

	return newNode

}
