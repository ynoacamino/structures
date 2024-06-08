package binarytree

import "testing"

type User struct {
	id   int
	data int
}

func compareTo(a, b User) int {
	if a.id < b.id {
		return -1
	}
	if a.id == b.id {
		return 0
	}
	return 1
}

func TestNewMethod(t *testing.T) {
	tree := New(compareTo)

	if tree == nil {
		t.Fatal("The tree struct is null")
	}
}

func TestAddMethod(t *testing.T) {
	tree := New(compareTo)

	e1 := User{
		id:   10,
		data: 11,
	}

	tree.Add(&e1)

	if tree.Size() != 1 {
		t.Fatal("Error in insert root, the size must be 1, but is ", tree.Size())
	}

	e2 := User{
		id:   15,
		data: 16,
	}

	e3 := User{
		id:   20,
		data: 21,
	}

	e4 := User{
		id:   5,
		data: 6,
	}

	tree.Add(&e2)
	tree.Add(&e3)
	tree.Add(&e4)

	if tree.Size() != 4 {
		t.Fatal("Error in insert new node, the size must be 4, but is ", tree.Size())
	}
}

func TestGetMethod(t *testing.T) {
	tree := New(compareTo)

	e1 := User{
		id:   10,
		data: 11,
	}
	tree.Add(&e1)

	get1, err := tree.Get(&User{id: 10})

	if err != nil {
		t.Fatal(err)
	}

	if get1.data != e1.data {
		t.Fatal("Error in search correct node in root")
	}

	e2 := User{
		id:   15,
		data: 16,
	}

	e3 := User{
		id:   20,
		data: 21,
	}

	e4 := User{
		id:   5,
		data: 6,
	}

	tree.Add(&e2)
	tree.Add(&e3)
	tree.Add(&e4)

	get2, err := tree.Get(&User{id: 15})

	if err != nil {
		t.Fatal(err)
	}

	if get2.data != e2.data {
		t.Fatal("Error in search correct node in branchs")
	}

	tree.print()
}

func TestTest(t *testing.T) {
	tree := New(compareTo)

	e1 := User{
		id:   10,
		data: 11,
	}

	e2 := User{
		id:   15,
		data: 16,
	}

	e3 := User{
		id:   20,
		data: 21,
	}

	e4 := User{
		id:   25,
		data: 26,
	}

	tree.Add(&e1)
	tree.Add(&e2)

	tree.print()

	tree.Add(&e3)
	tree.Add(&e4)

	tree.print()
}
