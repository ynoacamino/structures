package arraylist

type ArrayList struct {
	capacity int
	size     int
	arr      [3]int
}

func NewArrayList(capacity int) *ArrayList {
	arrayList := ArrayList{
		capacity: capacity,
		size:     0,
		arr:      [3]int{},
	}

	return &arrayList
}
