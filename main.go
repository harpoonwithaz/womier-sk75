package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"womier-sk75/frontend"
)

func main() {
	p := tea.NewProgram(frontend.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}
