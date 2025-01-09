package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/cmd/controls"
)

// commitTreeCmd represents the commitTree command
var commitTreeCmd = &cobra.Command{
	Use:   "commit-tree",
	Short: "Record changes to the repository",
	Long:  `The git commit-tree command is used to record changes to the repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		author := Settings["author"]
		email := Settings["email"]
		if author == "" || email == "" {
			fmt.Println("Author name and email are not set. Please run config command")
			return
		}
		message := cmd.Flag("message").Value.String()
		parent := cmd.Flag("parent").Value.String()
		fmt.Println(controls.CreateCommit(args[0], parent, message, author, email))
	},
}

func init() {
	rootCmd.AddCommand(commitTreeCmd)
	commitTreeCmd.Flags().StringP("message", "m", "", "Commit message")
	commitTreeCmd.Flags().StringP("parent", "p", "", "Parent commit")
}
