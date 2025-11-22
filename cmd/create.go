package cmd

import (
	"fmt"

	"github.com/AungKyawPhyo1142/create-express-cli/internal/generator"
	"github.com/AungKyawPhyo1142/create-express-cli/internal/setup"
	"github.com/AungKyawPhyo1142/create-express-cli/internal/tui"
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
		// Check if any flags were explicitly set (excluding project name as positional arg)
		flagsChanged := cmd.Flags().Changed("template") ||
			cmd.Flags().Changed("pm") ||
			cmd.Flags().Changed("no-install") ||
			cmd.Flags().Changed("git") ||
			cmd.Flags().Changed("dry-run")

		// If no flags were set, use TUI
		if !flagsChanged {
			result, err := tui.Run()
			if err != nil {
				fmt.Println("Error running interactive UI:", err)
				return
			}

			// Use TUI results
			projectName = result.ProjectName
			if projectName == "" {
				fmt.Println("Project creation cancelled.")
				return
			}

			// Use progress bar for project creation
			err = tui.RunProgressWithCallback(func(updateProgress tui.ProgressCallback) error {
				// Step 1: Generate project files
				updateProgress(tui.StepGenerating, "Generating project files...", 0.1)
				opts := generator.Options{
					ProjectName: projectName,
					Template:    "express-ts", // Always TypeScript
					NoInstall:   noInstall,
					DryRun:      dryRun,
				}

				if err := generator.Generate(opts); err != nil {
					return fmt.Errorf("error creating project: %w", err)
				}
				updateProgress(tui.StepGenerating, "Project files generated!", 0.4)

				// Step 2: Setup (install deps, init git)
				setupOpts := setup.Options{
					ProjectPath:    projectName,
					PackageManager: packageMgr,
					InstallDeps:    !noInstall,
					InitGit:        !initGit,
				}

				if setupOpts.InstallDeps {
					updateProgress(tui.StepInstalling, "Installing dependencies...", 0.5)
				}
				if setupOpts.InitGit {
					updateProgress(tui.StepInitializingGit, "Initializing Git repository...", 0.9)
				}

				if err := setup.Run(setupOpts); err != nil {
					return fmt.Errorf("error during project setup: %w", err)
				}

				updateProgress(tui.StepComplete, "Project created successfully!", 1.0)
				return nil
			})

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			return
		}

		// Flag-based flow (backward compatible)
		// if project-name was provided as an argument, use it
		if len(args) > 0 {
			projectName = args[0]
		}

		if projectName == "" {
			fmt.Println("Please provide a project name. Usage: create-express-cli create <project-name>")
			return
		}

		// TypeScript is now the default and only option
		templateName := "express-ts"

		// Use progress bar for project creation
		var err error
		err = tui.RunProgressWithCallback(func(updateProgress tui.ProgressCallback) error {
			// Step 1: Generate project files
			updateProgress(tui.StepGenerating, "Generating project files...", 0.1)
			opts := generator.Options{
				ProjectName: projectName,
				Template:    templateName,
				NoInstall:   noInstall,
				DryRun:      dryRun,
			}

			if err := generator.Generate(opts); err != nil {
				return fmt.Errorf("error creating project: %w", err)
			}
			updateProgress(tui.StepGenerating, "Project files generated!", 0.4)

			// Step 2: Setup (install deps, init git)
			setupOpts := setup.Options{
				ProjectPath:    projectName,
				PackageManager: packageMgr,
				InstallDeps:    !noInstall,
				InitGit:        initGit,
			}

			if setupOpts.InstallDeps {
				updateProgress(tui.StepInstalling, "Installing dependencies...", 0.5)
			}
			if setupOpts.InitGit {
				updateProgress(tui.StepInitializingGit, "Initializing Git repository...", 0.9)
			}

			if err := setup.Run(setupOpts); err != nil {
				return fmt.Errorf("error during project setup: %w", err)
			}

			updateProgress(tui.StepComplete, "Project created successfully!", 1.0)
			return nil
		})

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&template, "template", "t", "express-ts", "Template to use (default: express-ts)")
	createCmd.Flags().StringVarP(&packageMgr, "pm", "p", "npm", "Package manager to use (npm, yarn, pnpm)")
	createCmd.Flags().BoolVarP(&noInstall, "no-install", "n", false, "Skip installation of dependencies")
	createCmd.Flags().BoolVarP(&initGit, "git", "g", true, "Initialize a Git repository")
	createCmd.Flags().BoolVar(&typescript, "typescript", true, "Use TypeScript (default: true)")
	createCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Simulate the creation without making changes")
}
