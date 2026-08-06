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
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pfolta/cdrdao2audio"
	"github.com/pfolta/cdrdao2audio/internal/formatter"
)

type versionOptions struct {
	short  bool
	format string
}

// NewVersionCommand creates a Cobra command that displays application version
// information.
func NewVersionCommand(appInfo *cdrdao2audio.AppInfo) *cobra.Command {
	opts := versionOptions{}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Example: "# Print human-readable version and build information:\n" +
			appInfo.Name + " version\n" +
			"\n" +
			"# Only print the version number:\n" +
			appInfo.Name + " version --short\n" +
			appInfo.Name + " version -s\n" +
			"\n" +
			"# Print the version and build information formatted as YAML:\n" +
			appInfo.Name + " version --format yaml\n" +
			appInfo.Name + " version -f yaml\n" +
			"\n" +
			"# Mix and match:\n" +
			appInfo.Name + " version --short --format json\n" +
			appInfo.Name + " version -s -f json\n",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runVersion(cmd.OutOrStdout(), appInfo, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(
		&opts.format,
		"format",
		"f",
		string(formatter.TEXT),
		fmt.Sprintf(
			"output format: [%s]",
			strings.Join(formatter.Formats(), "|"),
		),
	)
	flags.BoolVarP(&opts.short, "short", "s", false, "show version number only")

	return cmd
}

func runVersion(
	w io.Writer,
	appInfo *cdrdao2audio.AppInfo,
	opts versionOptions,
) error {
	f, err := formatter.NewFormatter(formatter.Format(opts.format))
	if err != nil {
		return err
	}

	if opts.short {
		return f.Write(w, appInfo.ShortVersion())
	}

	return f.Write(w, appInfo)
}
