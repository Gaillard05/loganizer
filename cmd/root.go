package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "loganizer",
	Short: "Analyseur de fichiers de logs",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
