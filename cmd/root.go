package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-express-cli",
	Short: "A CLI tool to scaffold Express.js applications quickly.",
	Long: `create-express-cli is a tool for quickly creating Node.js Express
applications with pre-configured middlewares, JWT authentication, and example endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to create-express-cli! Use 'create-express-cli create <project-name>' to start.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
