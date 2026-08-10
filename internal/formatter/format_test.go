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
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFormats(t *testing.T) {
	formats := Formats()
	want := []string{"text", "json", "yaml"}

	assert.Assert(t, is.Equal(len(formats), len(want)))

	for _, format := range formats {
		assert.Assert(t, is.Contains(want, format))
	}
}

func TestFormatString(t *testing.T) {
	for _, format := range Formats() {
		t.Run(format, func(t *testing.T) {
			assert.Assert(t, is.Equal(Format(format).String(), format))
		})
	}
}

func TestFormatType(t *testing.T) {
	for _, format := range Formats() {
		t.Run(format, func(t *testing.T) {
			assert.Assert(t, is.Equal(Format(format).Type(), "string"))
		})
	}
}

func TestFormatSet(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Format
		wantErr error
	}{
		{
			name: "text",
			str:  "text",
			want: TEXT,
		},
		{
			name: "json",
			str:  "json",
			want: JSON,
		},
		{
			name: "yaml",
			str:  "yaml",
			want: YAML,
		},
		{
			name: "case insensitive",
			str:  "JsOn",
			want: JSON,
		},
		{
			name:    "unknown format",
			str:     "xml",
			wantErr: ErrFormat,
		},
		{
			name:    "empty format",
			str:     "",
			wantErr: ErrFormat,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			format := TEXT

			err := format.Set(test.str)

			if test.wantErr != nil {
				assert.Assert(t, is.ErrorIs(err, test.wantErr))
				assert.Assert(t, is.Equal(format, TEXT))
			} else {
				assert.NilError(t, err)
				assert.Assert(t, is.Equal(format, test.want))
			}
		})
	}
}
