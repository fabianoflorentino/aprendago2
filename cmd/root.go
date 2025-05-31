package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "aprendago",
	Short:   "Aprenda Go",
	Long:    "Aprenda Go com Aprendago é um projeto para ensinar Go de forma simples e prática.",
	Version: os.Getenv("APP_VERSION"),
	Run:     func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
