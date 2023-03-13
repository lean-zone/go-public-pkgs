package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

var (
	std *Logger
	mu  sync.Mutex
)

func init() {
	mu.Lock()
	defer mu.Unlock()
	if std == nil {
		std, _ = stdOptions().Build(true)
	}
}

func Default() *Logger {
	return std
}

func Log(lvl zapcore.Level, msg string, fields ...Field) {
	std.Log(lvl, msg, fields...)
}

func GetLogLevel() zapcore.Level {
	return std.Level()
}

func RunProductionMode(logOutputPaths ...string) {
	mu.Lock()
	defer mu.Unlock()
	std, _ = NewProductionOptions(logOutputPaths...).Build(true)
}

func StdFromOptions(o *Options) {
	mu.Lock()
	defer mu.Unlock()
	std, _ = o.Build(true)
}

// StdErrLogger returns logger of standard library which writes to supplied zap logger at error level.
func StdErrLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std, zapcore.ErrorLevel); err == nil {
		return l
	}

	return nil
}

// StdInfoLogger returns logger of standard library which writes to supplied zap
// logger at info level.
func StdInfoLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std, zapcore.InfoLevel); err == nil {
		return l
	}

	return nil
}

// Sync calls the underlying Core's Sync method, flushing any buffered
// log entries. Applications should take care to call Sync before exiting.
func Sync() {
	err := std.Sync()
	if err != nil {
		return
	}
}

// WithName adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func WithName(s string) *Logger { return std.Named(s) }

// WithKeyAndValues creates a child logger and adds Zap fields to it.
func WithKeyAndValues(keysAndValues ...Field) *Logger { return std.With(keysAndValues...) }

// Debug method output debug level log.
func Debug(msg string, fields ...Field) {
	std.Debug(msg, fields...)
}

// Debugf method output debug level log.
func Debugf(format string, v ...any) {
	std.Sugar().Debugf(format, v...)
}

// Debugw method output debug level log.
func Debugw(msg string, keysAndValues ...any) {
	std.Sugar().Debugw(msg, keysAndValues...)
}

// Info method output info level log.
func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

// Infof method output info level log.
func Infof(format string, v ...any) {
	std.Sugar().Infof(format, v...)
}

// Infow method output info level log.
func Infow(msg string, keysAndValues ...any) {
	std.Sugar().Infow(msg, keysAndValues...)
}

// Warn method output warning level log.
func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}

// Warnf method output warning level log.
func Warnf(format string, v ...any) {
	std.Sugar().Warnf(format, v...)
}

// Warnw method output warning level log.
func Warnw(msg string, keysAndValues ...any) {
	std.Sugar().Warnw(msg, keysAndValues...)
}

// Error method output error level log.
func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}

// Errorf method output error level log.
func Errorf(format string, v ...any) {
	std.Sugar().Errorf(format, v...)
}

// Errorw method output error level log.
func Errorw(msg string, keysAndValues ...any) {
	std.Sugar().Errorw(msg, keysAndValues...)
}

// Panic method output panic level log and shutdown application.
func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}

// Panicf method output panic level log and shutdown application.
func Panicf(format string, v ...any) {
	std.Sugar().Panicf(format, v...)
}

// Panicw method output panic level log.
func Panicw(msg string, keysAndValues ...any) {
	std.Sugar().Panicw(msg, keysAndValues...)
}

func DPanic(msg string, fields ...Field) {
	std.DPanic(msg, fields...)
}

func DPanicf(format string, v ...any) {
	std.Sugar().DPanicf(format, v...)
}

func DPanicw(msg string, keysAndValues ...any) {
	std.Sugar().DPanicw(msg, keysAndValues...)
}

// Fatal method output fatal level log.
func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}

// Fatalf method output fatal level log.
func Fatalf(format string, v ...any) {
	std.Sugar().Fatalf(format, v...)
}

// Fatalw method output Fatalw level log.
func Fatalw(msg string, keysAndValues ...any) {
	std.Sugar().Fatalw(msg, keysAndValues...)
}

// Compatible with the standard log library

func Println(msg ...any) {
	std.Sugar().Infoln(msg...)
}

func Print(msg ...any) {
	std.Sugar().Infoln(msg...)
}

func Printf(format string, v ...any) {
	std.Sugar().Infof(format, v...)
}

func Panicln(msg ...any) {
	std.Sugar().Panicln(msg...)
}

func Fatalln(msg ...any) {
	std.Sugar().Fatalln(msg...)
}
