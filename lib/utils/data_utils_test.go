/**
 * Copyright 2020 IBM Corp.
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

// Package util ...
package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeStringValue(t *testing.T) {
	assert.Equal(t, "", SafeStringValue(nil))
	hello := "hello"
	assert.Equal(t, "hello", SafeStringValue(&hello))
}

func TestStringHasValue(t *testing.T) {
	hello := "hello"
	assert.True(t, true, StringHasValue(&hello))
}
