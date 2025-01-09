package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/cmd/controls"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Adding or removing remote repository",
	Long:  `The git remote command is used to add or remove remote repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remote called")
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			controls.SetRemote(args[0], args[1])
			controls.SaveRemotes()
		},
	})

	remoteCmd.AddCommand(&cobra.Command{
		Use:   "remove",
		Short: "Remove remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			controls.RemoveRemote(args[0])
			controls.SaveRemotes()
		},
	})

	remoteCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			controls.GetRemote(args[0])
		},
	})
}
