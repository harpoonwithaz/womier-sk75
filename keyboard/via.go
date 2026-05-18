package keyboard

import "errors"

// VIA PROTOCOL BYTE INSTRUCTIONS FOR WOMIER SK-75
const (
	CmdGetProtocolVersion = 0x01
	CmdSetKeyboardValue   = 0x07
	CmdGetKeyboardValue   = 0x08

	ChannelRGBMatrix = 0x03

	PropertyRGBBrightness  = 0x01
	PropertyRGBEffect      = 0x02
	PropertyRGBEffectSpeed = 0x03
	PropertyRGBColor       = 0x04
)

/*
	cmd (byte): The command to execute on the keyboard
	channel (byte): The component being targeted.
	property (byte): Specific property being changed.
	vals ([]byte): The value(s) the property will be changed to.

First byte contains the keyboard command (set, get, etc.).
Seconed byte includes which channel the packet will be sent to.
Third byte contains the property that will be set.
Fourth byte onward contains the value the property will be set to.
*/
func BuildPacket(cmd, channel, property byte, vals []byte) ([]byte, error) {
	// 32 bytes minus the header bytes
	if len(vals) > 29 {
		return nil, errors.New("must be less than 29 value parameters")
	}

	payload := make([]byte, 32)
	payload[0] = cmd
	payload[1] = channel
	payload[2] = property

	for i, val := range vals {
		payload[i+3] = val // offsets by the first bytes already included in the payload
	}

	return payload, nil
}

/*
	channel (byte): The component being targeted.
	property (byte): Specific property being changed.

First byte contains the set keyboard command.
Seconed byte includes which channel the packet will be sent to.
Third byte contains the property that will be set.
*/
// func BuiltGetPacket(channel, property byte) []byte {
// 	payload := make([]byte, 32)
// 	payload[0] = CmdGetKeyboardValue
// 	payload[1] = channel
// 	payload[2] = property

// 	return payload
// }
