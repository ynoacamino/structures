package trienode

import "fmt"

type TrieNode[T any] struct {
	key   byte
	value *T

	childs []*TrieNode[T]
	end    bool
}

func NewTrieNode[T any](end bool, value *T, key byte) *TrieNode[T] {
	return &TrieNode[T]{
		childs: make([]*TrieNode[T], 27),
		end:    end,
		value:  value,
		key:    key,
	}
}

func (node *TrieNode[T]) AddInNode(bytes *[]byte, i int, value *T) {
	key := (*bytes)[i]
	isFinal := len(*bytes) == i+1

	if node.childs[key] == nil {
		fmt.Println(" - No existe")
		if isFinal {
			node.childs[key] = NewTrieNode[T](true, value, key)
		} else {
			fmt.Println(" - No es final")
			node.childs[key] = NewTrieNode[T](false, nil, key)
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	} else {
		fmt.Println(" - Existe")
		if isFinal {
			fmt.Println(" - Es final y se repite")
			return
		} else {
			fmt.Println(" - No es final")
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	}
}
