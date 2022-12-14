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

package dyclog

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebug(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Debug("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "DEBUG - logger_test.go:31 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxDebug(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "DEBUG 1234567890 logger_test.go:36 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Debug("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "DEBUG - logger_test.go:42 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxDebug(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "DEBUG 1234567890 logger_test.go:46 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}

func TestInfo(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Info("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "INFO - logger_test.go:57 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxInfo(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "INFO 1234567890 logger_test.go:62 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Info("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "INFO - logger_test.go:68 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxInfo(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "INFO 1234567890 logger_test.go:72 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}

func TestNotice(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Notice("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "NOTICE - logger_test.go:83 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxNotice(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "NOTICE 1234567890 logger_test.go:88 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Notice("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "NOTICE - logger_test.go:94 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxNotice(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "NOTICE 1234567890 logger_test.go:98 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}

func TestWarn(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Warn("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "WARN - logger_test.go:109 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxWarn(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "WARN 1234567890 logger_test.go:114 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Warn("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "WARN - logger_test.go:120 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxWarn(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "WARN 1234567890 logger_test.go:124 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}

func TestError(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Error("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "ERROR - logger_test.go:135 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxError(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "ERROR 1234567890 logger_test.go:140 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Error("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "ERROR - logger_test.go:146 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxError(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "ERROR 1234567890 logger_test.go:150 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}

func TestFatal(t *testing.T) {
	logger := NewLogger(new(BufferWriter))
	logger.SetFormatter(NewTextFormatter(false))
	err := "params is not valid"
	ip := GetLocalIP()

	logger.Fatal("err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "FATAL - logger_test.go:161 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	ctx := InjectLogIDToCtx(context.Background(), "1234567890")
	logger.CtxFatal(ctx, "err: %s", err)
	assert.Equal(t, logger.GetWriter().(*BufferWriter).String(), "FATAL 1234567890 logger_test.go:166 "+ip+" err: params is not valid\n")
	logger.GetWriter().(*BufferWriter).Reset()

	SetWriter(new(BufferWriter))
	SetFormatter(NewTextFormatter(false))
	Fatal("err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "FATAL - logger_test.go:172 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()

	CtxFatal(ctx, "err: %s", err)
	assert.Equal(t, GetLogger().GetWriter().(*BufferWriter).String(), "FATAL 1234567890 logger_test.go:176 "+ip+" err: params is not valid\n")
	GetLogger().GetWriter().(*BufferWriter).Reset()
}
