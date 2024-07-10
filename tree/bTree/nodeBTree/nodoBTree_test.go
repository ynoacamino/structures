package nodebtree

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	myS := make([]*NodeBTree[int], 4)

	fmt.Println(len(myS))

	for i, slot := range myS {
		fmt.Println(i, slot)
	}
}
