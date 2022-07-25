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
	"strconv"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	number := 10000
	limitGoroutines := 1000
	currentGoroutines := 0
	for i := 0; i < number; i++ {
		if currentGoroutines < limitGoroutines {
			go func(i int) {
				ctx := InjectLogIDToCtx(context.Background(), strconv.FormatInt(time.Now().UnixNano()+int64(i), 10))
				CtxDebug(ctx, "number: %d", i)
			}(i)
			currentGoroutines++
		} else {
			ctx := InjectLogIDToCtx(context.Background(), strconv.FormatInt(time.Now().UnixNano()+int64(i), 10))
			CtxDebug(ctx, "number: %d", i)
		}

	}
	_ = Close()
}
