package nodotree

import (
	"testing"
)

func TestNewMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node == nil {
		t.Fatal("Node is nil")
	}
}

func TestGetDataMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if *node.GetData() != 10 {
		t.Fatal("Data must be 10, but got ", *node.GetData())
	}
}

func TestSetDataMehod(t *testing.T) {
	data := 10
	node := New(&data)

	newData := 20
	node.SetData(&newData)

	if *node.GetData() != 20 {
		t.Fatal("Data must be 20, but got ", *node.GetData())
	}
}

func TestGetLeftMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node.GetLeft() != nil {
		t.Fatal("Node left must be nil")
	}
}

func TestSetLeftMethod(t *testing.T) {
	data := 10
	node := New(&data)

	left := New(&data)
	node.SetLeft(left)

	if node.GetLeft() == nil {
		t.Fatal("Node left must not be nil")
	}
}

func TestGetRightMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node.GetRight() != nil {
		t.Fatal("Node right must be nil")
	}
}

func TestSetRightMethod(t *testing.T) {
	data := 10
	node := New(&data)

	right := New(&data)
	node.SetRight(right)

	if node.GetRight() == nil {
		t.Fatal("Node right must not be nil")
	}
}

func TestGetBalanceMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node.GetBalance() != 0 {
		t.Fatal("Node balance must be 0")
	}
}

func TestSetBalanceMethod(t *testing.T) {
	data := 10
	node := New(&data)

	node.SetBalance(1)

	if node.GetBalance() != 1 {
		t.Fatal("Node balance must be 1")
	}
}
