package main

import "womier-sk75/keyboard"

func main() {
	kb := &keyboard.Keyboard{}

	kb.Connect()

	// var setPreferences keyboard.SetPreferences
	// setPreferences.property = PropEffect
	// setPreferences.colorSat = 0
	// kb.SetKeyboard()

	// kb.GetEffect()
	// kb.GetColor()

	defer kb.Disconnect()
}
