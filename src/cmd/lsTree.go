package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/cmd/controls"
)

// lsTreeCmd represents the lsTree command
var lsTreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "List the contents of a tree object",
	Long:  `ls-tree is used to list the contents of a tree object.`,
	Run: func(cmd *cobra.Command, args []string) {
		treeHash := args[0]
		res, err := controls.ReadObject("p", treeHash)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(lsTreeCmd)
}
