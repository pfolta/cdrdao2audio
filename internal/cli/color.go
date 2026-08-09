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
	"fmt"
	"os"
	"strings"
)

// ColorMode specifies when colored output should be enabled.
type ColorMode string

const (
	// ColorNever disables colored output.
	ColorNever ColorMode = "never"

	// ColorAuto enables colored output without forcing it.
	// Fang will enable color when stdout is a TTY and disable it otherwise.
	ColorAuto ColorMode = "auto"

	// ColorAlways forces colored output.
	// This is useful when piping the output to a program that can display
	// color, such as `less`.
	ColorAlways ColorMode = "always"
)

// ColorModes returns the supported color modes.
func ColorModes() []string {
	return []string{string(ColorNever), string(ColorAuto), string(ColorAlways)}
}

func (colorMode ColorMode) String() string {
	return string(colorMode)
}

func (colorMode *ColorMode) Set(str string) error {
	color := ColorMode(strings.ToLower(str))

	switch color {
	case ColorNever:
		disableColor()
	case ColorAuto:
		autoColor()
	case ColorAlways:
		forceColor()
	default:
		return fmt.Errorf("must be one of [%s]", strings.Join(ColorModes(), "|"))
	}

	*colorMode = color
	return nil
}

func (ColorMode) Type() string {
	return "string"
}

func disableColor() {
	os.Setenv("NO_COLOR", "true")
	os.Unsetenv("CLICOLOR")
	os.Unsetenv("CLICOLOR_FORCE")
}

func autoColor() {
	os.Unsetenv("NO_COLOR")
	os.Setenv("CLICOLOR", "true")
	os.Unsetenv("CLICOLOR_FORCE")
}

func forceColor() {
	os.Unsetenv("NO_COLOR")
	os.Setenv("CLICOLOR", "true")
	os.Setenv("CLICOLOR_FORCE", "true")
}
