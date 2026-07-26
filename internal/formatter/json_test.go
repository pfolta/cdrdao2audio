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

func TestJSONFormatterWrite(t *testing.T) {
	tests := []struct {
		name     string
		opts     []JSONOption
		expected string
	}{
		{
			name: "default formatted json",
			expected: `{
    "name": "cdrdao2audio",
    "version": "1.0.0"
}
`,
		},
		{
			name: "compact json",
			opts: []JSONOption{CompactJSON},
			expected: `{"name":"cdrdao2audio","version":"1.0.0"}
`,
		},
	}

	value := map[string]string{
		"name":    "cdrdao2audio",
		"version": "1.0.0",
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			f := NewJSONFormatter(test.opts...)

			err := f.Write(out, value)

			assert.NilError(t, err)
			assert.Equal(t, out.String(), test.expected)
		})
	}
}

func TestJSONFormatterDisablesHTMLEscape(t *testing.T) {
	out := new(bytes.Buffer)
	f := NewJSONFormatter()

	err := f.Write(out, map[string]string{
		"html": "<script>",
	})

	assert.NilError(t, err)
	assert.Assert(t, is.Contains(out.String(), "<script>"))
}

func TestJSONFormatterWritesValidJSON(t *testing.T) {
	out := new(bytes.Buffer)
	f := NewJSONFormatter()

	err := f.Write(out, map[string]string{
		"name": "cdrdao2audio",
	})

	assert.NilError(t, err)

	var result map[string]string

	err = json.Unmarshal(out.Bytes(), &result)

	assert.NilError(t, err)
	assert.Equal(t, result["name"], "cdrdao2audio")
}

func TestJSONFormatterWriterError(t *testing.T) {
	f := NewJSONFormatter()

	err := f.Write(newErrorWriter(), map[string]string{
		"name": "cdrdao2audio",
	})

	assert.ErrorIs(t, err, errWriteFailed)
}
