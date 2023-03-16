// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 19:45:04

package heap

import (
	"fmt"
	"github.com/lean-zone/go-public-pkgs/common"
	"github.com/lean-zone/go-public-pkgs/functools"
	"strings"
)

type Item[T common.GenericComparableType] struct {
	Key   string
	Value T
}

type MaxHeap[T common.GenericComparableType] struct {
	Data []*Item[T]
}

func (mh MaxHeap[T]) Len() int { return len(mh.Data) }

func (mh MaxHeap[T]) Less(i, j int) bool {
	return (mh.Data)[i].Value > (mh.Data)[j].Value
}

func (mh *MaxHeap[T]) Swap(i, j int) {
	((*mh).Data)[i], ((*mh).Data)[j] = ((*mh).Data)[j], ((*mh).Data)[i]
}

func (mh *MaxHeap[T]) Push(x interface{}) {
	(*mh).Data = append((*mh).Data, x.(*Item[T]))
}

func (mh *MaxHeap[T]) Pop() interface{} {
	old := (*mh).Data
	n := len(old)
	x := old[n-1]
	(*mh).Data = old[0 : n-1]
	return x
}

func (mh MaxHeap[T]) String() string {
	var str strings.Builder

	for i := 0; functools.Pow(2, i) < len(mh.Data); i++ {
		startIndex := functools.Pow(2, i) - 1
		endIndex := functools.Min(functools.Pow(2, i+1)-1, len(mh.Data))
		for j := startIndex; j < endIndex; j++ {
			str.WriteString(fmt.Sprintf("%v ", (*mh.Data[j]).Value))
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (mh *MaxHeap[T]) Down(startPos, endPos int) {
	if startPos > endPos {
		panic("heap down endPos must be greater or equal startPos!")
	}
	parent := startPos
	tmpItem := ((*mh).Data)[startPos]

	for {
		leftChild := 2*parent + 1
		if leftChild >= endPos || leftChild < 0 {
			break
		}
		swapPos := leftChild
		if rightChild := leftChild + 1; rightChild < endPos && mh.Less(rightChild, leftChild) {
			swapPos = rightChild
		}

		if ((*mh).Data)[swapPos].Value < tmpItem.Value {
			break
		}
		((*mh).Data)[parent] = ((*mh).Data)[swapPos]
		parent = swapPos
	}
	((*mh).Data)[parent] = tmpItem
}

type MinHeap[T common.GenericComparableType] struct {
	MaxHeap[T]
}

func (mh *MinHeap[T]) Less(i, j int) bool {
	return (mh.Data)[i].Value < (mh.Data)[j].Value
}
