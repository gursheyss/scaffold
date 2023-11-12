package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/gursheyss/scaffoldb/cmd/docker"
	"github.com/spf13/cobra"
)

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

var (
	logoStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)
	tipMsgStyle          = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FF8C00")).Italic(true)
	endingMsgStyle       = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FF7F50")).Bold(true)
	allowedDatabaseTypes = []string{"mysql", "postgres", "sqlite", "mongo"}
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Scaffold different databases with Docker",
	Long:  "StreamlineDB's 'init' command is a CLI tool for quick database setup in Docker. It offers interactive selection, configuration, and optional AI-based schema generation. Ideal for rapid database scaffolding.",

	Run: func(cmd *cobra.Command, args []string) {
		err := docker.CheckDockerRunning()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		connectionString, err := docker.StartMySQLContainer("user", "pass", "test", 3306)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Connection string: %v", connectionString)
	},
}
