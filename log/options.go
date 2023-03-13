package log

import (
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	StdLogOutputPath = "stdout"
	ConsoleFormat    = "console"
	JsonFormat       = "json"
)

// Options contains configuration items related to log.
type Options struct {
	Name string `json:"name"               yaml:"name"`
	// OutputPaths is a list of URLs or file paths to write logging output to.

	OutputPaths []string `json:"output-paths"       yaml:"output-paths"`

	// LogLevel is the minimum enabled logging level.
	LogLevel string `json:"level"              yaml:"level"`

	// Format sets the logger's encoding. Valid values are "json" and "console"
	Format string `json:"format"             yaml:"format"`

	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disable-caller"     yaml:"disable-caller"`

	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disable-stacktrace" yaml:"disable-stacktrace"`

	// Sampling sets a sampling policy. A nil SamplingConfig disables sampling.
	Sampling *zap.SamplingConfig `json:"sampling" yaml:"sampling"`

	// EnableColor set the logging level color
	EnableColor bool `json:"enable-color"       yaml:"enable-color"`

	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development"        yaml:"development"`

	// StacktraceLevel configures the Logger to record a stack trace for all messages at
	// or above a given level.
	StacktraceLevel string `json:"stacktraceLevel"    yaml:"stacktraceLevel"`

	// CallerSkip is the number of callers skipped by caller annotation
	CallerSkip int `json:"callerSkip"         yaml:"callerSkip"`
}

func (o *Options) String() string {
	data, _ := jsoniter.Marshal(o)
	return string(data)
}

func (o *Options) AddOutputPath(outputPath ...string) {
	if len(outputPath) == 0 {
		return
	} else {
		for i := 0; i < len(outputPath); i++ {
			o.OutputPaths = append(o.OutputPaths, outputPath[i])
		}
	}
}

func (o *Options) SetLogLevel(level string) {
	o.LogLevel = level
}

func (o *Options) SetStacktraceLevel(level string) {
	o.StacktraceLevel = level
}

func (o *Options) SetCallerSkip(skip int) {
	o.CallerSkip = skip
}

// Build constructs a global zap logger with the Options.
func (o *Options) Build(redirectStdLog bool) (*Logger, error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.LogLevel)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encodeLevel := zapcore.CapitalLevelEncoder
	if o.Format == ConsoleFormat && o.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zc := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling:          o.Sampling,
		Encoding:          o.Format,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "timestamp",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      zapcore.OmitKey,
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      encodeLevel,
			EncodeTime:       timeEncoder,
			EncodeDuration:   milliSecondsDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "  ",
		},
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}

	var traceLevel zapcore.Level
	if err := traceLevel.UnmarshalText([]byte(o.StacktraceLevel)); err != nil {
		traceLevel = zapcore.ErrorLevel
	}

	logger, err := zc.Build(zap.AddStacktrace(traceLevel), zap.AddCallerSkip(o.CallerSkip))

	if err != nil {
		return nil, err
	}
	if o.Name != "" {
		logger = logger.Named(o.Name)
	}
	if redirectStdLog {
		loggerForStd, _ := zc.Build(zap.AddStacktrace(traceLevel))
		zap.RedirectStdLog(loggerForStd)
	}
	return logger, nil
}

// NewOptions creates a new Options object with default parameters.
func NewOptions(logOutputPaths ...string) *Options {
	return &Options{
		OutputPaths:       logOutputPaths,
		LogLevel:          InfoLevel.String(),
		Format:            ConsoleFormat,
		DisableCaller:     true,
		DisableStacktrace: true,
		EnableColor:       false,
		Development:       false,
		StacktraceLevel:   ErrorLevel.String(),
		CallerSkip:        0,
	}
}

func stdOptions() *Options {
	return &Options{
		OutputPaths:       []string{StdLogOutputPath},
		LogLevel:          DebugLevel.String(),
		Format:            ConsoleFormat,
		DisableCaller:     false,
		DisableStacktrace: false,
		EnableColor:       true,
		Development:       true,
		StacktraceLevel:   ErrorLevel.String(),
		CallerSkip:        1,
	}
}

// NewProductionOptions creates an Options object in production environment.
func NewProductionOptions(logOutputPaths ...string) *Options {
	return &Options{
		OutputPaths:       logOutputPaths,
		LogLevel:          InfoLevel.String(),
		Format:            JsonFormat,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EnableColor:     false,
		Development:     false,
		StacktraceLevel: PanicLevel.String(),
		CallerSkip:      1,
	}
}
