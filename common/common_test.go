// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2023/2/27 21:53:14

package common

import (
	"fmt"
	"testing"
)

func TestCommon(t *testing.T) {
	var mp GenericMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
	fmt.Println(mp)
}
