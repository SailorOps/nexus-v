package templates

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed files/**
var templateFS embed.FS

type Context struct {
	Name        string
	Identifier  string
	Description string
	CommandName string
	Template    string
	Force       bool
	DryRun      bool
}

// ListTemplates returns all template variants in files/
func ListTemplates() ([]string, error) {
	entries, err := templateFS.ReadDir("files")
	if err != nil {
		return nil, err
	}

	var out []string
	for _, e := range entries {
		if e.IsDir() {
			out = append(out, e.Name())
		}
	}
	return out, nil
}

func GenerateProject(ctx Context, targetDir string) error {
	if ctx.Template == "" {
		ctx.Template = "default"
	}

	searchRoots := []string{
		filepath.Join("files", ctx.Template),
		filepath.Join("files", "default"),
	}

	if _, err := templateFS.ReadDir(searchRoots[0]); err != nil {
		return fmt.Errorf("unknown template variant: %s", ctx.Template)
	}

	seen := map[string]bool{}

	for _, root := range searchRoots {
		err := fsWalk(root, func(path string, isDir bool) error {
			rel := strings.TrimPrefix(path, root)
			rel = strings.TrimPrefix(rel, string(os.PathSeparator))

			if seen[rel] {
				return nil
			}
			seen[rel] = true

			outPath := filepath.Join(targetDir, rel)

			if isDir {
				if ctx.DryRun {
					fmt.Println("[dir]  ", outPath)
					return nil
				}
				return os.MkdirAll(outPath, 0o755)
			}

			outPath = strings.TrimSuffix(outPath, ".tmpl")

			if ctx.DryRun {
				fmt.Println("[file] ", outPath)
				return nil
			}

			if !ctx.Force {
				if _, err := os.Stat(outPath); err == nil {
					return fmt.Errorf(
						"refusing to overwrite existing file: %s (use --force to override)",
						outPath,
					)
				}
			}

			return renderFile(path, outPath, ctx)
		})

		if err != nil {
			return err
		}
	}

	return nil
}
