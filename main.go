package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"womier-sk75/frontend"
	"womier-sk75/keyboard"
)

func main() {
	kb := &keyboard.Keyboard{}

	kb.DetectKeyboard()
	err := kb.Connect()
	if err != nil {
		fmt.Printf("Error connecting to keyboard: %v\n", err)
		os.Exit(1)
	}

	defer kb.Disconnect()

	p := tea.NewProgram(frontend.InitialModel(kb))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v", err)
		os.Exit(1)
	}
}
