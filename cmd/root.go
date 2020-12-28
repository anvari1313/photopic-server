package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCMD = &cobra.Command{
		Use:   "photopic",
		Short: "A static photo album application",
		Long:  `Photopic is a simple and easy to run static photo album web application.`,
	}
)

func init() {
	cobra.OnInitialize(configure)

	rootCMD.AddCommand(serveCMD)
}

func configure() {
	// TODO: Put some configuration codes here
}

// Execute executes the root command.
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Printf("failed to execute root command: %s\n", err.Error())
		os.Exit(1)
	}
}
