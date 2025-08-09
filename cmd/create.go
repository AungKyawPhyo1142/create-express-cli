package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	projectName string
	template    string
	packageMgr  string
	noInstall   bool
	dryRun      bool
)

var createCmd = &cobra.Command{
	Use:   "create <project-name>",
	Short: "Create a new Express.js application",
	Long: `Scaffold a new Express.js application with pre-configured middlewares,
JWT authentication, and example endpoints. Example:

express-cli create myapp --template express-jwt --pm npm`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// if project-name was provided as an argument, use it
		if len(args) > 0 {
			projectName = args[0]
		}

		if projectName == "" {
			fmt.Println("Please provide a project name. Usage: create-express-cli create <project-name>")
			return
		}

		fmt.Println("Creating project with:")
		fmt.Println("  Name       :", projectName)
		fmt.Println("  Template   :", template)
		fmt.Println("  PackageMgr :", packageMgr)
		fmt.Println("  No Install :", noInstall)
		fmt.Println("  Dry Run    :", dryRun)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&template, "template", "t", "express-jwt", "Template to use (express-basic, express-jwt, etc.)")
	createCmd.Flags().StringVarP(&packageMgr, "pm", "p", "npm", "Package manager to use (npm, yarn, pnpm)")
	createCmd.Flags().BoolVarP(&noInstall, "no-install", "n", false, "Skip installation of dependencies")
	createCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Simulate the creation without making changes")
}
