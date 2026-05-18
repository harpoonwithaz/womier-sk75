// RGB effects for keyboard
package keyboard

import (
	"errors"
)

type RGBProperty byte

const (
	PropBrightness RGBProperty = 0x01 // From JSON: id_qmk_rgb_matrix_brightness
	PropEffect     RGBProperty = 0x02 // From JSON: id_qmk_rgb_matrix_effect
	PropSpeed      RGBProperty = 0x03 // From JSON: id_qmk_rgb_matrix_effect_speed
	PropColor      RGBProperty = 0x04 // From JSON: id_qmk_rgb_matrix_color
)

type RGBState struct {
	Brightness  int
	Effect      int
	EffectSpeed int
	ColorHue    int
	ColorSat    int
}

// Set methods
func (k *Keyboard) SetRGB(property RGBProperty, values []byte) error {
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
		if values[0] > 4 {
			return errors.New("speed must be between 0-4")
		}
	case PropColor:
		// if values[0] > 44 {
		// 	return errors.New("color must be between 0-44")
		// }
	}

	payload, err := BuildPacket(CmdSetKeyboardValue, ChannelRGBMatrix, byte(property), values)
	if err != nil {
		return err
	}

	_, err = k.SendPacket(payload)
	if err != nil {
		return err
	}

	// fmt.Printf("Raw res: %v\n", res)

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

	return response, nil
}

func (k *Keyboard) ReadState() (*RGBState, error) {
	var err error

	getProp := func(prop RGBProperty) []byte {
		if err != nil {
			return nil
		}
		var val []byte
		val, err = k.GetKeyboard(prop)
		return val
	}

	b := getProp(PropBrightness)
	e := getProp(PropEffect)
	es := getProp(PropSpeed)
	c := getProp(PropColor)

	if err != nil {
		return nil, err
	}

	return &RGBState{
		// fourth byte onwards holds values
		Brightness:  int(b[3]),
		Effect:      int(e[3]),
		EffectSpeed: int(es[3]),
		ColorHue:    int(c[3]), // fourth byte is hue
		ColorSat:    int(c[4]), // fifth byte is sat
	}, nil
}
