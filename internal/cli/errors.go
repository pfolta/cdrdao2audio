package cli

import "errors"

// ErrNoCommand indicates that the CLI was invoked without a subcommand.
var ErrNoCommand = errors.New("no command specified")
