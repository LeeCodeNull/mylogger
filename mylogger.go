/*
*
@author biubiu
@date:2025/3/21
*/
package mylogger

import (
	"context"
	"log/slog"
)

var globalLogger *slog.Logger

// 初始化日志
func Init(config Config) error {
	handler, err := NewMultiHandler(config)
	if err != nil {
		return err
	}

	// 创建 Logger
	globalLogger = slog.New(handler)

	return nil
}

// 全局日志函数
func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

// 上下文日志记录
func InfoContext(ctx context.Context, msg string, args ...any) {
	globalLogger.InfoContext(ctx, msg, args...)
}

// 关闭日志资源（确保关闭 lumberjack.Logger）
func Close() {
	once.Do(closeFn)
}
