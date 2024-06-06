package main

import (
	linkedlist "structures/list/linkedList"
)

func main() {
	node := linkedlist.NewNode(2)

	println(*node.GetData())

	num := 3

	node.SetData(&num)

	println(*node.GetData())

	node2 := linkedlist.NewNode(4)

	node.SetNext(node2)

	println(*node.GetNext().GetData())

}
