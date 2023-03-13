// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 20:28:48

package common

// GenericComparableType After adding "~", the system is compatible with the base class type
type GenericComparableType interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string | ~int | ~uint | ~uintptr
}

type GenericSlice[T any] []T

type GenericMap[Key GenericComparableType, Value any] map[Key]Value

type GenericChan[T any] chan T

type GenericInterface[T any] interface {
	Print(data T)
}

type GenericStruct[T1 int | string, T2 int | bool] struct {
	Key   T1
	Value T2
}
