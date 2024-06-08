package binarytree

import (
	"errors"
	"fmt"
	nodotree "structures/tree/nodoTree"
)

type BinaryTree[T any] struct {
	root      *nodotree.Node[T]
	height    int
	length    int
	compareTo func(a, b T) int
}

func (tree *BinaryTree[T]) Size() int {
	return tree.length
}

func (tree *BinaryTree[T]) Add(data *T) {
	newNode := nodotree.New(data)

	if tree.root == nil {
		tree.root = newNode
		tree.length = tree.length + 1
		return
	}

	tree.height = tree.addNode(newNode, tree.root)

	tree.length = tree.length + 1

	tree.searchUnbalance()
}

func (tree *BinaryTree[T]) addNode(newNode *nodotree.Node[T], node *nodotree.Node[T]) int {
	var isRight bool = tree.compareTo(*newNode.GetData(), *node.GetData()) > 0

	if isRight {
		if node.GetRight() != nil {
			newHeigth := tree.addNode(newNode, node.GetRight())

			node.SetHeightRight(newHeigth + 1)
		} else {
			node.SetRight(newNode)
			node.SetHeightRight(1)
		}
	} else {
		if node.GetLeft() != nil {
			newHeigth := tree.addNode(newNode, node.GetLeft())

			node.SetHeightLeft(newHeigth + 1)
		} else {
			node.SetLeft(newNode)
			node.SetHeightLeft(1)
		}
	}

	return node.GetMaxHeight()
}

func (tree *BinaryTree[T]) searchUnbalance() {
	if tree.length < 3 {
		return
	}

	currentNode := tree.root

	if tree.root.GetBalance() == 2 || tree.root.GetBalance() == -2 {
		tree.root = balancer(tree.root)
		return
	}

	for currentNode.GetBalance() > 1 || currentNode.GetBalance() < -1 {
		fmt.Println("currentNode", *currentNode.GetData(), "balance:", currentNode.GetBalance())

		if currentNode.GetRight() != nil && (currentNode.GetRight().GetBalance() == 2 || currentNode.GetRight().GetBalance() == -2) {
			currentNode.SetRight(balancer(currentNode.GetRight()))
			return
		}

		if currentNode.GetLeft() != nil && (currentNode.GetLeft().GetBalance() == 2 || currentNode.GetLeft().GetBalance() == -2) {
			currentNode.SetLeft(balancer(currentNode.GetLeft()))
			return
		}

		if currentNode.GetBalance() > 1 {
			currentNode = currentNode.GetRight()
		} else {
			currentNode = currentNode.GetLeft()
		}
	}

}

func balancer[T any](node *nodotree.Node[T]) *nodotree.Node[T] {
	var firstBalance int = node.GetBalance()
	var secondBalance int

	if firstBalance > 0 {
		secondBalance = node.GetRight().GetBalance()
	} else {
		secondBalance = node.GetLeft().GetBalance()
	}

	var newHead *nodotree.Node[T]

	if firstBalance == 2 {
		if secondBalance == 1 {
			newHead = rotationSL(node)
		}
		if secondBalance == -1 {
			newHead = rotationDL(node)
		}
	}
	if firstBalance == -2 {
		if secondBalance == -1 {
			newHead = rotationSR(node)
		}
		if secondBalance == 1 {

			newHead = rotationDR(node)
		}
	}

	return newHead
}

// rotacion simple izquierda porque jalas a la izquierda un nodo 2 1
func rotationSL[T any](node *nodotree.Node[T]) *nodotree.Node[T] {
	newHead := node.GetRight()

	node.SetRight(newHead.GetLeft())
	node.SetHeightRight(newHead.GetHeightLeft())

	newHead.SetLeft(node)
	newHead.SetHeightLeft(node.GetMaxHeight() + 1)

	return newHead
}

// -2 -1
func rotationSR[T any](node *nodotree.Node[T]) *nodotree.Node[T] {
	newHead := node.GetLeft()

	node.SetLeft(newHead.GetRight())
	node.SetHeightLeft(newHead.GetHeightRight())

	newHead.SetRight(node)
	newHead.SetHeightRight(node.GetMaxHeight() + 1)

	return newHead
}

// 2 -1
func rotationDL[T any](node *nodotree.Node[T]) *nodotree.Node[T] {
	newHead := node.GetRight().GetLeft()

	node.GetRight().SetLeft(newHead.GetRight())
	node.GetRight().SetHeightLeft(newHead.GetHeightRight())

	newHead.SetRight(node.GetRight())
	newHead.SetHeightRight(node.GetRight().GetMaxHeight() + 1)

	node.SetRight(newHead.GetLeft())
	node.SetHeightRight(newHead.GetHeightLeft())

	newHead.SetLeft(node)
	newHead.SetHeightLeft(node.GetMaxHeight() + 1)

	return newHead
}

func rotationDR[T any](node *nodotree.Node[T]) *nodotree.Node[T] { // <--------------- AQUI
	newHead := node.GetLeft().GetRight()

	node.GetLeft().SetRight(newHead.GetLeft())
	node.GetLeft().SetHeightRight(newHead.GetHeightLeft())

	newHead.SetLeft(node.GetLeft())
	newHead.SetHeightLeft(node.GetLeft().GetMaxHeight() + 1)

	node.SetLeft(newHead.GetRight())
	node.SetHeightLeft(newHead.GetHeightRight())

	newHead.SetRight(node)
	newHead.SetHeightRight(node.GetMaxHeight() + 1)

	return newHead
}

func (tree *BinaryTree[T]) Get(data *T) (*T, error) {
	currentNode := tree.root

	for {
		if tree.compareTo(*currentNode.GetData(), *data) == 0 {
			return currentNode.GetData(), nil
		}

		if !currentNode.IsFather() {
			break
		}

		if tree.compareTo(*currentNode.GetData(), *data) > 0 {
			currentNode = currentNode.GetLeft()
		} else {
			currentNode = currentNode.GetRight()
		}
	}

	return nil, errors.New("element not found")
}

func New[T any](compareTo func(a, b T) int) *BinaryTree[T] {
	return &BinaryTree[T]{
		root:      nil,
		height:    0,
		length:    0,
		compareTo: compareTo,
	}
}

func (tree *BinaryTree[T]) print() {
	if tree.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	tree.printNode(tree.root, "", true)
}

func (tree *BinaryTree[T]) printNode(node *nodotree.Node[T], prefix string, isTail bool) {
	if node == nil {
		return
	}

	if node.GetRight() != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│     "
		} else {
			newPrefix += "      "
		}
		tree.printNode(node.GetRight(), newPrefix, false)
	}

	fmt.Printf("%s", prefix)
	if isTail {
		fmt.Printf("└──── ")
	} else {
		fmt.Printf("┌──── ")
	}
	fmt.Println(node.GetHeightLeft(), *node.GetData(), node.GetHeightRight())

	if node.GetLeft() != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "      "
		} else {
			newPrefix += "│     "
		}
		tree.printNode(node.GetLeft(), newPrefix, true)
	}
}
