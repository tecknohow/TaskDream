package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskdream",
	Short: "TaskDream - A task and project management application",
	Long: `TaskDream is a powerful task and project management application
inspired by Vikunja's architecture, with features from Super Productivity
and Tududi.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(webCmd)
}
