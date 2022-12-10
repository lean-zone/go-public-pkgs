// @Author: Michael Lean
// @E-mail: 1013851072@qq.com
// @Create Time: UTC +8:00 2022/12/10 11:47:39

package test

import (
	"github.com/zhanglin-zl/go-public-pkgs/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	log := logger.Default()
	log.Info("test logger")

	log.AddHook(logger.NewFileHookByTime("logs", "[goTest]", "log"))
	log.Info("test file logger")
}
