// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 16:55:43

package functools

import (
	"fmt"
	"testing"
)

type newInt int

func TestFunctools(t *testing.T) {
	backNums := []newInt{1, 2}
	bn := newInt(3)
	fmt.Println(Min(backNums[0], bn))
	fmt.Println(Min(1, 2))
	fmt.Println(Min("2", "4", "1"))

	fmt.Println(Max(backNums[0], bn))
	fmt.Println(Max(1, 2))
	fmt.Println(Max("2", "4", "1"))

}
