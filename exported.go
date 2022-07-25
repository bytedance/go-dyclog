/*
 * Copyright 2022 douyincloud
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

import "context"

var defaultLogger *Logger

func init() {
	defaultLogger = NewDefaultLogger()
}

func GetLogger() *Logger {
	return defaultLogger
}

func SetWriter(writer LogWriter) {
	defaultLogger.SetWriter(writer)
}

func SetFormatter(formatter Formatter) {
	defaultLogger.SetFormatter(formatter)
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

func SetCallDepth(depth int) {
	defaultLogger.SetCallDepth(depth)
}

func Flush() error {
	return defaultLogger.Flush()
}

func Close() error {
	return defaultLogger.Close()
}

func Debug(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), DEBUG, format, args...)
}

func Info(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), INFO, format, args...)
}

func Notice(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), NOTICE, format, args...)
}

func Warn(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), WARN, format, args...)
}

func Error(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), ERROR, format, args...)
}

func Fatal(format string, args ...interface{}) {
	defaultLogger.Logf(context.Background(), FATAL, format, args...)
}

func CtxDebug(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, DEBUG, format, args...)
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, INFO, format, args...)
}

func CtxNotice(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, NOTICE, format, args...)
}

func CtxWarn(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, WARN, format, args...)
}

func CtxError(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, ERROR, format, args...)
}

func CtxFatal(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.Logf(ctx, FATAL, format, args...)
}
