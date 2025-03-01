package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/cmd/controls"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Display the contents of a file",
	Long: `cat-file is used to display the contents of a file.
This command is used to display the contents of a file in the git repository.`,

	Run: func(cmd *cobra.Command, args []string) {
		res, err := controls.ReadObject(args[0], args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}
