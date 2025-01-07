/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"GitClone/cmd/controls"
	"fmt"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Long:  `The git commit command is used to record changes to the repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := cmd.Flag("message").Value.String()
		parentHash := controls.GetHead("main")
		treeHash := controls.WriteTree()
		author := Settings["author"]
		email := Settings["email"]
		commitHash := controls.CreateCommit(treeHash, parentHash, message, author, email)
		controls.UpdateHead("main", commitHash)
		fmt.Println("Commit created with hash:", commitHash)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringP("message", "m", "", "Commit message")
}
