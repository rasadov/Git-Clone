/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GitClone/cmd/controls"
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set the author name and email",
	Long:  `The git config command is a convenience function that is used to set Git configuration values on a global or local project level.`,
	Run: func(cmd *cobra.Command, args []string) {
		author, _ := cmd.Flags().GetString("author")
		email, _ := cmd.Flags().GetString("email")
		if author != "" {
			Settings["author"] = author
		}
		if email != "" {
			Settings["email"] = email
		}
		controls.SaveConfig(Settings)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get the author name and email",
		Run: func(cmd *cobra.Command, args []string) {
			if Settings["author"] != "" {
				fmt.Println("Author: ", Settings["author"])
			} else {
				fmt.Println("Author is not set")
			}
			if Settings["email"] != "" {
				fmt.Println("Email: ", Settings["email"])
			} else {
				fmt.Println("Email is not set")
			}
		},
	})
	configCmd.Flags().StringP("author", "a", "", "Set the author name and email")
	configCmd.Flags().StringP("email", "e", "", "Set the author email")
}
