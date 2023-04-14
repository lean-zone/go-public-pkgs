// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/4/14 21:16:10

package random

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Println(Array1dFloat32(3))
	SetSeed(2)
	fmt.Println(Array1dFloat32(3))
}
