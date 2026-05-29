package main

import (
	"log"

	"github.com/th3shadowbroker/hymetric/internal/cmd"
)

func main() {
	rootCmd := cmd.CreateRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not execute command: %s", err)
	}
}
