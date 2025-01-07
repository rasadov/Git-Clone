package cmd

import (
	"GitClone/cmd/utils"
	"fmt"
	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "Compute object ID and optionally creates a blob from a file",
	Long:  `hashObject is used to compute object ID and optionally creates a blob from a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePath, _ := cmd.Flags().GetString("write")
		objectType, _ := cmd.Flags().GetString("type")
		path := args[0]
		fmt.Println(utils.CreateObject(path, writePath, objectType))
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
	hashObjectCmd.Flags().StringP("type", "t", "", "Specify the type")
	hashObjectCmd.Flags().StringP("write", "w", "", "Actually write the object into the database")

}
