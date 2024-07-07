package hashmap

import (
	l "structures/list/linkedList"
	city "structures/map/hash"
	n "structures/map/hashMap/nodeHash"
)

type HashMap[V any] struct {
	table []*l.LinkedList[n.NodeHash[string, V]] // slice de linkedList de nodeHash
	size  int

	capacity int

	loadFactor float32
}

func NewHashMap[V any](capacity int) *HashMap[V] {
	const (
		LOAD_FACTOR float32 = 0.75
		SIZE        int     = 0
	)

	return &HashMap[V]{
		table:      make([]*l.LinkedList[n.NodeHash[string, V]], capacity),
		size:       SIZE,
		capacity:   capacity,
		loadFactor: LOAD_FACTOR,
	}
}

func (hashMap *HashMap[V]) Put(key string, value *V) *V {
	nodeHash := n.NewNodeHash(key, value)

	hashedKey := city.CityHash32([]byte(key)) % uint32(hashMap.capacity)

	if hashMap.table[hashedKey] == nil {
		hashMap.table[hashedKey] = l.NewLinkedList(func(a, b n.NodeHash[string, V]) bool {
			return a.GetKey() == b.GetKey()
		})
	}

	list := hashMap.table[hashedKey]

	hashMap.size = hashMap.size + 1

	currentNode := list.GetFirstNode()

	for currentNode != nil {
		if currentNode.GetData().GetKey() == key {
			oldValue := currentNode.GetData().GetValue()
			currentNode.GetData().SetValue(value)

			return oldValue
		}

		currentNode = currentNode.GetNext()
	}

	list.AddFirst(nodeHash)

	return value
}

func (hashmap *HashMap[V]) Get(key string) *V {
	hashedKey := city.CityHash32([]byte(key)) % uint32(hashmap.capacity)

	slot := hashmap.table[hashedKey]

	if slot == nil {
		return nil
	}

	currentNode := slot.GetFirstNode()

	for currentNode != nil {
		if currentNode.GetData().GetKey() == key {
			return currentNode.GetData().GetValue()
		}

		currentNode = currentNode.GetNext()
	}

	return nil
}

func (hashMap *HashMap[V]) Remove(key string) *V {
	hashedKey := city.CityHash32([]byte(key)) % uint32(hashMap.capacity)

	matchNode := n.NewNodeHash[string, V](key, nil)

	slot := hashMap.table[hashedKey]

	oldValue := slot.RemoveMatch(matchNode)

	if oldValue != nil {
		hashMap.size = hashMap.size - 1
	}

	return oldValue.GetValue()
}

func (hashMap *HashMap[V]) CotainsKey(key string) bool {
	return true
}

func (hashMap *HashMap[V]) IsEmpty() bool {
	return hashMap.size == 0
}

func (hashMap *HashMap[V]) Size() int {
	return hashMap.size
}

func (hasMap *HashMap[V]) Clear() {
	hasMap.table = make([]*l.LinkedList[n.NodeHash[string, V]], hasMap.capacity)
	hasMap.size = 0
}
