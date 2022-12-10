package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

type Color int

// Foreground text colors.
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors.
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func stringColored(msg string, c Color) string {
	return fmt.Sprintf("\033[1;%v;40m%s\033[0m", c, msg)
}

func stringColoredHighlight(msg string, c Color) string {
	return fmt.Sprintf("\033[%vm%s\033[0m", c, msg)
}

func printColored(msg string, logLevel logrus.Level) string {
	var msgColored string
	switch logLevel {
	case logrus.DebugLevel, logrus.TraceLevel:
		msgColored = stringColoredHighlight(msg, FgHiMagenta)
	case logrus.WarnLevel:
		msgColored = stringColoredHighlight(msg, FgHiYellow)
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		msgColored = stringColoredHighlight(msg, FgHiRed)
	case logrus.InfoLevel:
		msgColored = stringColoredHighlight(msg, FgHiBlue)
	default:
		msgColored = stringColoredHighlight(msg, FgHiBlue)
	}
	return msgColored
}

type StdoutFormatter struct {
	Formatter     logrus.Formatter
	DisableColors bool
}

func NewStdoutFormatter() *StdoutFormatter {
	return &StdoutFormatter{
		DisableColors: false,
	}
}

func (m *StdoutFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	var newLog string

	if entry.HasCaller() {
		fName, _ := filepath.Abs(entry.Caller.File)
		var levelName string
		if m.DisableColors {
			levelName = fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String()))
		} else {
			levelName = printColored(fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())), entry.Level)
		}
		newLog = fmt.Sprintf("%s[%s][%s:%d] %s",
			levelName, timestamp, fName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s][%s] %s", strings.ToUpper(entry.Level.String()), timestamp, entry.Message)
	}
	b.WriteString(newLog)
	b.WriteString("\n")
	return b.Bytes(), nil
}

type FileFormatter struct {
	Formatter logrus.Formatter
}

func NewFileFormatter() *FileFormatter {
	return &FileFormatter{}
}

func (ff *FileFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	var newLog string

	if entry.HasCaller() {
		fName, _ := filepath.Abs(entry.Caller.File)
		levelName := fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String()))
		newLog = fmt.Sprintf("%s[%s][%s:%d] %s",
			levelName, timestamp, fName, entry.Caller.Line, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s][%s] %s", strings.ToUpper(entry.Level.String()), timestamp, entry.Message)
	}
	b.WriteString(newLog)
	b.WriteString("\n")
	return b.Bytes(), nil
}
