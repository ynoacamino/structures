package trie

import (
	"strings"
	n "structures/trie/trieNode"
)

type Trie[T any] struct {
	root *n.TrieNode[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		root: n.NewTrieNode[T](false, nil, 0),
	}
}

func (trie Trie[T]) Add(str string, value *T) {
	bytes := convertToBytes(str)

	trie.root.AddInNode(&bytes, 0, value)
}

func convertToBytes(str string) []byte {
	bytes := []byte(strings.ToLower(str))

	for i, b := range bytes {
		bytes[i] = b - 97
	}

	return bytes
}
