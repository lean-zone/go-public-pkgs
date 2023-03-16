// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 19:44:29

package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func generateData() ([]string, []float32, []string, []float32) {
	keys1 := []string{
		"a00", "a01", "a02", "a03", "a04", "a05", "a06", "a07", "a08", "a09",
		"a10", "a11", "a12", "a13", "a14", "a15", "a16", "a17", "a18", "a19",
	}

	values1 := []float32{
		10, 19, 7, 11, 4, 5, 6, 8, 9, 12,
		16, 18, 15, 20, 17, 3, 13, 1, 14, 2,
	}

	keys2 := []string{
		"b00", "b01", "b02", "b03", "b04", "b05", "b06", "b07", "b08", "b09",
		"b10", "b11", "b12", "b13", "b14", "b15", "b16", "b17", "b18", "b19",
	}

	values2 := []float32{
		25, 36, 34, 32, 21, 39, 33, 28, 29, 26,
		24, 35, 38, 27, 31, 23, 37, 30, 40, 22,
	}

	return keys1, values1, keys2, values2
}

func TestMaxHeap(t *testing.T) {
	var mh1 = MaxHeap[float32]{}
	mh1.Data = []*Item[float32]{}

	keys1, values1, keys2, values2 := generateData()

	for i := 0; i < len(keys1); i++ {
		heap.Push(&mh1, &Item[float32]{
			Key:   keys1[i],
			Value: values1[i],
		})
	}
	fmt.Println("构建堆成功：")
	fmt.Println(mh1)
	Verify(&mh1, 0)

	SortTopK(&mh1)
	fmt.Println("排序之后：")
	fmt.Println(mh1)

	mh1.Data = []*Item[float32]{}
	for i := 0; i < len(keys2); i++ {
		heap.Push(&mh1, &Item[float32]{
			Key:   keys2[i],
			Value: values2[i],
		})
	}
	fmt.Println("构建堆成功：")
	fmt.Println(mh1)
	Verify(&mh1, 0)

	SortTopK(&mh1)
	fmt.Println("堆排序之后：")
	fmt.Println(mh1)
}

func TestMinHeap(t *testing.T) {
	var mh1 = MinHeap[float32]{}
	mh1.Data = []*Item[float32]{}

	keys1, values1, keys2, values2 := generateData()

	for i := 0; i < 20; i++ {
		heap.Push(&mh1, &Item[float32]{
			Key:   keys1[i],
			Value: values1[i],
		})
	}
	fmt.Println("构建堆成功：")
	fmt.Println(mh1)
	Verify(&mh1, 0)

	mh1.Data = []*Item[float32]{}
	for i := 0; i < 20; i++ {
		heap.Push(&mh1, &Item[float32]{
			Key:   keys2[i],
			Value: values2[i],
		})
	}
	fmt.Println("构建堆成功：")
	fmt.Println(mh1)
	Verify(&mh1, 0)

}
