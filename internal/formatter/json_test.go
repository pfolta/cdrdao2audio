// Copyright (c) 2026 Peter Folta
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package formatter

import (
	"bytes"
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

type testJSONInput struct {
	A struct {
		B int    `json:"b"`
		C string `json:"c"`
	} `json:"a"`
}

func TestJSONFormatterWrite(t *testing.T) {
	tests := []struct {
		name     string
		opts     []JSONOption
		expected string
	}{
		{
			name:     "default formatted json",
			expected: "{\n    \"a\": {\n        \"b\": 1,\n        \"c\": \"test\"\n    }\n}\n",
		},
		{
			name:     "compact json",
			opts:     []JSONOption{CompactJSON},
			expected: `{"a":{"b":1,"c":"test"}}` + "\n",
		},
	}

	var input testJSONInput
	input.A.B = 1
	input.A.C = "test"

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			f := NewJSONFormatter(test.opts...)

			err := f.Write(out, input)

			assert.NilError(t, err)
			assert.Assert(t, is.Equal(out.String(), test.expected))
		})
	}
}

func TestJSONFormatterOptionsOverrideDefaults(t *testing.T) {
	out := new(bytes.Buffer)

	f := NewJSONFormatter(func(enc *json.Encoder) {
		enc.SetIndent("", "  ")
	})

	var input testJSONInput
	input.A.B = 1
	input.A.C = "test"

	expected := "{\n  \"a\": {\n    \"b\": 1,\n    \"c\": \"test\"\n  }\n}\n"

	err := f.Write(out, input)

	assert.NilError(t, err)
	assert.Assert(t, is.Equal(out.String(), expected))
}

func TestJSONFormatterDisablesHTMLEscape(t *testing.T) {
	out := new(bytes.Buffer)
	f := NewJSONFormatter()

	input := map[string]string{
		"html": "<script>",
	}

	expected := "{\n    \"html\": \"<script>\"\n}\n"

	err := f.Write(out, input)

	assert.NilError(t, err)
	assert.Assert(t, is.Equal(out.String(), expected))
}

func TestJSONFormatterWritesValidJSON(t *testing.T) {
	out := new(bytes.Buffer)
	f := NewJSONFormatter()

	var input testJSONInput
	input.A.B = 1
	input.A.C = "test"

	err := f.Write(out, input)

	assert.NilError(t, err)

	var result testJSONInput

	err = json.Unmarshal(out.Bytes(), &result)

	assert.NilError(t, err)
	assert.Assert(t, is.Equal(result.A.C, "test"))
}

func TestJSONFormatterWriterError(t *testing.T) {
	f := NewJSONFormatter()

	var input testJSONInput
	input.A.B = 1
	input.A.C = "test"

	err := f.Write(newErrorWriter(), input)
	assert.ErrorIs(t, err, errWriteFailed)
}
