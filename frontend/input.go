package frontend

import tea "charm.land/bubbletea/v2"

// Define states for TUI menu system
type menuState int

const (
	stateMainMenu menuState = iota
	stateRGBMenu
	stateMacroMenu
	stateAdjusting
)

// Define what property is being adjusted
type adjustTarget int

const (
	adjustNone adjustTarget = iota
	adjustBrightness
	adjustEffect
	adjustEffectSpeed
	adjustColor
)

// Input handler for Main Menu
func (m model) handleMainMenuInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.mainMenuIdx > 0 {
			m.mainMenuIdx--
		}
	case "down", "j":
		if m.mainMenuIdx < 1 {
			m.mainMenuIdx++
		}
	case "enter":
		if m.mainMenuIdx == 0 {
			m.state = stateRGBMenu
		} else {
			m.state = stateMacroMenu
		}
	}
	return m, nil
}

// Input handler for RGB Sub-Menu
func (m model) handleRGBMenuInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "b":
		m.state = stateMainMenu
		return m, nil
	case "up", "k":
		if m.rgbMenuIdx > 0 {
			m.rgbMenuIdx--
		}
	case "down", "j":
		if m.rgbMenuIdx < 3 {
			m.rgbMenuIdx++
		}
	case "enter":
		m.state = stateAdjusting

		m.adjustTarget = adjustTarget(m.rgbMenuIdx + 1)
	}
	return m, nil
}

// Input handler for Macro Sub-Menu
func (m model) handleMacroMenuInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if msg.String() == "esc" || msg.String() == "b" {
		m.state = stateMainMenu
	}
	return m, nil
}

// Input handler when adjusting a specific property
func (m model) handleAdjustmentInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc", "enter":
		m.state = stateRGBMenu
		m.adjustTarget = adjustNone
		return m, nil
	}

	switch m.adjustTarget {
	case adjustBrightness:
		// Up/Down arrows adjust brightness
		switch msg.String() {
		case "up":
			if m.brightness < 9 { // 0-9 hardware limit from JSON
				m.brightness++
				cmd = m.setBrightnessCmd(m.brightness)
			}
		case "down":
			if m.brightness > 0 {
				m.brightness--
				cmd = m.setBrightnessCmd(m.brightness)
			}
		}

	case adjustEffect:
		switch msg.String() {
		case "up":
			if m.effect > 0 {
				m.effect--
				cmd = m.setEffectCmd(m.effect)
			}
		case "down":
			if m.effect < 18 {
				m.effect++
				cmd = m.setEffectCmd(m.effect)
			}
		}

	case adjustColor:
		// Left/Right arrows adjust color
		switch msg.String() {
		case "right":
			if m.colorHue < 255 {
				m.colorHue += 5
				cmd = m.setColorCmd(m.colorHue, m.colorSat)
			}
		case "left":
			if m.colorHue > 0 {
				m.colorHue -= 5
				cmd = m.setColorCmd(m.colorHue, m.colorSat)
			}
		}
	}

	return m, cmd
}
