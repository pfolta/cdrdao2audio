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
	"image/color"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
)

func FangColorScheme(lightDark lipgloss.LightDarkFunc) fang.ColorScheme {
	return fang.ColorScheme{
		Base:         lipgloss.NoColor{},
		Title:        lipgloss.Blue,
		Description:  lipgloss.NoColor{},
		Comment:      lightDark(lipgloss.BrightWhite, lipgloss.BrightBlack),
		Flag:         lipgloss.Magenta,
		FlagDefault:  lipgloss.BrightMagenta,
		Command:      lipgloss.Cyan,
		QuotedString: lipgloss.Green,
		Argument:     lipgloss.NoColor{},
		Help:         lipgloss.NoColor{},
		Dash:         lipgloss.NoColor{},
		ErrorHeader:  [2]color.Color{lipgloss.White, lipgloss.Red},
		ErrorDetails: lipgloss.Red,
	}
}
