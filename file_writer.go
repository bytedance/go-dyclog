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
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

const dateFormat = "2006-01-02_15"

// RotationWindow allows to claim which rotation window provider uses.
type RotationWindow int8

const (
	// Daily means rotate daily.
	Daily RotationWindow = iota
	// Hourly means rotate hourly.
	Hourly
)

// FileWriter provides a file rotated output to loggers,
// it is thread-safe and uses memory buffer to boost file writing performance.
type FileWriter struct {
	file           *rotatedFile
	filename       string
	rotationWindow RotationWindow
	fileCountLimit int

	currentTimeSeg time.Time
	sync.RWMutex
}

// NewFileWriter creates a FileWriter.
func NewFileWriter(filename string, window RotationWindow, options ...FileOption) LogWriter {
	w := &FileWriter{
		filename:       filename,
		rotationWindow: window,
	}
	file, err := w.loadFile()
	if err != nil {
		panic(err)
	}
	w.file = newRotatedFile(file)
	for _, op := range options {
		op(w)
	}
	return w
}

func (w *FileWriter) loadFile() (io.WriteCloser, error) {
	timedName, currentTimeSeg, err := timedFilename(w.filename)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(filepath.Dir(timedName), os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(timedName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	if _, err := os.Lstat(w.filename); err == nil {
		_ = os.Remove(w.filename)
	}
	_ = os.Symlink(filepath.Base(timedName), w.filename)
	w.currentTimeSeg = currentTimeSeg
	return file, nil
}

func (w *FileWriter) checkIfNeedRotate(logTime time.Time) error {
	var needRotate bool

	switch w.rotationWindow {
	case Daily:
		if w.currentTimeSeg.YearDay() != logTime.YearDay() {
			needRotate = true
		}
	case Hourly:
		if w.currentTimeSeg.Hour() != logTime.Hour() || w.currentTimeSeg.YearDay() != logTime.YearDay() {
			needRotate = true
		}
	}

	if needRotate {
		defer func() {
			go w.cleanFiles(w.fileCountLimit)
		}()
		if err := w.rotate(); err != nil {
			return err
		}
	}
	return nil
}

func getFileDate(name string) time.Time {
	sn := strings.Split(name, ".")
	t, _ := time.Parse(dateFormat, sn[len(sn)-1])
	return t
}

func (w *FileWriter) cleanFiles(limit int) {
	if limit <= 0 {
		return
	}
	logs := make([]string, 0)
	_ = filepath.Walk(filepath.Dir(w.filename), func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(path, w.filename+".") {
			logs = append(logs, path)
		}
		return nil
	})

	if len(logs) <= limit {
		return
	}
	sort.Slice(logs, func(i, j int) bool {
		return getFileDate(logs[i]).After(getFileDate(logs[j]))
	})
	for _, f := range logs[limit:] {
		_ = os.Remove(f)
	}
}

func (w *FileWriter) rotate() error {
	file, err := w.loadFile()
	if err != nil {
		return err
	}
	w.file.Rotate(file)
	return nil
}

func (w *FileWriter) Write(formatLog []byte) error {
	w.Lock()
	err := w.checkIfNeedRotate(time.Now())
	w.Unlock()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "write file %s error: %s\n", w.filename, err)
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "write file %s error: %s\n", w.filename, err)
	}
	_, err = w.file.Write(formatLog)
	return err
}

func (w *FileWriter) Close() error {
	return w.file.Close()
}

func (w *FileWriter) Flush() error {
	return w.file.Flush()
}

func timedFilename(filename string) (string, time.Time, error) {
	var now time.Time
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return "", now, err
	}
	now = time.Now()
	return absPath + "." + now.Format(dateFormat), now, nil
}

type FileOption func(writer *FileWriter)

func SetLimitFiles(n int) FileOption {
	return func(writer *FileWriter) {
		writer.fileCountLimit = n
	}
}
