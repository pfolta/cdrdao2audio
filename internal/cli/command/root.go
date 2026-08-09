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

package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pfolta/cdrdao2audio"
	"github.com/pfolta/cdrdao2audio/internal/cli"
)

type rootOptions struct {
	colorMode cli.ColorMode
}

// ErrNoCommand indicates that the CLI was invoked without a subcommand.
var ErrNoCommand = errors.New("no command specified")

func NewRootCommand() *cobra.Command {
	appInfo := cdrdao2audio.GetAppInfo()

	opts := rootOptions{
		// Initialize with "auto" so that Cobra's help text includes the default
		// value for `--color`. However, this doesn't trigger the side effect in
		// `*colorMode.Set()`. In this case, this is desired as the environment
		// configuration should only be applied if a `--color` flag was actually
		// passed. This means that without the flag the current environment
		// variables still take effect.
		colorMode: cli.ColorAuto,
	}

	cmd := &cobra.Command{
		Use:                appInfo.Name,
		Short:              "Convert a cdrdao dump to individual audio tracks",
		DisableSuggestions: true,
		SilenceErrors:      true,
		SilenceUsage:       true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		// By default, if neither Run nor RunE is defined, Cobra displays the
		// help text and exits successfully. Treat invoking the root command as
		// an error instead.
		RunE: func(cmd *cobra.Command, _ []string) error {
			_ = cmd.Help()
			return ErrNoCommand
		},
	}

	pflags := cmd.PersistentFlags()
	pflags.Var(
		&opts.colorMode,
		"color",
		fmt.Sprintf("colored output: [%s]", strings.Join(cli.ColorModes(), "|")),
	)

	cmd.AddCommand(NewVersionCommand(&appInfo))

	return cmd
}
