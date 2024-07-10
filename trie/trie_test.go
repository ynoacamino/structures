package trie

import "testing"

func TestMain(t *testing.T) {
	trie := NewTrie[int]()

	e1 := 1

	trie.Add("holaed", &e1)
	println("----")
	trie.Add("hola", &e1)
	println("----")
	trie.Add("hol", &e1)
}
