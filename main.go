package main

import (
	"fmt"
	"os"

	"womier-sk75/frontend"
	"womier-sk75/keyboard"

	tea "charm.land/bubbletea/v2"
)

func main() {
	kb := &keyboard.Keyboard{}

	kb.DetectKeyboard()
	err := kb.Connect()
	if err != nil {
		fmt.Printf("Error connecting to keyboard: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(frontend.InitialModel(kb))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}
