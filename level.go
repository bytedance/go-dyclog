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
	"errors"
)

var ErrInvalidLogLevel = errors.New("logger: invalid log level")

type Level int

const (
	InvalidLevel Level = iota - 1
	DEBUG
	INFO
	NOTICE
	WARN
	ERROR
	FATAL
)

const (
	sDebug  = "DEBUG"
	sInfo   = "INFO"
	sNotice = "NOTICE"
	sWarn   = "WARN"
	sError  = "ERROR"
	sFatal  = "FATAL"
)

var (
	levelNames = []string{
		sDebug,
		sInfo,
		sNotice,
		sWarn,
		sError,
		sFatal,
	}

	levelStrings = map[string]Level{
		sDebug:  DEBUG,
		sInfo:   INFO,
		sNotice: NOTICE,
		sWarn:   WARN,
		sError:  ERROR,
		sFatal:  FATAL,
	}
)

func (l Level) String() string {
	return levelNames[l]
}

func ParseLevel(s string) (Level, error) {
	l, ok := levelStrings[s]
	if !ok {
		return InvalidLevel, ErrInvalidLogLevel
	}

	return l, nil
}

func MustParseLevel(s string) Level {
	l, err := ParseLevel(s)
	if err != nil {
		panic(ErrInvalidLogLevel)
	}

	return l
}
