package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/cmd/controls"
)

// writeTreeCmd represents the writeTree command
var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "Create a tree object",
	Long: `The git write-tree command creates a tree object from the current state
			of the "staging area". The staging area is a place where changes go when you run git add.
			In this challenge we won't implement a staging area,
			we'll just assume that all files in the working directory are staged.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(controls.WriteTree())
	},
}

func init() {
	rootCmd.AddCommand(writeTreeCmd)
}
