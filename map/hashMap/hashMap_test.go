package hashmap

import (
	"testing"
)

func TestMain(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	v1 := hashMap.Get("hola")

	println("v1", *v1)
}
