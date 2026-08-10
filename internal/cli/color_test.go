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
	"os"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestColorModes(t *testing.T) {
	modes := ColorModes()
	want := []string{"never", "auto", "always"}

	assert.Assert(t, is.Equal(len(modes), len(want)))

	for _, mode := range modes {
		assert.Assert(t, is.Contains(want, mode))
	}
}

func TestColorModeString(t *testing.T) {
	for _, mode := range ColorModes() {
		t.Run(mode, func(t *testing.T) {
			assert.Assert(t, is.Equal(ColorMode(mode).String(), mode))
		})
	}
}

func TestColorModeType(t *testing.T) {
	for _, mode := range ColorModes() {
		t.Run(mode, func(t *testing.T) {
			assert.Assert(t, is.Equal(ColorMode(mode).Type(), "string"))
		})
	}
}

func TestColorModeSet(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantMode  ColorMode
		wantEnv   map[string]string
		wantUnset []string
	}{
		{
			name:     "never",
			input:    "never",
			wantMode: ColorNever,
			wantEnv: map[string]string{
				"NO_COLOR": "true",
			},
			wantUnset: []string{
				"CLICOLOR",
				"CLICOLOR_FORCE",
			},
		},
		{
			name:     "auto",
			input:    "auto",
			wantMode: ColorAuto,
			wantEnv: map[string]string{
				"CLICOLOR": "true",
			},
			wantUnset: []string{
				"NO_COLOR",
				"CLICOLOR_FORCE",
			},
		},
		{
			name:     "always",
			input:    "always",
			wantMode: ColorAlways,
			wantEnv: map[string]string{
				"CLICOLOR":       "true",
				"CLICOLOR_FORCE": "true",
			},
			wantUnset: []string{
				"NO_COLOR",
			},
		},
		{
			name:     "uppercase",
			input:    "ALWAYS",
			wantMode: ColorAlways,
			wantEnv: map[string]string{
				"CLICOLOR":       "true",
				"CLICOLOR_FORCE": "true",
			},
			wantUnset: []string{
				"NO_COLOR",
			},
		},
		{
			name:     "mixed case",
			input:    "AuTo",
			wantMode: ColorAuto,
			wantEnv: map[string]string{
				"CLICOLOR": "true",
			},
			wantUnset: []string{
				"NO_COLOR",
				"CLICOLOR_FORCE",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Cleanup(resetEnv)

			var mode ColorMode
			err := mode.Set(test.input)
			assert.NilError(t, err)
			assert.Assert(t, is.Equal(mode, test.wantMode))

			for key, want := range test.wantEnv {
				assert.Assert(t, is.Equal(os.Getenv(key), want))
			}

			for _, key := range test.wantUnset {
				_, ok := os.LookupEnv(key)
				assert.Assert(t, is.Equal(ok, false))
			}
		})
	}
}

func TestColorModeSetInvalid(t *testing.T) {
	t.Cleanup(resetEnv)

	t.Setenv("NO_COLOR", "original")
	t.Setenv("CLICOLOR", "original")
	t.Setenv("CLICOLOR_FORCE", "original")

	colorMode := ColorAuto

	err := colorMode.Set("invalid")
	assert.Assert(t, is.Error(err, "must be one of [never|auto|always]"))
	assert.Assert(t, is.Equal(colorMode, ColorAuto))

	assert.Assert(t, is.Equal(os.Getenv("NO_COLOR"), "original"))
	assert.Assert(t, is.Equal(os.Getenv("CLICOLOR"), "original"))
	assert.Assert(t, is.Equal(os.Getenv("CLICOLOR_FORCE"), "original"))
}

func resetEnv() {
	os.Unsetenv("NO_COLOR")
	os.Unsetenv("CLICOLOR")
	os.Unsetenv("CLICOLOR_FORCE")
}
