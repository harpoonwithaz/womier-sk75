// Send commands to keyboard backend

package frontend

import (
	"womier-sk75/keyboard"

	tea "charm.land/bubbletea/v2"
)

// Custom message types for asynchronous keyboard operations
type keyboardOpMsg struct {
	success bool
	err     error
}

// Placeholder tea.Cmd functions for hardware interaction
func (m model) setBrightnessCmd(level int) tea.Cmd {
	return func() tea.Msg {
		// TODO: Call kb.SetRGB(PropBrightness, byte(level)) here
		m.kb.SetRGB(keyboard.PropBrightness, []byte{byte(level)})
		return keyboardOpMsg{success: true}
	}
}

func (m model) setEffectCmd(value int) tea.Cmd {
	return func() tea.Msg {
		m.kb.SetRGB(keyboard.PropEffect, []byte{byte(value)})
		return keyboardOpMsg{success: true}
	}
}

func (m model) setColorCmd(hue, sat int) tea.Cmd {
	return func() tea.Msg {
		// TODO: Call kb.SetRGB(PropColor, byte(hue)) here
		values := []byte{byte(hue), byte(sat)}
		m.kb.SetRGB(keyboard.PropColor, values)
		return keyboardOpMsg{success: true}
	}
}
