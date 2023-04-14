// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 19:45:04

package heap

import (
	"fmt"
	"github.com/lean-zone/go-public-pkgs/common"
	"github.com/lean-zone/go-public-pkgs/functools"
)

type ComparableData interface {
	LessThan(ComparableData) bool
	Print()
}

type DataHeap []ComparableData

func (mh *DataHeap) Len() int { return len(*mh) }

func (mh *DataHeap) Less(i, j int) bool {
	return (*mh)[i].LessThan((*mh)[j])
}

func (mh *DataHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *DataHeap) Push(x interface{}) {
	*mh = append(*mh, x.(ComparableData))
}

func (mh *DataHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

func (mh DataHeap) Print() {
	for i := 0; functools.Pow(2, i) < len(mh); i++ {
		startIndex := functools.Pow(2, i) - 1
		endIndex := functools.Min(functools.Pow(2, i+1)-1, len(mh))
		for j := startIndex; j < endIndex; j++ {
			(mh)[j].Print()
		}
		fmt.Println()
	}
	fmt.Println()
}

func (mh *DataHeap) Down(startPos, endPos int) {
	if startPos > endPos {
		panic("heap down endPos must be greater than or equal to startPos!")
	}
	parent := startPos
	tmpItem := (*mh)[startPos]

	for {
		leftChild := 2*parent + 1
		if leftChild >= endPos || leftChild < 0 {
			break
		}
		swapPos := leftChild
		if rightChild := leftChild + 1; rightChild < endPos && mh.Less(rightChild, leftChild) {
			swapPos = rightChild
		}

		if !(*mh)[swapPos].LessThan(tmpItem) {
			break
		}
		(*mh)[parent] = (*mh)[swapPos]
		parent = swapPos
	}
	(*mh)[parent] = tmpItem
}

type MinHeapItem[T common.GenericComparableType] struct {
	Key   string
	Value T
}

func (mhi MinHeapItem[T]) LessThan(another ComparableData) bool {
	return mhi.Value < another.(MinHeapItem[T]).Value
}

func (mhi MinHeapItem[T]) Print() {
	fmt.Print(fmt.Sprintf("%v ", mhi.Value))
}

type MaxHeapItem[T common.GenericComparableType] struct {
	MinHeapItem[T]
}

func (mhi MaxHeapItem[T]) LessThan(another ComparableData) bool {
	return mhi.Value > another.(MaxHeapItem[T]).Value
}
