// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 15:13:45

package functools

import (
	"github.com/lean-zone/go-public-pkgs/common"
)

func Distance2similarity(dist float32) float32 {
	return 1 - dist/2
}

func MapSum[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Min[T common.GenericComparableType](x ...T) T {
	if len(x) == 0 {
		panic("no enough param")
	} else if len(x) == 1 {
		return x[0]
	} else {
		minV := x[0]
		for i := 1; i < len(x); i++ {
			if minV > x[i] {
				minV = x[i]
			}
		}
		return minV
	}
}

func Max[T common.GenericComparableType](x ...T) T {
	if len(x) == 0 {
		panic("no enough param")
	} else if len(x) == 1 {
		return x[0]
	} else {
		maxV := x[0]
		for i := 1; i < len(x); i++ {
			if maxV < x[i] {
				maxV = x[i]
			}
		}
		return maxV
	}
}

func Pow(x, n int) int {
	ans := 1
	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}
