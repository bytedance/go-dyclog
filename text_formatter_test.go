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
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTextFormatter(t *testing.T) {
	now := time.Now()
	l := log{}
	l.level = DEBUG
	l.time = now
	l.message = "test text formatter!"
	l.context = context.Background()
	l.caller = GetCaller(1)

	f := NewDefaultTextFormatter()
	f.SetColor(true)
	b, e := f.Format(l)
	if e != nil {
		t.Errorf("err: %v", e)
	}
	ip := GetLocalIP()
	assert.Equal(t, "\x1b[37m DEBUG - text_formatter_test.go:33 "+ip+" test text formatter!\n", string(b))

	f.SetColor(false)
	b, e = f.Format(l)
	assert.Equal(t, "DEBUG - text_formatter_test.go:33 "+ip+" test text formatter!\n", string(b))

	f.SetTimestamp(true)
	b, e = f.Format(l)
	assert.Equal(t, now.Format(f.timestampFormat)+" DEBUG - text_formatter_test.go:33 "+ip+" test text formatter!\n", string(b))

	f.SetQuote(true)
	b, e = f.Format(l)
	assert.Equal(t, "\""+now.Format(f.timestampFormat)+"\""+" \"DEBUG\" \"-\" \"text_formatter_test.go:33\" \""+ip+"\" \"test text formatter!\"\n", string(b))
}
