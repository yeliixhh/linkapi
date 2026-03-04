package logger

import (
	"log/slog"
	"os"
	"sync"
)

var Log *slog.Logger

var once sync.Once

// 初始化日志
func InitLogger() {

	once.Do(func() {
		opts := &slog.HandlerOptions{
			Level:     slog.LevelDebug, // 允许输出 Debug 级别
			AddSource: true,            // (可选) 添加文件名和行号，会有性能开销
		}

		handler := slog.NewJSONHandler(os.Stdout, opts)

		logger := slog.New(handler)

		Log = logger
	})

}
