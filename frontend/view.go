package frontend

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View rendering logic
func (m model) View() tea.View {
	var s strings.Builder

	s.WriteString("=== Womier SK75 Control Panel ===\n\n")

	switch m.state {
	case stateMainMenu:
		s.WriteString("Main Menu:\n")
		options := []string{"RGB Settings", "Macro Settings"}
		for i, opt := range options {
			cursor := " "
			if m.mainMenuIdx == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
		}

	case stateRGBMenu:
		s.WriteString("Main Menu > RGB Settings:\n")
		options := []string{"Change Brightness", "Change Effect", "Change Effect Speed", "Change Color"}
		for i, opt := range options {
			cursor := " "
			if m.rgbMenuIdx == i {
				cursor = ">"
			}
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
		}
		s.WriteString("\n[esc] Back to Main Menu")

	case stateMacroMenu:
		s.WriteString("Main Menu > Macro Settings:\n")
		s.WriteString("  [Placeholder] Macro configuration module.\n")
		s.WriteString("\n[esc] Back to Main Menu")

	case stateAdjusting:
		s.WriteString("Adjustment Mode:\n\n")

		switch m.adjustTarget {

		case adjustBrightness:
			s.WriteString(fmt.Sprintf("Adjust Brightness (Use Up/Down Arrow):\n"))
			bar := strings.Repeat("█", m.brightness) + strings.Repeat("░", 9-m.brightness)
			s.WriteString(fmt.Sprintf("[%s] %d/9\n", bar, m.brightness)) // Checked against JSON bounds

		case adjustEffect:
			effectOptions := []string{
				"Off",
				"Wave",
				"Color Cloud",
				"Vortex",
				"Mix Color",
				"Breathe",
				"Light",
				"Slowly Off",
				"Stone",
				"Laser",
				"Starry",
				"Flowers Open",
				"Traverse",
				"Wave Bar",
				"Meteor",
				"Rain",
				"Scan",
				"Trigger Color",
				"Center Spread",
			}

			for i, opt := range effectOptions {
				cursor := " "
				if m.effect == i {
					cursor = ">"
				}
				s.WriteString(fmt.Sprintf("%s %s\n", cursor, opt))
			}

		case adjustColor:
			s.WriteString(fmt.Sprintf("Adjust Color Hue (Use Left/Right Arrow):\n"))
			s.WriteString(fmt.Sprintf("< Hue Value: %d >\n", m.colorHue))
		}

		s.WriteString("\n[enter/esc] Save and return to menu")
	}

	// Status bar area
	s.WriteString("\n" + strings.Repeat("-", 40) + "\n")
	if m.statusMessage != "" {
		s.WriteString(fmt.Sprintf("Status: %s\n", m.statusMessage))
	} else {
		s.WriteString("System Ready.\n")
	}

	if m.state != stateAdjusting {
		s.WriteString("[q] Quit Application\n")
	}

	return tea.NewView(s.String())
}
