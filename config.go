package mylogger

import "log/slog"

// 日志级别
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func mapLogLevel(level LogLevel) slog.Level {
	switch level {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// 配置结构体
type Config struct {
	Level       LogLevel `json:"level"`        // 日志级别
	OutputPaths []string `json:"output_paths"` // 输出路径（支持 "stdout" 或文件路径）
	FilePath    string   `json:"file_path"`    // 日志文件路径（当 OutputPaths 包含文件时）
	MaxSize     int      `json:"max_size"`     // 单个日志文件最大大小（字节）10M=10*1024*1024
	MaxBackups  int      `json:"max_backups"`  // 保留的日志文件数量
	MaxAge      int      `json:"max_age"`      // 保留日志的天数（天）
	Compress    bool     `json:"compress"`     // 是否压缩旧日志
	Console     bool     `json:"console"`      // 是否输出到控制台
}
