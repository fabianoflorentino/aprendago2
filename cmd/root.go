package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     rootCmdUse,
	Short:   rootCmdShort,
	Long:    rootCmdLong,
	Version: rootVersion,
	Run:     func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
