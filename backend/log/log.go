package log

import (
	"time"
	"wordma/config"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger = zap.Logger

func InitLog() *zap.Logger {
	if config.DisableLog {
		zapLogger := zap.NewNop()
		zap.ReplaceGlobals(zapLogger)
		return zapLogger
	}

	level := zapcore.InfoLevel
	if config.DevelopMode {
		level = zapcore.DebugLevel
	}

	core := newZapCore(zapInitParams{
		logFile:    config.LogPath,
		logLevel:   level,
		maxSize:    500, // MB
		maxBackups: 30,  // the maximum number of old log files to retain
		maxAge:     15,  // the maximum number of days to retain old log files based on the
		compress:   true,
	})

	zapOpts := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
	if config.DevelopMode {
		zapOpts = append(zapOpts, zap.Development(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
	zapLogger := zap.New(core, zapOpts...)
	zap.ReplaceGlobals(zapLogger)
	return zapLogger
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

type zapInitParams struct {
	logFile    string
	logLevel   zapcore.Level
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
}

func newZapCore(p zapInitParams) zapcore.Core {
	// set log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(p.logLevel)

	// set encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = syslogTimeEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	encoderConfig.ConsoleSeparator = " "

	// Log to console
	methods := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(colorable.NewColorableStdout()),
			atomicLevel),
	}

	// Log to file
	if p.logFile != "" {
		fileEncoderConfig := encoderConfig
		fileEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // disable color

		methods = append(methods, zapcore.NewCore(
			zapcore.NewJSONEncoder(fileEncoderConfig), // write json format
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   p.logFile,
				MaxSize:    p.maxSize,
				MaxBackups: p.maxBackups,
				MaxAge:     p.maxAge,
				Compress:   p.compress,
			}),
			atomicLevel))
	}

	return zapcore.NewTee(methods...)
}
