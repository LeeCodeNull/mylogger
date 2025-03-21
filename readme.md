# mylogger: 轻量级 Go 日志库

**mylogger** 是一个基于 Go 标准库 `log/slog` 封装的高性能日志库，支持控制台输出、文件轮转、上下文字段传递和灵活配置。通过集成 `lumberjack` 库实现日志文件的自动轮转（按大小、时间、备份数量），适用于微服务、Web 应用等场景。

---

## 特性
- **多输出支持**：同时输出到控制台（`stdout`）和文件。
- **文件轮转**：按大小、天数或备份数量自动分割日志文件。
- **上下文传递字段**：通过 `context` 传递日志字段（如请求ID、用户信息）。
- **灵活配置**：支持日志级别、格式、路径等参数动态配置。
- **零第三方依赖**：仅依赖 Go 标准库和 `lumberjack`。

---

## 安装

```bash
go get github.com/LeeCodeNull/mylogger  # 替换为你的仓库地址
```
## 快速开始
```go
package main

import (
    "context"
    "github.com/LeeCodeNull/mylogger"
)

func main() {
    // 配置日志
    config := mylogger.Config{
        Level:       mylogger.InfoLevel,
        Console:     true,
        FilePath:    "./app.log",
        MaxSize:     10 * 1024 * 1024, // 10MB
        MaxBackups:  5,
        MaxAge:      7,    // 保留7天日志
        Compress:    true, // 压缩旧日志
    }

    // 初始化日志
    if err := mylogger.Init(config); err != nil {
        panic(err)
    }

    // 记录日志
    mylogger.Info("系统启动成功")
    mylogger.Debug("调试信息", "key", "value")
    mylogger.Error("发生错误", "error", "database connection failed")

}
```
## 高级用法
1. 上下文字段传递 : 通过 context 传递日志字段（如请求ID、用户信息）：
```go
ctx := mylogger.AppendContext(context.Background(), slog.String("request_id", "12345"))

// 在日志中使用 Context
mylogger.InfoContext(ctx, "处理请求", "status", "success")
```
2. 文件轮转配置
```go
config := mylogger.Config{
    MaxSize:     100 * 1024 * 1024, // 100MB
    MaxBackups:  10,
    MaxAge:      30, // 保留30天
    Compress:    true,
}
```
3. 全局日志函数
   函数名	说明	         示例
   Info	     信息级日志	   mylogger.Info("操作成功")
   Debug	 调试级日志	   mylogger.Debug("参数校验")
   Warn	     警告级日志	   mylogger.Warn("资源不足")
   Error	 错误级日志	    mylogger.Error("请求失败")
   InfoContext	带 Context 的信息日志	mylogger.InfoContext(ctx, "请求完成")
   ....
4. 反馈与问题
    提交 [GitHub 问题](https://github.com/LeeCodeNull/mylogger/issues) 或 [Pull Request](https://github.com/LeeCodeNull/mylogger/pulls)。