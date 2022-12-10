package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
)

type Hook struct {
	Writer     io.Writer
	Formatter  logrus.Formatter
	LogLevels  []logrus.Level
	ShowAtLine bool
}

func (h Hook) Levels() []logrus.Level {
	return h.LogLevels
}

func (h Hook) Fire(entry *logrus.Entry) error {
	if h.ShowAtLine {
		_, file, line, ok := runtime.Caller(5)
		if ok {
			entry.Data["at"] = fmt.Sprintf("%v:%v", file, line)
		}
	}

	serialized, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(serialized)
	return err
}

func makeLevels(maxLevel logrus.Level) []logrus.Level {
	allLevel := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	var filtered []logrus.Level
	for _, l := range allLevel {
		if l <= maxLevel {
			filtered = append(filtered, l)
		}
	}
	return filtered
}

// NewFileHookByTime 按时间滚动
func NewFileHookByTime(folder, filePrefix, fileExt string) *Hook {
	return &Hook{
		ShowAtLine: false,
		Writer:     NewFileCtx(folder, filePrefix, fileExt),
		Formatter:  NewFileFormatter(),
		LogLevels:  makeLevels(logrus.InfoLevel),
	}
}
