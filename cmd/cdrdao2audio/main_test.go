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

package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestColorFlag(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		env       map[string]string
		wantColor bool
	}{
		{
			name:      "default does not force color",
			args:      []string{"--help"},
			wantColor: false,
		},
		{
			name:      "`always` enables color",
			args:      []string{"--color", "always", "--help"},
			wantColor: true,
		},
		{
			name: "`never` flag overrides `CLICOLOR` and `CLICOLOR_FORCE`",
			args: []string{"--color", "never", "--help"},
			env: map[string]string{
				"CLICOLOR":       "true",
				"CLICOLOR_FORCE": "true",
			},
			wantColor: false,
		},
		{
			name: "`always` flag overrides `CLICOLOR`",
			args: []string{"--color", "always", "--help"},
			env: map[string]string{
				"CLICOLOR": "false",
			},
			wantColor: true,
		},
		{
			name: "`always` flag overrides `NO_COLOR`",
			args: []string{"--color", "always", "--help"},
			env: map[string]string{
				"NO_COLOR": "true",
			},
			wantColor: true,
		},
		{
			name: "existing `CLICOLOR_FORCE` is preserved without flag",
			args: []string{"--help"},
			env: map[string]string{
				"CLICOLOR_FORCE": "true",
			},
			wantColor: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for key, value := range test.env {
				t.Setenv(key, value)
			}

			stdout := captureStdout(t, func() {
				err := run(test.args)
				assert.NilError(t, err)
			})

			assert.Equal(t, strings.Contains(stdout, "\x1b["), test.wantColor)
		})
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = writer

	t.Cleanup(func() {
		os.Stdout = original
		reader.Close()
		writer.Close()
	})

	fn()

	writer.Close()

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	return string(output)
}
