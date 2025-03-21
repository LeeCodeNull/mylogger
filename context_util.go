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

const slogContextKey = "slog_fields"

// 将字段附加到 Context
func AppendContext(ctx context.Context, attrs ...slog.Attr) context.Context {
	existingAttrs, ok := ctx.Value(slogContextKey).([]slog.Attr)
	if !ok {
		existingAttrs = []slog.Attr{}
	}
	return context.WithValue(ctx, slogContextKey, append(existingAttrs, attrs...))
}

// 从 Context 中提取字段
func GetAttrsFromContext(ctx context.Context) []slog.Attr {
	if attrs, ok := ctx.Value(slogContextKey).([]slog.Attr); ok {
		return attrs
	}
	return nil
}
