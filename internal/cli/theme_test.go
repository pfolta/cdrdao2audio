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
	"testing"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFangColorScheme(t *testing.T) {
	tests := []struct {
		name      string
		lightDark lipgloss.LightDarkFunc
		want      fang.ColorScheme
	}{
		{
			name: "light",
			lightDark: func(light, dark color.Color) color.Color {
				return light
			},
			want: fang.ColorScheme{
				Base:         lipgloss.NoColor{},
				Title:        lipgloss.Blue,
				Description:  lipgloss.NoColor{},
				Comment:      lipgloss.Yellow,
				Flag:         lipgloss.Magenta,
				FlagDefault:  lipgloss.BrightMagenta,
				Command:      lipgloss.Cyan,
				QuotedString: lipgloss.Green,
				Argument:     lipgloss.NoColor{},
				Help:         lipgloss.NoColor{},
				Dash:         lipgloss.NoColor{},
				ErrorHeader:  [2]color.Color{lipgloss.White, lipgloss.Red},
				ErrorDetails: lipgloss.Red,
			},
		},
		{
			name: "dark",
			lightDark: func(light, dark color.Color) color.Color {
				return dark
			},
			want: fang.ColorScheme{
				Base:         lipgloss.NoColor{},
				Title:        lipgloss.Blue,
				Description:  lipgloss.NoColor{},
				Comment:      lipgloss.Yellow,
				Flag:         lipgloss.Magenta,
				FlagDefault:  lipgloss.BrightMagenta,
				Command:      lipgloss.Cyan,
				QuotedString: lipgloss.Green,
				Argument:     lipgloss.NoColor{},
				Help:         lipgloss.NoColor{},
				Dash:         lipgloss.NoColor{},
				ErrorHeader:  [2]color.Color{lipgloss.White, lipgloss.Red},
				ErrorDetails: lipgloss.Red,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			theme := FangColorScheme(test.lightDark)

			assert.Assert(t, is.Equal(theme.Base, test.want.Base))
			assert.Assert(t, is.Equal(theme.Title, test.want.Title))
			assert.Assert(t, is.Equal(theme.Description, test.want.Description))
			assert.Assert(t, is.Equal(theme.Comment, test.want.Comment))
			assert.Assert(t, is.Equal(theme.Flag, test.want.Flag))
			assert.Assert(t, is.Equal(theme.FlagDefault, test.want.FlagDefault))
			assert.Assert(t, is.Equal(theme.Command, test.want.Command))
			assert.Assert(t, is.Equal(theme.QuotedString, test.want.QuotedString))
			assert.Assert(t, is.Equal(theme.Argument, test.want.Argument))
			assert.Assert(t, is.Equal(theme.Help, test.want.Help))
			assert.Assert(t, is.Equal(theme.Dash, test.want.Dash))
			assert.Assert(t, is.Equal(theme.ErrorHeader, test.want.ErrorHeader))
			assert.Assert(t, is.Equal(theme.ErrorDetails, test.want.ErrorDetails))
		})
	}
}
