package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/th3shadowbroker/hymetric/internal/api"
	"github.com/th3shadowbroker/hymetric/internal/config"
)

func createServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Listen for and serve HTTP requests",
		Run: func(cmd *cobra.Command, args []string) {
			configFlag, _ := cmd.Flags().GetString("config")
			if err := config.LoadConfiguration(configFlag); err != nil {
				log.Fatalf("Could not load configuration: %s", err)
			}

			api.Listen(config.Active.Http.Address)
		},
	}
}
