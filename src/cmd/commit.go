package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	controls2 "src/cmd/controls"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Long:  `The git commit command is used to record changes to the repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := cmd.Flag("message").Value.String()
		parentHash := controls2.GetHead("main")
		treeHash := controls2.WriteTree()
		author := Settings["author"]
		email := Settings["email"]
		commitHash := controls2.CreateCommit(treeHash, parentHash, message, author, email)
		controls2.UpdateHead("main", commitHash)
		fmt.Println("Commit created with hash:", commitHash)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringP("message", "m", "", "Commit message")
}
