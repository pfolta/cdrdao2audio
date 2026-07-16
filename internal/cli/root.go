package cli

import (
	"github.com/spf13/cobra"

	"github.com/pfolta/cdrdao2audio"
)

func NewRootCommand() *cobra.Command {
	metadata := cdrdao2audio.GetMetadata()

	cmd := &cobra.Command{
		Use:                metadata.Name,
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

	cmd.AddCommand(NewVersionCommand(metadata))

	return cmd
}
