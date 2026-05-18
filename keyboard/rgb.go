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

// Set methods
func (k *Keyboard) SetKeyboard(property RGBProperty, values []byte) error {
	switch property {
	case PropBrightness:
		if values[0] > 9 {
			return errors.New("brightness level must be >= 0 and <= 9")
		}
	case PropEffect:
		if values[0] > 18 {
			return errors.New("effect preset must be between 0-18")
		}
	case PropSpeed:
		break
	case PropColor:
		if values[0] > 44 {
			return errors.New("color must be between 0-44")
		}
	}

	payload, err := BuildPacket(CmdSetKeyboardValue, ChannelRGBMatrix, byte(property), values)
	if err != nil {
		return err
	}

	response, err := k.SendPacket(payload)
	if err != nil {
		return err
	}

	fmt.Printf("Raw response: %v\n", response)
	return nil
}

// Get methods
func (k *Keyboard) GetKeyboard(property RGBProperty) ([]byte, error) {
	payload, err := BuildPacket(CmdGetKeyboardValue, ChannelRGBMatrix, byte(property), []byte{0x00}) // get packet doesnt need values
	if err != nil {
		return nil, err
	}

	response, err := k.SendPacket(payload)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw response: %v\n", response) // for testing purposes
	return response, nil
}
