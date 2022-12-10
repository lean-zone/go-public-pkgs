package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Default() *logrus.Logger {
	var Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	Log.SetReportCaller(true)
	Log.SetFormatter(NewStdoutFormatter())
	return Log
}

func New() *logrus.Logger {
	return logrus.New()
}
