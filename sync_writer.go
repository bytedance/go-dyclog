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
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type syncWriter struct {
	*bufio.Writer
	sync.Mutex
}

func newSyncWriter(w io.Writer) *syncWriter {
	return &syncWriter{
		Writer: bufio.NewWriterSize(w, 8*1024),
	}
}

type rotatedFile struct {
	w *syncWriter
	sync.WaitGroup
	done chan bool
}

func newRotatedFile(file io.WriteCloser) *rotatedFile {
	f := &rotatedFile{
		newSyncWriter(file),
		sync.WaitGroup{},
		make(chan bool),
	}
	f.Add(1)
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-f.done:
				ticker.Stop()
				_ = f.Flush()
				f.Done()
				return
			case <-ticker.C:
				err := f.Flush()
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "log writes file error: %s", err)
				}
			}
		}
	}()
	return f
}

func (f *rotatedFile) Close() error {
	f.done <- true
	f.Wait()
	return nil
}

func (f *rotatedFile) Rotate(w io.WriteCloser) {
	f.w.Lock()
	defer f.w.Unlock()
	_ = f.w.Flush()
	f.w.Reset(w)
}

func (f *rotatedFile) Flush() error {
	var err error
	f.w.Lock()
	defer f.w.Unlock()
	err = f.w.Writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (f *rotatedFile) Write(c []byte) (int, error) {
	f.w.Lock()
	defer f.w.Unlock()
	if f.w.Buffered()+len(c) > 4096 {
		_ = f.w.Flush()
	}
	n, err := f.w.Write(c)
	if err != nil {
		return n, err
	}
	if len(c) == 0 || c[len(c)-1] != '\n' {
		_, _ = f.w.Write([]byte{'\n'})
	}
	return n + 1, nil
}
