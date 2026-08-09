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
	"errors"
	"io"
	"reflect"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var errWriteFailed = errors.New("write failed")

type errorWriter struct {
	err error
}

func (w errorWriter) Write([]byte) (int, error) {
	return 0, w.err
}

func newErrorWriter() io.Writer {
	return errorWriter{err: errWriteFailed}
}

func TestNewFormatter(t *testing.T) {
	tests := []struct {
		name    string
		format  Format
		want    any
		wantErr error
	}{
		{
			name:   "text",
			format: TEXT,
			want:   &TextFormatter{},
		},
		{
			name:   "json",
			format: JSON,
			want:   &JSONFormatter{},
		},
		{
			name:   "yaml",
			format: YAML,
			want:   &YAMLFormatter{},
		},
		{
			name:   "text is case insensitive",
			format: "TeXt",
			want:   &TextFormatter{},
		},
		{
			name:   "json is case insensitive",
			format: "JsOn",
			want:   &JSONFormatter{},
		},
		{
			name:   "yaml is case insensitive",
			format: "yAML",
			want:   &YAMLFormatter{},
		},
		{
			name:    "unknown format",
			format:  "xml",
			wantErr: ErrFormat,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := NewFormatter(test.format)

			if test.wantErr != nil {
				assert.ErrorIs(t, err, test.wantErr)
				assert.Assert(t, f == nil)
			} else {
				assert.NilError(t, err)
				assert.Assert(t, is.Equal(reflect.TypeOf(f), reflect.TypeOf(test.want)))
			}
		})
	}
}

func TestFormats(t *testing.T) {
	formats := Formats()
	want := []string{"text", "json", "yaml"}

	assert.Equal(t, len(formats), len(want))

	for _, format := range want {
		assert.Assert(t, is.Contains(formats, format))
	}
}
