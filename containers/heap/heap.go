// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/28 16:17:17

package heap

import (
	"log"
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

func Heapify(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n-1)
	}
}

func Push(h Interface, x interface{}) {
	h.Push(x)
	Up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n-1)
	return h.Pop()
}

func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n-1) {
			Up(h, i)
		}
	}
	return h.Pop()
}

func Fix(h Interface, i int) {
	if !down(h, i, h.Len()-1) {
		Up(h, i)
	}
}

func Up(h Interface, i int) {
	for {
		parent := (i - 1) / 2 // parent
		if i == parent || !h.Less(i, parent) {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}

func down(h Interface, startPos, endPos int) bool {
	if startPos > endPos {
		panic("the 'endPos' must be greater than or equal to 'startPos'!")
	}
	i := startPos
	for {
		leftChild := 2*i + 1
		if leftChild > endPos || leftChild < 0 {
			break
		}
		swapPos := leftChild
		if rightChild := leftChild + 1; rightChild <= endPos && h.Less(rightChild, leftChild) {
			swapPos = rightChild
		}
		if !h.Less(swapPos, i) {
			break
		}
		h.Swap(i, swapPos)
		i = swapPos
	}

	return i > startPos
}

// SortTopK 堆建好了之后，重新排序，利用最大（小）堆获取TopK最小（大）值
func SortTopK(h Interface) {
	if h.Len() <= 1 {
		return
	}

	for i := h.Len() - 1; i > 0; {
		h.Swap(0, i)
		i--
		down(h, 0, i)
	}
}

func Verify(h Interface, i int) {
	n := h.Len()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h.Less(j1, i) {
			log.Fatalf("failed to verify, parent index: %v, leftchild index: %v", i, j1)
			return
		}
		Verify(h, j1)
	}
	if j2 < n {
		if h.Less(j2, i) {
			log.Fatalf("failed to verify, parent index: %v, rightchild index: %v", i, j2)
			return
		}
		Verify(h, j2)
	}
}
