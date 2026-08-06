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
	"testing"

	"github.com/goccy/go-yaml"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

type testYAMLInput struct {
	A struct {
		B int    `yaml:"b"`
		C string `yaml:"c"`
	} `yaml:"a"`
}

func TestYAMLFormatterWrite(t *testing.T) {
	out := new(bytes.Buffer)
	f := NewYAMLFormatter()

	var input testYAMLInput
	input.A.B = 1
	input.A.C = "test"

	want := "---\n" +
		"a:\n" +
		"  b: 1\n" +
		"  c: test\n"

	err := f.Write(out, input)
	assert.NilError(t, err)
	assert.Assert(t, is.Equal(out.String(), want))
}

func TestYAMLFormatterOptionsOverrideDefaults(t *testing.T) {
	out := new(bytes.Buffer)

	f := NewYAMLFormatter(
		yaml.Indent(4),
	)

	var input testYAMLInput
	input.A.B = 1
	input.A.C = "test"

	want := "---\n" +
		"a:\n" +
		"    b: 1\n" +
		"    c: test\n"

	err := f.Write(out, input)
	assert.NilError(t, err)
	assert.Assert(t, is.Equal(out.String(), want))
}

func TestYAMLFormatterWriterError(t *testing.T) {
	f := NewYAMLFormatter()
	err := f.Write(newErrorWriter(), testYAMLInput{})
	assert.ErrorIs(t, err, errWriteFailed)
}
