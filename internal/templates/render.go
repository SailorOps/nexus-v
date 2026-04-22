package templates

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func renderFile(srcPath, outPath string, ctx Context) error {
	data, err := templateFS.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template: %w", err)
	}

	tmpl, err := template.New(filepath.Base(srcPath)).Parse(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ctx); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	if ctx.DryRun {
		return nil
	}

	return os.WriteFile(outPath, buf.Bytes(), 0o644)
}
