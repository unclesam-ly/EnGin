package logger

import (
	"EnGin/internal/conf"
	"fmt"
	"os"
	"path"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(cfg conf.LoggerConfig) (*zap.Logger, error) {
	dir := path.Dir(cfg.Filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}
	var cores []zapcore.Core
	// 1. 文件输出 (Info 级别及以上)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= getLogLevel(cfg.Level) && lvl < zap.ErrorLevel
	})

	infoWriter := getLogWriter(cfg.Filename, cfg)
	cores = append(cores, zapcore.NewCore(getFileEncoder(cfg.Prefix), infoWriter, infoLevel))

	// 2. 错误文件输出 (Error 级别及以上)
	errFilename := strings.Replace(cfg.Filename, ".log", "_error.log", 1)
	errWriter := getLogWriter(errFilename, cfg)
	cores = append(cores, zapcore.NewCore(getFileEncoder(cfg.Prefix), errWriter, zap.ErrorLevel))

	// 3. 控制台输出
	cores = append(cores, zapcore.NewCore(getConsoleEncoder(cfg.Prefix), zapcore.AddSync(os.Stdout), getLogLevel(cfg.Level)))
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger, nil
}

func getLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getLogWriter(filename string, cfg conf.LoggerConfig) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	})
}

func getFileEncoder(prefix string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	if prefix != "" {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(fmt.Sprintf("[%s] 2006-01-02 15:04:05", prefix))
	} else {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder(prefix string) zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	if prefix != "" {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(fmt.Sprintf("[%s] 2006-01-02 15:04:05", prefix))
	} else {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
