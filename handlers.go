package mylogger

import (
	"os"
	"sync"

	"context"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
)

// 全局锁，用于安全关闭 lumberjack.Logger
var (
	once    sync.Once
	lg      *lumberjack.Logger
	closeFn func()
)

// 创建文件 Handler
func NewFileHandler(config Config) (slog.Handler, error) {
	// 初始化 lumberjack.Logger 并保存实例
	lg = &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	// 注册关闭函数
	closeFn = func() {
		lg.Close()
	}

	return slog.NewJSONHandler(lg, &slog.HandlerOptions{
		AddSource: true,
		Level:     mapLogLevel(config.Level),
	}), nil
}

// 组合 Handler（支持控制台和文件）
type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(config Config) (slog.Handler, error) {
	var handlers []slog.Handler

	// 添加控制台输出
	if config.Console {
		handlers = append(handlers, slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     mapLogLevel(config.Level),
		}))
	}

	// 添加文件输出
	if config.FilePath != "" {
		fileHandler, err := NewFileHandler(config)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, fileHandler)
	}

	return &MultiHandler{
		handlers: handlers,
	}, nil
}

// 实现 Handler 接口
func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	//添加context中的字段
	if attrs := GetAttrsFromContext(ctx); attrs != nil {
		r.AddAttrs(attrs...)
	}
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}
func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}
