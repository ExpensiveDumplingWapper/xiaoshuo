package zap

import (
	"os"
	"time"

	"xiaoshuo/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func encoderConfig(enableCaller bool) zapcore.EncoderConfig {
	if enableCaller {
		return zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	} else {
		return zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		}
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func RegisterLog(lc config.LogConfig) error {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lc.FileName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	var level zapcore.Level
	switch lc.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig(lc.EnableCaller)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		level,
	)
	zLogger := zap.New(core, zap.AddCaller())
	defer zLogger.Sync()
	zLog := zLogger.Sugar()
	zLog.Info()

	//logger.SetLogger(zLog)
	return nil
}
