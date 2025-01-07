package cmd

import (
	"GitClone/cmd/controls"
	"github.com/spf13/cobra"
	"os"
)

var Settings map[string]string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GitClone",
	Short: "Remote version control system",
	Long: `GitClone is a distributed version control system. GitClone is a free and open source software
designed to handle everything from small to very large projects with speed and efficiency.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	Settings = controls.LoadConfig()
}
