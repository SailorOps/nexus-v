package cli

import (
	"fmt"
	"time"
)

type Spinner struct {
	stop chan struct{}
}

func NewSpinner() *Spinner {
	return &Spinner{stop: make(chan struct{})}
}

func (s *Spinner) Start(msg string) {
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for {
			select {
			case <-s.stop:
				return
			default:
				fmt.Printf("\r%s %s", frames[i], msg)
				i = (i + 1) % len(frames)
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	close(s.stop)
	fmt.Print("\r")
}
