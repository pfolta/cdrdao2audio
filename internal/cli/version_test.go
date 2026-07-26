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

package cli

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/pfolta/cdrdao2audio"
)

var testAppInfo = cdrdao2audio.AppInfo{
	Name:      "cdrdao2audio",
	Version:   "v1.4.2-test",
	BuildDate: "2026-07-15T22:29:10Z",
	License:   "MIT",
	OS:        "darwin",
	Arch:      "arm64",
}

const (
	expectedShortVersionText = `1.4.2-test
`

	expectedVersionText = `cdrdao2audio version v1.4.2-test-darwin-arm64 (2026-07-15T22:29:10Z)

MIT
`

	expectedVersionJSON = `{
    "name": "cdrdao2audio",
    "version": "v1.4.2-test",
    "buildDate": "2026-07-15T22:29:10Z",
    "license": "MIT",
    "os": "darwin",
    "arch": "arm64"
}
`

	expectedShortVersionJSON = `{
    "version": "1.4.2-test"
}
`
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expected    string
		expectedErr string
	}{
		{
			name:     "default text output",
			args:     nil,
			expected: expectedVersionText,
		},
		{
			name:     "short version text output",
			args:     []string{"--short"},
			expected: expectedShortVersionText,
		},
		{
			name:     "json output",
			args:     []string{"--format", "json"},
			expected: expectedVersionJSON,
		},
		{
			name:     "short json output",
			args:     []string{"--short", "--format", "json"},
			expected: expectedShortVersionJSON,
		},
		{
			name:        "unknown output format",
			args:        []string{"--format", "xml"},
			expectedErr: "unknown format",
		},
		{
			name:        "unknown flag",
			args:        []string{"--unknown"},
			expectedErr: "unknown flag",
		},
		{
			name:        "unknown command",
			args:        []string{"unknown"},
			expectedErr: "unknown command",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := new(bytes.Buffer)

			cmd := NewVersionCommand(testAppInfo)
			cmd.SetOut(out)
			cmd.SetArgs(test.args)

			err := cmd.Execute()

			if test.expectedErr == "" {
				assert.NilError(t, err)
				assert.Equal(t, out.String(), test.expected)
			} else {
				assert.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}
