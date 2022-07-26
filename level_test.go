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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelString(t *testing.T) {
	data := map[Level]string{
		DEBUG: sDebug,
		INFO:  sInfo,
		WARN:  sWarn,
		ERROR: sError,
		FATAL: sFatal,
	}

	for k, v := range data {
		assert.Equal(t, k.String(), v)
	}
}

func TestParseLevel(t *testing.T) {
	for k, v := range levelStrings {
		l, err := ParseLevel(k)
		assert.Equal(t, err, nil)
		assert.Equal(t, l, v)
	}
}
