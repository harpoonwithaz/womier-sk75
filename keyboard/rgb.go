// RGB effects for keyboard
package keyboard

import (
	"errors"
	"fmt"
)

type RGBProperty byte

const (
	PropBrightness RGBProperty = 0x01 // From JSON: id_qmk_rgb_matrix_brightness
	PropEffect     RGBProperty = 0x02 // From JSON: id_qmk_rgb_matrix_effect
	PropSpeed      RGBProperty = 0x03 // From JSON: id_qmk_rgb_matrix_effect_speed
	PropColor      RGBProperty = 0x04 // From JSON: id_qmk_rgb_matrix_color
)

type SetPreferences struct {
	property RGBProperty
	value    byte
	colorSat byte
}

// Set methods
// Gener
func (k *Keyboard) SetKeyboard(state SetPreferences) error {
	switch state.property {
	case PropBrightness:
		if state.value > 9 {
			return errors.New("brightness level must be >= 0 and <= 9")
		}
	case PropEffect:
		if state.value > 18 {
			return errors.New("effect preset must be between 0-18")
		}
	case PropSpeed:
		break
	case PropColor:
		if state.value > 44 {
			return errors.New("color must be between 0-44")
		}

		payload, _ := BuildSetPacket(ChannelRGBMatrix, byte(state.property), state.value, state.colorSat) // include a 255 because idk womier did it
		response, err := k.SendPacket(payload)
		// TODO: this part is just for testing to see the response bytes
		if err != nil {
			return err
		}

		fmt.Printf("Raw response: %v\n", response)
		return nil
	}

	payload, _ := BuildSetPacket(ChannelRGBMatrix, byte(state.property), state.value)
	response, err := k.SendPacket(payload)
	if err != nil {
		return err
	}

	fmt.Printf("Raw response: %v\n", response)
	return nil
}

// Get methods
func (k *Keyboard) GetEffect() ([]byte, error) {
	payload := BuiltGetPacket(ChannelRGBMatrix, PropertyRGBEffect)

	response, err := k.SendPacket(payload)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw response: %v\n", response)
	return response, nil
}

func (k *Keyboard) GetColor() ([]byte, error) {
	payload := BuiltGetPacket(ChannelRGBMatrix, PropertyRGBColor)

	response, err := k.SendPacket(payload)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw response: %v\n", response)
	return response, nil
}
