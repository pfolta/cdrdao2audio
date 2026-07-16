package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pfolta/cdrdao2audio/internal/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		if !errors.Is(err, cli.ErrNoCommand) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		os.Exit(1)
	}
}
