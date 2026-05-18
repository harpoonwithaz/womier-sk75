// This file defines the Keyboard struct. It is the only place in your project that directly touches the hid library.

package keyboard

import (
	"errors"
	"fmt"
	"sync"

	"github.com/karalabe/hid"
)

type Keyboard struct {
	device     *hid.Device
	interfaces []hid.DeviceInfo
	mu         sync.Mutex
}

// Finds interface of known keyboards
func (k *Keyboard) DetectKeyboard() {
	// hard coded for now for my keyboards wireless and usb mode
	// in the future, we can read via json config files and have
	// a list of known vendor and product ids, and expand the functionality for
	// mulitple known keyboards detected
	const targetVID uint16 = 0x320F
	const targetPIDwired uint16 = 0x5055
	const targetPIDwireless uint16 = 0x5088

	var interfaces []hid.DeviceInfo

	hids := hid.Enumerate(0, 0)

	for _, h := range hids {
		// yes ik this is lazy, i will make it better later
		if h.VendorID == targetVID && (h.ProductID == targetPIDwired || h.ProductID == targetPIDwireless) {
			fmt.Printf("Found a known device with vid: %x, pid: %x and interface #%v\n", h.VendorID, h.ProductID, h.Interface)
			interfaces = append(interfaces, h) // add the interface which matches our target
		}
	}

	k.mu.Lock()
	k.interfaces = interfaces
	k.mu.Unlock()
}

func (k *Keyboard) Connect() error {
	// copy slice under lock, then operate on the copy (avoid pointers into k.interfaces)
	k.mu.Lock()
	if len(k.interfaces) == 0 {
		k.mu.Unlock()
		return errors.New("cannot connect, keyboard has not been detected")
	}
	interfacesCopy := make([]hid.DeviceInfo, len(k.interfaces))
	copy(interfacesCopy, k.interfaces)
	k.mu.Unlock()

	// Find interface 1 on the copy
	var target hid.DeviceInfo
	found := false
	for _, info := range interfacesCopy {
		if info.Interface == 1 {
			target = info // copy the struct
			found = true
			fmt.Println("Found lighting interface")
			break
		}
	}

	if !found {
		return errors.New("could not find the Raw HID lighting interface")
	}

	// Open without holding the mutex (avoid blocking other ops)
	d, err := target.Open()
	if err != nil {
		return err
	}

	k.mu.Lock()
	k.device = d
	k.mu.Unlock()

	fmt.Println("Successfully connected to keyboard")
	return nil
}

func (k *Keyboard) Disconnect() error {
	k.mu.Lock()
	d := k.device
	k.device = nil
	k.mu.Unlock()

	if d != nil {
		err := d.Close()
		if err != nil {
			return err
		}
	}

	fmt.Println("Successfully disconnected keyboard")
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

	k.mu.Lock()
	defer k.mu.Unlock()

	if k.device == nil {
		return nil, errors.New("device not connected")
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

func (k *Keyboard) GetProtocol() ([]byte, error) {
	// get protocol version only requires the cmd,
	// so the channel, property and values are 0
	payload, err := BuildPacket(CmdGetProtocolVersion, 0x00, 0x00, []byte{0x00})
	if err != nil {
		return nil, err
	}

	response, err := k.SendPacket(payload)
	if err != nil {
		return nil, err
	}

	return response, err
}
