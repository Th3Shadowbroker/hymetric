package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/th3shadowbroker/hymetric/internal/config"
)

func createInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new configuration with default values",
		Run: func(cmd *cobra.Command, args []string) {
			configFlag, _ := cmd.Flags().GetString("config")
			if err := config.LoadConfiguration(configFlag); err != nil {
				log.Fatalln("Could not load configuration for initialization")
			}

			if err := viper.SafeWriteConfigAs(configFlag); err != nil {
				if _, ok := errors.AsType[viper.ConfigFileAlreadyExistsError](err); ok {
					log.Printf("There already is a configuration file at %s", configFlag)
				} else {
					log.Panicf("Could not wirte config: %s", err)
				}
			} else {
				log.Println("Configuration file initialized")
			}
		},
	}
}
