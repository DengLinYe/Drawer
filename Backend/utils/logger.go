package utils

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *slog.Logger
var LogWriter io.Writer

// 初始化日志记录器
func InitLogger() {
	if err := os.MkdirAll("log", os.ModePerm); err != nil {
		panic("无法创建日志目录: " + err.Error())
	}

	// 设置日志文件滚动
	fileWriter := &lumberjack.Logger{
		Filename:   filepath.Join("log", "app.log"),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	LogWriter = io.MultiWriter(os.Stdout, fileWriter)

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(LogWriter, opts)
	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}
