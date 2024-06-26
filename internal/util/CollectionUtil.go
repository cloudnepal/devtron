/*
 * Copyright (c) 2020-2024. Devtron Inc.
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

package util

import (
	"github.com/google/go-cmp/cmp"
	"sort"
)

func CompareUnOrdered(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	return cmp.Equal(a, b)
}

func GetTruncatedMessage(message string, maxLength int) string {
	_length := len(message)
	if _length == 0 {
		return message
	}
	_truncatedLength := maxLength
	if _length < _truncatedLength {
		return message
	} else {
		if _truncatedLength > 3 {
			return message[:(_truncatedLength-3)] + "..."
		}
		return message[:_truncatedLength]
	}
}
