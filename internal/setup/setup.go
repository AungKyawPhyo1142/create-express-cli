package setup

import (
	"fmt"
	"os"
	"os/exec"
)

type Options struct {
	ProjectPath    string
	PackageManager string
	InstallDeps    bool
	InitGit        bool
}

func Run(opts Options) error {
	fmt.Println("Checking enviornments...")

	if err := checkNode(); err != nil {
		return err
	}

	pm := opts.PackageManager
	if pm == "" {
		detected, err := detectPackageManager()
		if err != nil {
			return err
		}
		pm = detected
	}

	fmt.Println("Using package manager: ", pm)

	if opts.InstallDeps {
		if err := installDependencies(pm, opts.ProjectPath); err != nil {
			return err
		}
	} else {
		fmt.Println("Skipping dependency installation")
	}

	if opts.InitGit {
		if err := initGit(opts.ProjectPath); err != nil {
			return err
		}
	}

	fmt.Println("Setup complete!")
	fmt.Printf("👉 Next steps:\n  cd %s\n  %s run dev\n", opts.ProjectPath, pm)
	return nil

}

func checkNode() error {
	if _, err := exec.LookPath("node"); err != nil {
		return fmt.Errorf("node.js is not installed or not in PATH")
	}
	return nil
}

func detectPackageManager() (string, error) {
	if _, err := exec.LookPath("npm"); err == nil {
		return "npm", nil
	}
	if _, err := exec.LookPath("yarn"); err == nil {
		return "yarn", nil
	}
	if _, err := exec.LookPath("pnpm"); err == nil {
		return "pnpm", nil
	}
	return "", fmt.Errorf("no supported package manager found (npm, pnpm, yarn)")

}

func installDependencies(pm, path string) error {
	fmt.Println("Installing dependencies...")
	cmd := exec.Command(pm, "install")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func initGit(path string) error {
	fmt.Println("Initializing Git repository...")
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	gitignore := []byte("node_modules\n.DS_Store\n.env\n")
	return os.WriteFile(path+"/.gitignore", gitignore, 0644)

}
