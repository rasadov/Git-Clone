package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize git repository",
	Long: `Initialize git repository in the current directory.
This command is used to create a new git repository in the current directory.`,

	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize() {
	for _, dir := range []string{GitDir, GitDir + "/objects", GitDir + "/refs", GitDir + "/heads"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Sprintf("Error creating directory: %s\n", dir)
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(GitDir+"/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}
	os.Create(GitDir + "/config")
	fmt.Println("Initialized git directory")
}
