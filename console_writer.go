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

import (
	"io"
	"os"
)

type ConsoleWriter struct {
	writer io.Writer
}

func NewConsoleWriter() LogWriter {
	w := &ConsoleWriter{
		writer: os.Stdout,
	}

	return w
}

func (cw *ConsoleWriter) Write(formatLog []byte) error {
	_, err := cw.writer.Write(formatLog)
	return err
}

func (cw *ConsoleWriter) Close() error {
	return nil
}

func (cw *ConsoleWriter) Flush() error {
	return nil
}
