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
	"context"
	"os"

	"charm.land/fang/v2"

	"github.com/pfolta/cdrdao2audio/internal/cli/command"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(args []string) error {
	opts := []fang.Option{
		fang.WithColorSchemeFunc(fang.AnsiColorScheme),
		fang.WithoutManpage(),

		// Use `version` subcommand instead of `--version`/`-v` flag.
		fang.WithoutVersion(),
	}

	ctx := context.Background()

	cmd := command.NewRootCommand()
	cmd.SetArgs(args)

	// Parse flags early, so the color mode is configured before running Cobra.
	// Ignore any parsing errors at this stage. Cobra will parse the flags again
	// when running `Execute()` and report any parsing errors.
	_ = cmd.ParseFlags(args)

	return fang.Execute(ctx, cmd, opts...)
}
