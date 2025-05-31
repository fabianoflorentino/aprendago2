package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "apd",
	Short: "Aprenda Go",
	Long:  "Aprenda Go com Aprendago é um projeto para ensinar Go de forma simples e prática.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
