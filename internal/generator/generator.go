package generator

import (
	"bytes"
	"fmt"

	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	tmplFS "github.com/AungKyawPhyo1142/create-express-cli/internal/templates"
)

type Options struct {
	ProjectName string
	Template    string
	NoInstall   bool
	DryRun      bool
}

func Generate(opts Options) error {
	srcPath := opts.Template
	destPath := opts.ProjectName

	if opts.DryRun {
		fmt.Println("[Dry Run] Would create project at: ", destPath)
		return nil
	}

	// create target dir
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// copy files from template
	return fs.WalkDir(tmplFS.Templates, srcPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, _ := filepath.Rel(srcPath, path)
		target := filepath.Join(destPath, relPath)

		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}

		return renderOrCopyEmbeddedFile(path, target, opts)

	})

}

func renderOrCopyEmbeddedFile(src, dst string, opts Options) error {
	content, err := tmplFS.Templates.ReadFile(src)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), "{{") {
		tmpl, err := template.New(filepath.Base(src)).Parse(string(content))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, opts); err != nil {
			return fmt.Errorf("template execute error in %s: %w", src, err)
		}
		return os.WriteFile(dst, buf.Bytes(), 0644)
	}
	return os.WriteFile(dst, content, 0644)
}

/*
func renderOrCopyFile(src, dst string, opts Options) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// check if file contains Go template syntax
	if strings.Contains(string(content), "{{") {
		tmpl, err := template.New(filepath.Base(src)).Parse(string(content))
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, opts); err != nil {
			return fmt.Errorf("template execute error in %s: %w", src, err)
		}
		return os.WriteFile(dst, buf.Bytes(), 0644)

	}
	return copyFile(src, dst)

}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err

}
*/
