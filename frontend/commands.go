package frontend

import tea "charm.land/bubbletea/v2"

// Custom message types for asynchronous keyboard operations
type keyboardOpMsg struct {
	success bool
	err     error
}

// Placeholder tea.Cmd functions for hardware interaction
func setBrightnessCmd(level int) tea.Cmd {
	return func() tea.Msg {
		// TODO: Call kb.SetRGB(PropBrightness, byte(level)) here
		return keyboardOpMsg{success: true}
	}
}

func setColorCmd(hue int) tea.Cmd {
	return func() tea.Msg {
		// TODO: Call kb.SetRGB(PropColor, byte(hue)) here
		return keyboardOpMsg{success: true}
	}
}
