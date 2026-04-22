package telemetry

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type LocalData struct {
	TemplateCounts map[string]int `json:"templateCounts"`
	TotalRuns      int            `json:"totalRuns"`
}

type LocalSink struct{}

func (l *LocalSink) Record(ev Event) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(home, ".nexusv-telemetry.json")

	var data LocalData
	data.TemplateCounts = map[string]int{}

	if b, err := os.ReadFile(path); err == nil {
		json.Unmarshal(b, &data)
	}

	data.TotalRuns++
	data.TemplateCounts[ev.Template]++

	out, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(path, out, 0o644)
}
