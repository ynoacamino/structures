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

func (trie Trie[T]) Get(str string) *T {
	bytes := convertToBytes(str)

	return trie.root.GetInNode(&bytes, 0)
}

func (trie Trie[T]) Add(str string, value *T) {
	bytes := convertToBytes(str)

	trie.root.AddInNode(&bytes, 0, value)
}

func (trie Trie[T]) Search(str string) bool {
	bytes := convertToBytes(str)

	return trie.root.SearchInNode(&bytes, 0)
}

func convertToBytes(str string) []byte {
	bytes := []byte(strings.ToLower(str))

	for i, b := range bytes {
		bytes[i] = b - 97
	}

	return bytes
}

func (trie *Trie[T]) Remove(str string) *T {
	bytes := convertToBytes(str)

	return trie.root.RemoveInNode(&bytes, 0)
}

func (trie *Trie[T]) suggest(str string) []*n.TrieNode[T] {
	bytes := convertToBytes(str)

	node := trie.root.SearchPreFix(&bytes, 0)

	if node == nil {
		return nil
	}

	suggestSlice := make([]*n.TrieNode[T], 0) 

	return suggestSlice
}
