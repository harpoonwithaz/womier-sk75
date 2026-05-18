package utils

import (
	"fmt"
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

func GetDevices() {
	hids := hid.Enumerate(0, 0)
	printHidsInfo(hids)
}
