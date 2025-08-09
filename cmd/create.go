package cmd

import (
	"fmt"

	"github.com/AungKyawPhyo1142/create-express-cli/internal/generator"
	"github.com/AungKyawPhyo1142/create-express-cli/internal/setup"
	"github.com/spf13/cobra"
)

var (
	projectName string
	template    string
	packageMgr  string
	noInstall   bool
	initGit     bool
	dryRun      bool
	typescript  bool
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

		// Determine the correct template based on flags.
		// The --typescript flag takes precedence.
		var templateName string
		if typescript {
			templateName = "express-ts"
		} else {
			templateName = template
		}

		opts := generator.Options{
			ProjectName: projectName,
			Template:    templateName,
			NoInstall:   noInstall,
			DryRun:      dryRun,
		}

		if err := generator.Generate(opts); err != nil {
			fmt.Println("Error creating project:", err)
			return
		}

		if err := setup.Run(setup.Options{
			ProjectPath:    projectName,
			PackageManager: packageMgr,
			InstallDeps:    !noInstall,
			InitGit:        !initGit,
		}); err != nil {
			fmt.Println("Error during project setup:", err)
			fmt.Println("Project files were created, but dependency installation failed. Please run 'npm install' manually.")
			return
		}

		fmt.Println("Project created successfully")

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&template, "template", "t", "express-basic", "Template to use for JavaScript projects (e.g., express-basic)")
	createCmd.Flags().StringVarP(&packageMgr, "pm", "p", "npm", "Package manager to use (npm, yarn, pnpm)")
	createCmd.Flags().BoolVarP(&noInstall, "no-install", "n", false, "Skip installation of dependencies")
	createCmd.Flags().BoolVarP(&initGit, "git", "g", true, "Initialize a Git repository")
	createCmd.Flags().BoolVar(&typescript, "typescript", false, "Use TypeScript instead of JavaScript")
	createCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Simulate the creation without making changes")
}
