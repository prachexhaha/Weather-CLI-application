package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Weather generator ",
		Long:  `Weather generator API`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
