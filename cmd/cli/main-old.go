package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/karalabe/hid"
)

func printHidInfo(hid hid.DeviceInfo) {
	fmt.Printf(" OS Path: %s\n", hid.Path)
	fmt.Printf(" Vendor ID: %#04x\n", hid.VendorID)
	fmt.Printf(" Product ID: %#04x\n", hid.ProductID)
	fmt.Printf(" Release: %d\n", hid.Release)
	fmt.Printf(" Serial: %s\n", hid.Serial)
	fmt.Printf(" Manufacturer: %s\n", hid.Manufacturer)
	fmt.Printf(" Product: %s\n", hid.Product)
	fmt.Printf(" Usage Page: %#04x\n", hid.UsagePage)
	fmt.Printf(" Usage: %d\n", hid.Usage)
	fmt.Printf(" Interface: %d\n", hid.Interface)
}

func printHidsInfo(hids []hid.DeviceInfo) {
	fmt.Printf("hid.Supported() %v\n", hid.Supported())

	for i, hid := range hids {
		fmt.Println(strings.Repeat("-", 64))
		fmt.Printf("HID #%d\n", i)
		printHidInfo(hid)
	}

	fmt.Println(strings.Repeat("=", 64))
}

func checkProtocol(device *hid.Device) {
	packet := make([]byte, 33)
	packet[0] = 0x00 // Report ID
	packet[1] = 0x01 // Command: id_get_protocol_version

	_, err := device.Write(packet)
	if err != nil {
		log.Fatalf("Write error: %v", err)
	}

	// Read the response (32 bytes)
	res := make([]byte, 32)
	_, err = device.Read(res)
	if err != nil {
		log.Fatalf("Read error: %v", err)
	}

	// VIA usually returns [0x01, v_high, v_low]
	fmt.Printf("Protocol Version Response: % x\n", res[:3])
	// fmt.Printf("Raw response: %v\n", res)
}

func printResponse(bytes []byte) {
	fmt.Println("Byte   | Value")
	for i, b := range bytes {
		fmt.Printf("%v     | %v\n", i, b)
	}
}

func getBrightness(device *hid.Device) {
	packet := make([]byte, 33)

	packet[0] = 0x00
	packet[1] = 0x08 // id_custom_get_value (V2 style)
	packet[2] = 0x03 // Channel: RGB Matrix
	packet[3] = 0x01 // Property: Brightness

	device.Write(packet)

	res := make([]byte, 32)
	device.Read(res)

	// VIA returns [Command, Channel, Property, Value]
	fmt.Println("Current Brightness Data:")

	printResponse(res)
}

func setBrightness(device *hid.Device) {
	packet := make([]byte, 33)

	packet[0] = 0x00
	packet[1] = 0x07 // id_custom_get_value (V2 style)
	packet[2] = 0x03 // Channel: RGB Matrix
	packet[3] = 0x02 // Property: Brightness
	packet[4] = 0x04 // set it to off

	device.Write(packet)

	res := make([]byte, 32)
	device.Read(res)

	// VIA returns [Command, Channel, Property, Value]
	fmt.Println("Current Brightness Data:")
	fmt.Printf("Raw bytes: % x\n", res)

}

type Payload struct {
	command  byte
	channel  byte
	property byte
	value    byte
}

func sendPacket(device *hid.Device, payload Payload) ([]byte, error) {
	// Create a 33 byte slice containing all zeros
	packet := make([]byte, 33)

	packet[0] = 0x00 // First byte is
	packet[1] = payload.command
	packet[2] = payload.channel
	packet[3] = payload.property
	packet[4] = payload.value

	_, err := device.Write(packet)
	if err != nil {
		return nil, err // return an empty packet
	}

	res := make([]byte, 32)
	_, err = device.Read(res)

	if err != nil {
		return res, err // return an empty packet
	}

	return res, nil
}

func main() {
	const vendorID uint16 = 0x320F
	const productID uint16 = 0x5055
	hids := hid.Enumerate(vendorID, productID)

	var targetDevice *hid.DeviceInfo

	// printHidsInfo(hids)

	for _, info := range hids {
		// The Raw HID interface for VIA/QMK is typically 0xFF60.
		// If UsagePage shows 0, try Interface 3 (do not use 0, cuz your keyboard will stop working :) ).
		if info.Interface == 1 {
			targetDevice = &info
			fmt.Println("Found lighting interface")
			printHidInfo(info)
			break
		}
	}

	if targetDevice == nil {
		log.Fatal("Could not find the Raw HID lighting interface.")
	}

	// device, err := targetDevice.Open()
	// Inside your main after targetDevice.Open()
	device, err := targetDevice.Open()
	// fmt.Printf("Connected Usage Page: %#04x, Usage: %#04x\n", targetDevice.UsagePage, targetDevice.Usage)
	if err != nil {
		log.Fatalf("Failed to open lighting interface: %v", err)
	}

	// fmt.Println("Successfully connected to lighting controller.")

	// Perform your RGB experiments here
	setBrightness(device)
	// checkProtocol(device)
	// getBrightness(device)

	// printHidInfo(*targetDevice)

	defer device.Close()
	fmt.Println("Device successfully closed")
}
