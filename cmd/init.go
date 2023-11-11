package cmd

import (
	"fmt"
	"os"

	"github.com/gursheyss/scaffoldb/cmd/docker"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new database in Docker with optional AI-generated schema",
	Long: `The 'init' command in StreamlineDB is used to quickly set up a new database within a Docker container. 
It provides an interactive interface to choose the type of database (e.g., PostgreSQL, MySQL) and configure basic settings. 
Additionally, it leverages GPT-4 AI to optionally generate a database schema, streamlining the setup process for new or existing projects. 
This command is ideal for developers who need to quickly scaffold databases for development, testing, or production environments.

Example:

streamlinedb init

This will start the process of setting up a new database, offering choices for database type, version, and configuration. 
If opted for, it will also generate an initial schema based on AI suggestions.`,

	Run: func(cmd *cobra.Command, args []string) {
		err := docker.CheckDockerRunning()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initCmd.Flags().StringP("name", "n", "", "Name of project")
}
