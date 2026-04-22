package telemetry

import "fmt"

type SessionSink struct {
	Events []Event
}

func (s *SessionSink) Record(ev Event) {
	s.Events = append(s.Events, ev)
}

func (s *SessionSink) Dump() {
	for _, e := range s.Events {
		fmt.Printf("[session] template=%s dry=%v force=%v\n",
			e.Template, e.DryRun, e.Force)
	}
}
