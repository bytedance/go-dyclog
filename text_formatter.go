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
	"fmt"
	"sync"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

var formatPool = sync.Pool{
	New: func() interface{} {
		return bytes.Buffer{}
	},
}

type TextFormatter struct {
	enableColors    bool
	enableQuote     bool
	enableTimestamp bool
	timestampFormat string
}

func NewDefaultTextFormatter() *TextFormatter {
	return &TextFormatter{
		enableColors:    false,
		enableQuote:     false,
		enableTimestamp: false,
		timestampFormat: defaultTimestampFormat,
	}
}

func NewTextFormatter(enableColor bool) *TextFormatter {
	return &TextFormatter{
		enableColors:    enableColor,
		enableQuote:     false,
		enableTimestamp: false,
		timestampFormat: defaultTimestampFormat,
	}
}

func (f *TextFormatter) SetColor(enable bool) {
	f.enableColors = enable
}

func (f *TextFormatter) SetQuote(enable bool) {
	f.enableQuote = enable
}

func (f *TextFormatter) SetTimestamp(enable bool) {
	f.enableTimestamp = enable
}

func (f *TextFormatter) isColored() bool {
	return f.enableColors
}

func (f *TextFormatter) Format(l log) ([]byte, error) {
	fixedKeys := make([]string, 0, 6)
	if f.enableTimestamp {
		fixedKeys = append(fixedKeys, fieldKeyTime)
	}
	fixedKeys = append(fixedKeys, fieldKeyLevel)
	fixedKeys = append(fixedKeys, fieldKeyLogID)
	fixedKeys = append(fixedKeys, fieldKeyLocation)
	fixedKeys = append(fixedKeys, fieldKeyIP)
	fixedKeys = append(fixedKeys, fieldKeyMessage)

	b := formatPool.Get().(bytes.Buffer)
	defer formatPool.Put(b)

	if f.isColored() {
		f.encodeColorText(&b, l, fixedKeys)
	} else {
		f.encodeText(&b, l, fixedKeys)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *TextFormatter) encodeColorText(b *bytes.Buffer, entry log, fixedKeys []string) {
	var levelColor int
	switch entry.level {
	case DEBUG:
		levelColor = gray
	case WARN:
		levelColor = yellow
	case ERROR, FATAL:
		levelColor = red
	case INFO:
		levelColor = blue
	default:
		levelColor = blue
	}

	fmt.Fprintf(b, "\u001b[%dm", levelColor)
	f.encodeText(b, entry, fixedKeys)
}

func (f *TextFormatter) encodeText(b *bytes.Buffer, l log, fixedKeys []string) {
	for _, key := range fixedKeys {
		var value interface{}
		switch {
		case key == fieldKeyTime:
			value = l.time.Format(f.timestampFormat)
		case key == fieldKeyIP:
			value = GetLocalIP()
		case key == fieldKeyLogID:
			value = GetLogIDFromCtx(l.context)
		case key == fieldKeyLevel:
			value = l.level.String()
		case key == fieldKeyMessage:
			value = l.message
		case key == fieldKeyLocation:
			file, line := GetCallerLocation(l.caller)
			value = fmt.Sprintf("%s:%d", file, line)
		}

		if value == nil {
			continue
		}
		if b.Len() > 0 {
			b.WriteByte(' ')
		}

		stringVal, ok := value.(string)
		if !ok {
			stringVal = fmt.Sprint(value)
		}

		if !f.enableQuote {
			b.WriteString(stringVal)
		} else {
			b.WriteString(fmt.Sprintf("%q", stringVal))
		}
	}
}
