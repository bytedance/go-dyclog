## 抖音云开发者日志SDK

用于开发抖音云小程序Golang日志SDK及代码样例。

**注意：该项目处于内部测试阶段，API 行为可能会发生变更。**

### 快速上手

```Golang
// main.go
package main

import (
	"context"
	"fmt"

	"github.com/bytedance/go-dyclog"
)

func main() {
	// Start your vefaas function =D.
	vefaas.Start(handler)
}

// Define your handler function.
func handler(ctx context.Context, r *events.HTTPRequest) (*events.EventResponse, error) {
	// 注入logid
	dyclog.InjectLogIDToCtx(ctx, r.Headers["x-tt-logid"])
	// 刷新日志及关系日志处理器

	defer func() {
		_ = dyclog.Close()
	}()

	// 支持Debug, Info, Warn, Error, Fatal 
	dyclog.Debug("received new request: %s %s, request id: %s\n", r.HTTPMethod, r.Path, vefaascontext.RequestIdFrom)
	dyclog.CtxDebug(ctx, "received new request: %s %s, request id: %s\n", r.HTTPMethod, r.Path, vefaascontext.RequestIdFrom)

	return &events.EventResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte("Hello veFaaS!"),
	}, nil
}
```

### 接口方法

```Golang
// 支持Debug, Info, Notice, Warn, Error, Fatal 6中级别，可根据情况自行选择使用
func Debug(format string, args ...interface{})
func Info(format string, args ...interface{})
func Notice(format string, args ...interface{})
func Warn(format string, args ...interface{})
func Error(format string, args ...interface{})
func Fatal(format string, args ...interface{})

// 支持通过Context传递LogID等信息
func CtxDebug(ctx context.Context, format string, args ...interface{})
func CtxInfo(ctx context.Context, format string, args ...interface{})
func CtxNotice(ctx context.Context, format string, args ...interface{})
func CtxWarn(ctx context.Context, format string, args ...interface{})
func CtxError(ctx context.Context, format string, args ...interface{})
func CtxFatal(ctx context.Context, format string, args ...interface{})

// 设置writer
func SetWriter(writer LogWriter)
// 设置格式化
func SetFormatter(formatter Formatter)
// 日志日志级别，低于此级别的日志不会打印
func SetLevel(level Level)
// 刷新，在进行异步处理及带buffer的io时需要
func Flush() error
// 关系日志处理器，回收相关资源
func Close() error
```

### 使用方法

具体使用方法可以参考examples

*****以下两种方式默认使用ConsoleWriter，通过stdout输出日志*****
```go
// 第一种使用方式
import "github.com/bytedance/go-dyclog"

func func main() {
    dyclog.Debug("test go-dyclog!")
    dyclog.CtxDebug(context.Background(), "test go-dyclog!")
    _ = dyclog.Close()
}
```

```go
// 第二种使用方式
import "github.com/bytedance/go-dyclog"

var Logger *dyclog.Logger

func init() {
    Logger = NewDefaultLogger()
}
func func main() {
    Logger.Debug("test go-dyclog!")
    Logger.CtxDebug(context.Background(), "test go-dyclog!")
    _ = Logger.Close()
}
```

*****以下案例使用FileWriter写入日志文件，抖音云暂时不支持文件日志收集，敬请期待*****
```go

import "github.com/bytedance/go-dyclog"

var Logger *dyclog.Logger

func init() {
	// 第二个参数，限制文件个数
    Logger = dyclog.NewLogger(dyclog.NewFileWriter("./logs/dyc.log", 2))
}

func func main() {
    Logger.Debug("test go-dyclog!")
    Logger.CtxDebug(context.Background(), "test go-dyclog!")
    _ = Logger.Close()
}

```