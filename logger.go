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
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"
)

const minCallDepth = 4

type Logger struct {
	writer    LogWriter
	formatter Formatter
	level     Level
	callDepth int

	entryPool  sync.Pool
	formatPool sync.Pool
}

func NewDefaultLogger() *Logger {
	return &Logger{
		writer:    NewAsyncWriter(NewConsoleWriter(), false),
		formatter: NewDefaultTextFormatter(),
		level:     DEBUG,
		callDepth: minCallDepth,

		entryPool: sync.Pool{
			New: func() interface{} {
				return log{}
			},
		},
	}
}

func NewLogger(writer LogWriter) *Logger {
	return &Logger{
		writer:    writer,
		formatter: NewDefaultTextFormatter(),
		level:     DEBUG,
		callDepth: minCallDepth,

		entryPool: sync.Pool{},
		formatPool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (logger *Logger) GetWriter() LogWriter {
	return logger.writer
}

func (logger *Logger) SetWriter(writer LogWriter) {
	logger.writer = writer
}

func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.formatter = formatter
}

func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

func (logger *Logger) SetCallDepth(depth int) {
	logger.callDepth = depth
}

func (logger *Logger) Flush() error {
	return logger.writer.Flush()
}

func (logger *Logger) Close() error {
	return logger.writer.Close()
}

func (logger *Logger) newLog(ctx context.Context, level Level, format string, args ...interface{}) log {
	l, ok := logger.entryPool.Get().(log)
	if !ok {
		l = log{}
	}
	if ctx != nil {
		l.context = ctx
	}
	l.time = time.Now()
	l.level = level
	l.message = fmt.Sprintf(format, args...)
	l.caller = GetCaller(logger.callDepth)
	return l
}

func (logger *Logger) releaseLog(entry log) {
	logger.entryPool.Put(entry)
}

func (logger *Logger) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if level < logger.level {
		return
	}
	l := logger.newLog(ctx, level, format, args...)
	b, err := logger.formatter.Format(l)
	if err != nil {
		return
	}
	_ = logger.writer.Write(b)
	logger.releaseLog(l)
}

func (logger *Logger) Debug(format string, args ...interface{}) {
	logger.Logf(context.Background(), DEBUG, format, args...)
}

func (logger *Logger) Info(format string, args ...interface{}) {
	logger.Logf(context.Background(), INFO, format, args...)
}

func (logger *Logger) Notice(format string, args ...interface{}) {
	logger.Logf(context.Background(), NOTICE, format, args...)
}

func (logger *Logger) Warn(format string, args ...interface{}) {
	logger.Logf(context.Background(), WARN, format, args...)
}

func (logger *Logger) Error(format string, args ...interface{}) {
	logger.Logf(context.Background(), ERROR, format, args...)
}

func (logger *Logger) Fatal(format string, args ...interface{}) {
	logger.Logf(context.Background(), FATAL, format, args...)
}

func (logger *Logger) CtxDebug(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, DEBUG, format, args...)
}

func (logger *Logger) CtxInfo(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, INFO, format, args...)
}

func (logger *Logger) CtxNotice(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, NOTICE, format, args...)
}

func (logger *Logger) CtxWarn(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, WARN, format, args...)
}

func (logger *Logger) CtxError(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, ERROR, format, args...)
}

func (logger *Logger) CtxFatal(ctx context.Context, format string, args ...interface{}) {
	logger.Logf(ctx, FATAL, format, args...)
}
