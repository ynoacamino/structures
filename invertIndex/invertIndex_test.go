package invertindex

import (
	"testing"
)

func TestInvertIndex(t *testing.T) {
	index := NewInvertIndex[int](10)

	e1 := 10

	if index.Size() != 0 {
		t.Error("Expected size to be 0")
	}

	if !index.IsEmpty() {
		t.Error("Expected index to be empty")
	}

	index.Put("key", &e1)

	if index.Size() != 1 {
		t.Error("Expected size to be 1")
	}

	if index.IsEmpty() {
		t.Error("Expected index not to be empty")
	}

	list := index.Get("key")

	if list.Size() != 1 {
		t.Error("Expected list size to be 1")
	}

	if list.Get(0) != &e1 {
		t.Error("Expected value to be 10")
	}

	index.Remove("key")

	if index.Size() != 0 {
		t.Error("Expected size to be 0")
	}

	if !index.IsEmpty() {
		t.Error("Expected index to be empty")
	}
}
