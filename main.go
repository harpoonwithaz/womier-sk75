package main

import (
	"fmt"
	"os"
	"womier-sk75/keyboard"
)

func main() {
	kb := &keyboard.Keyboard{}

	kb.DetectKeyboard()
	err := kb.Connect()
	if err != nil {
		fmt.Printf("There was an error: %v\n", err)
		os.Exit(0)
	}

	kb.SetKeyboard(keyboard.PropertyRGBEffect, []byte{2})
	defer kb.Disconnect()

	// utils.GetDevices()

}
