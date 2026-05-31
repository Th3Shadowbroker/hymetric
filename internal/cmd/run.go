package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Listen for and serve HTTP requests",
		Run: func(cmd *cobra.Command, args []string) {
			panic(fmt.Errorf("Not implemented!"))
		},
	}
}
