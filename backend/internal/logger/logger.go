package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logFilePath string, debug bool) (*zap.Logger, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	endcoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	fileEndcoder := zapcore.NewJSONEncoder(endcoderConfig)
	consoleEndcoder := zapcore.NewConsoleEncoder(endcoderConfig)

	var logLevel zapcore.Level
	if debug {
		logLevel = zap.DebugLevel
	} else {
		logLevel = zap.InfoLevel
	}

	fileWriter := zapcore.AddSync(logFile)
	fileCore := zapcore.NewCore(fileEndcoder, fileWriter, logLevel)

	consoleWriter := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEndcoder, consoleWriter, logLevel)

	core := zapcore.NewTee(fileCore, consoleCore)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return logger, nil
}

func IsDebug() bool {
	if debugLog, exists := os.LookupEnv("DEBUG"); debugLog == "True" || !exists {
		return true
	}
	return false
}
