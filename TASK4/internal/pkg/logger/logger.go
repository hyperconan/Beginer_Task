package logger

import (
	"log"

	"go.uber.org/zap"
)

var (
	// L 是全局 zap.Logger
	L *zap.Logger
	// S 是便捷的 sugared logger
	S *zap.SugaredLogger
)

// Init 初始化日志组件，默认使用生产配置。
func Init() {
	if L != nil {
		return
	}

	cfg := zap.NewProductionConfig()
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}

	L = logger
	S = logger.Sugar()
}

// Sync 将日志缓冲落盘
func Sync() {
	if L != nil {
		_ = L.Sync()
	}
}
