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
	"bytes"
)

type BufferWriter struct {
	buff bytes.Buffer
}

func (bw *BufferWriter) Bytes() []byte {
	return bw.buff.Bytes()
}

func (bw *BufferWriter) String() string {
	return bw.buff.String()
}

func (bw *BufferWriter) Write(formatLog []byte) error {
	var err error
	var n int
	writeLen := 0
	for writeLen < len(formatLog) {
		n, err = bw.buff.Write(formatLog[writeLen:])
		if err != nil {
			return err
		}
		writeLen += n
	}
	return err
}

func (bw *BufferWriter) Reset() error {
	bw.buff.Reset()
	return nil
}

func (bw *BufferWriter) Close() error {
	return nil
}

func (bw *BufferWriter) Flush() error {
	return nil
}