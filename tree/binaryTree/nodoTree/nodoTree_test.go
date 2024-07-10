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

func TestGetHeightRightMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node.GetHeightRight() != 0 {
		t.Fatal("Height right must be 0 but is ", node.GetHeightRight())
	}
}

func TestSetHeightRightMethod(t *testing.T) {
	data := 10
	node := New(&data)

	node.SetHeightRight(10)

	if node.GetHeightRight() != 10 {
		t.Fatal("Height right must be 10, but is ", node.GetHeightRight())
	}
}

func TestGetHeightLeftMethod(t *testing.T) {
	data := 10
	node := New(&data)

	if node.GetHeightLeft() != 0 {
		t.Fatal("Height left must be 0, but is ", node.GetHeightLeft())
	}
}

func TestSetHeightLeftMethod(t *testing.T) {
	data := 10
	node := New(&data)

	node.SetHeightLeft(10)

	if node.GetHeightLeft() != 10 {
		t.Fatal("Height left must be 10, but is ", node.GetHeightLeft())
	}
}
