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
	"fmt"
	"os"
	"sync"
)

type AsyncWriter struct {
	LogWriter
	done    *sync.WaitGroup
	ch      chan []byte
	flush   chan bool
	flushed chan error
	omit    bool
}

func NewAsyncWriter(w LogWriter, omit bool) LogWriter {
	asyncWriter := &AsyncWriter{
		LogWriter: w,
		done:      &sync.WaitGroup{},
		ch:        make(chan []byte, 1024),
		flush:     make(chan bool),
		flushed:   make(chan error),
		omit:      omit,
	}
	go asyncWriter.runWorker()
	return asyncWriter
}

func (w *AsyncWriter) runWorker() {
	for {
		select {
		case formatLog, ok := <-w.ch:
			if !ok {
				return
			}
			err := w.LogWriter.Write(formatLog)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "log async writes error: %s\n", err)
			}
			w.done.Done()
		case <-w.flush:
			for i := 0; i < len(w.ch); i++ {
				formatLog := <-w.ch
				err := w.LogWriter.Write(formatLog)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "log async writes error: %s\n", err)
				}
				w.done.Done()
			}
			w.flushed <- w.LogWriter.Flush()
		}
	}
}

func (w *AsyncWriter) Write(log []byte) error {
	w.done.Add(1)
	if w.omit {
		select {
		case w.ch <- log:
		default:
			w.done.Done()
		}
	} else {
		w.ch <- log
	}
	return nil
}

func (w *AsyncWriter) Flush() error {
	w.flush <- true
	return <-w.flushed
}

func (w *AsyncWriter) Close() error {
	close(w.ch)
	w.done.Wait()
	return w.LogWriter.Close()
}
