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

import "context"

const ctxLogID = "DYC_LOGID"

func InjectLogIDToCtx(ctx context.Context, logID string) context.Context {
	return context.WithValue(ctx, ctxLogID, logID)
}

func GetLogIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return "-"
	}
	logID, ok := ctx.Value(ctxLogID).(string)
	if !ok {
		return "-"
	}

	return logID
}
