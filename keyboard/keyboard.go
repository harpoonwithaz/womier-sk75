// This file defines the Keyboard struct. It is the only place in your project that directly touches the hid library.

package keyboard

import (
	"errors"
	"fmt"

	"github.com/karalabe/hid"
)

type Keyboard struct {
	device *hid.Device
}

const vendorID uint16 = 0x320F
const productID uint16 = 0x5055

func (k *Keyboard) Connect() error {
	// Get a list of hids under the vendor and product ID
	hids := hid.Enumerate(vendorID, productID)
	var targetDevice *hid.DeviceInfo

	// Find interface 1 and connect to it
	// ** DO NOT USE INTERFACE 0, IT WILL CRASH THE KEYBOARD **
	for _, info := range hids {
		if info.Interface == 1 {
			targetDevice = &info
			fmt.Println("Found lighting interface")
			break
		}
	}

	if targetDevice == nil {
		return errors.New("could not find the Raw HID lighting interface")
	}

	var err error
	k.device, err = targetDevice.Open()
	if err != nil {
		return err
	}

	return nil
}

func (k *Keyboard) Disconnect() error {
	err := k.device.Close()
	if err != nil {
		return err
	}

	return nil
}

// SendPacket is the centralized "Write/Read" function.
// It adds the 0x00 Report ID and returns the keyboard's 32-byte response.

// Reads the response
// Returns slice of the raw response
func (k *Keyboard) SendPacket(payload []byte) ([]byte, error) {
	if len(payload) != 32 {
		return nil, errors.New("payload must be 32 bytes")
	}

	// First byte is the report id
	packet := []byte{0x00}
	packet = append(packet, payload...)

	_, err := k.device.Write(packet)
	if err != nil {
		return nil, err
	}

	res := make([]byte, 32)
	_, err = k.device.Read(res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
