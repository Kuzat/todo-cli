package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple cli application for managing todos",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
}
