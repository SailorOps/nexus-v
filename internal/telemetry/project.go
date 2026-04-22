package telemetry

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ProjectSink struct{}

func (p *ProjectSink) Record(ev Event) {
	if ev.ProjectDir == "" {
		return
	}

	path := filepath.Join(ev.ProjectDir, ".nexusv-usage.json")

	out, _ := json.MarshalIndent(ev, "", "  ")
	_ = os.WriteFile(path, out, 0o644)
}
