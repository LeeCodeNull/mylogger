/*
*
@author biubiu
@date:2025/3/21
*/
package mylogger

import (
	"context"
	"log/slog"
	"testing"
)

var conf = Config{
	Compress: false,
	Console:  true,
	FilePath: "/root/a.txt",
	Level:    InfoLevel,
}

func init() {
	Init(conf)
}

func TestLogger(t *testing.T) {
	globalLogger.WithGroup("api").Info("api test")
	ctx := AppendContext(context.Background(), slog.String("req_id", "this is testabc"))
	globalLogger.InfoContext(ctx, "我是 ctx")

}
func TestAbc(t *testing.T) {
	t.Log("abccc")
}
