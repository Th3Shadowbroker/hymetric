package cmd

import "github.com/spf13/cobra"

func CreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hymetric",
		Short: "A tiny service for displaying Hypixel timers on a LaMetric smart-clock.",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cmd.PersistentFlags().StringP("config", "c", "config.yml", "path to configuration file")
	cmd.MarkFlagFilename("config", "yml", "yaml")

	cmd.AddCommand(createInitCmd(), createServeCmd())

	return cmd
}
