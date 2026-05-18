package frontend

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

// Model stores the state of the TUI application
type model struct {
	state        menuState
	adjustTarget adjustTarget
	mainMenuIdx  int
	rgbMenuIdx   int
	macroMenuIdx int

	// keyboard state variables
	brightness    int // Range: 0-9
	colorHue      int // Range: 0-255
	statusMessage string
}

func InitialModel() model {
	return model{
		state:        stateMainMenu,
		adjustTarget: adjustNone,
		brightness:   5,   // start mid-range
		colorHue:     120, // start color
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.state != stateAdjusting {
				return m, tea.Quit
			}
		}

		// Route input handling based on current UI state
		switch m.state {
		case stateMainMenu:
			return m.handleMainMenuInput(msg)
		case stateRGBMenu:
			return m.handleRGBMenuInput(msg)
		case stateMacroMenu:
			return m.handleMacroMenuInput(msg)
		case stateAdjusting:
			return m.handleAdjustmentInput(msg)
		}

	case keyboardOpMsg:
		if msg.err != nil {
			m.statusMessage = fmt.Sprintf("Error: %v", msg.err)
		} else {
			m.statusMessage = "Keyboard updated successfully."
		}
	}

	return m, cmd
}
