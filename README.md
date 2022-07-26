## Overview

The go-dyclog SDK provides simple log APIs for douyincloud mini program

## How It Works
```go
// main.go
package main

import (
	"context"
	
	"github.com/bytedance/go-dyclog"
)

func main() {
	// Start your vefaas function =D.
	vefaas.Start(handler)
}

// Define your handler function.
func handler(ctx context.Context, r *events.HTTPRequest) (*events.EventResponse, error) {
	// inject logid into context
	dyclog.InjectLogIDToCtx(ctx, r.Headers["x-tt-logid"])
	// close logger
	defer func() {
		_ = dyclog.Close()
	}()

	// Support Debug, Info, Warn, Error, Fatal 
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
## Features

- Simple architectural design
- Concise interface abstraction
- A certain scalability
- Support base log levels
- Support customization Formatter
- Support customization Writer

## Interfaces

``basis log methods``
- func Debug(format string, args ...interface{})
- func Info(format string, args ...interface{})
- func Notice(format string, args ...interface{})
- func Warn(format string, args ...interface{})
- func Error(format string, args ...interface{})
- func Fatal(format string, args ...interface{})

``with context basis log methods``
- func CtxDebug(ctx context.Context, format string, args ...interface{})
- func CtxInfo(ctx context.Context, format string, args ...interface{})
- func CtxNotice(ctx context.Context, format string, args ...interface{})
- func CtxWarn(ctx context.Context, format string, args ...interface{})
- func CtxError(ctx context.Context, format string, args ...interface{})
- func CtxFatal(ctx context.Context, format string, args ...interface{})

``setting methods``
- func SetWriter(writer LogWriter)
- func SetFormatter(formatter Formatter)
- func SetLevel(level Level)
- func Flush() error
- func Close() error

## Examples
*****The following two methods use ConsoleWriter by default to output logs through stdout*****
```go
// the first method
import "github.com/bytedance/go-dyclog"

func func main() {
    dyclog.Debug("test go-dyclog!")
    dyclog.CtxDebug(context.Background(), "test go-dyclog!")
    _ = dyclog.Close()
}
```

```go
// the second method
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

*****The following cases use FileWriter to write log files. Douyin Cloud does not support file log collection for the time being, so stay tuned*****
```go

import "github.com/bytedance/go-dyclog"

var Logger *dyclog.Logger

func init() {
    // the second args, RotationWindow, enum 0 Daily 1 Hourly
    Logger = dyclog.NewLogger(dyclog.NewFileWriter("./logs/dyc.log", 0))
}

func func main() {
    Logger.Debug("test go-dyclog!")
    Logger.CtxDebug(context.Background(), "test go-dyclog!")
    _ = Logger.Close()
}

```

## Security

If you discover a potential security issue in this project, or think you may
have discovered a security issue, we ask that you notify Bytedance Security via our [security center](https://security.bytedance.com/src) or [vulnerability reporting email](sec@bytedance.com).

Please do **not** create a public GitHub issue.

## License

This project is licensed under the [Apache-2.0 License](LICENSE).