/*
 * Copyright 2022 ByteDance and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"strconv"
	"time"

	"github.com/bytedance/dyclog"
)

func main() {
	// test logger method
	logger := dyclog.NewDefaultLogger()
	req, rsp := "123", "456"
	logger.Debug("req: %s, rsp: %s", req, rsp)
	logger.Info("req: %s, rsp: %s", req, rsp)
	logger.Warn("req: %s, rsp: %s", req, rsp)
	logger.Error("req: %s, rsp: %s", req, rsp)
	logger.Fatal("req: %s, rsp: %s", req, rsp)

	// test context
	ctx := dyclog.InjectLogIDToCtx(context.Background(), strconv.FormatInt(time.Now().UnixNano(), 10))
	logger.CtxDebug(ctx, "hello logger!")

	_ = logger.Flush()
	_ = logger.Close()

	// test export method
	dyclog.Debug("hello, logger!")
	dyclog.CtxDebug(ctx, "hello, logger!")
	_ = dyclog.Flush()
	_ = dyclog.Close()

	// test sync file_writer
	logger = dyclog.NewLogger(dyclog.NewFileWriter("./logs/dyc.log", 2))
	logger.SetFormatter(dyclog.NewTextFormatter(false))
	err := "oops, something is wrong!"
	logger.Debug("err: %s", err)
	logger.Debug("123")
	logger.CtxFatal(ctx, "fatal error")
	_ = logger.Flush()
	_ = logger.Close()

	// test async file_writer
	logger = dyclog.NewLogger(dyclog.NewAsyncWriter(dyclog.NewFileWriter("./logs/dyc.log", 2), true))
	formatter := dyclog.NewTextFormatter(false)
	formatter.SetTimestamp(true)
	logger.SetFormatter(formatter)
	message := "this is async file writer!"
	logger.Debug("message: %s", message)
	logger.CtxFatal(ctx, message)
	_ = logger.Flush()
	_ = logger.Close()
}
