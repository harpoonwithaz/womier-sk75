# Womier SK75 Control Panel

A small Go-based terminal control panel for the **Womier SK75** mechanical keyboard. It connects to the keyboard over USB or 2.4 GHz wireless HID and lets you manage RGB lighting from an interactive TUI.

## Features

- Detects the Womier SK75 in wired and wireless modes
- Adjusts RGB brightness, effects, effect speed, and color
- Reads the keyboard's current RGB state
- Uses the keyboard's VIA-compatible HID protocol
- Lightweight terminal interface built with Bubble Tea

> This project is under active development. Macro configuration is currently a placeholder.

## Linux setup

The application needs permission to access the keyboard's HID device. Create a udev rule:

```bash
sudo nano /etc/udev/rules.d/99-womier.rules
```

Add the following line:

```text
SUBSYSTEM=="usb", ATTR{idVendor}=="320F", ATTR{idProduct}=="5055", MODE="0666"
```

Then reload the rules:

```bash
sudo udevadm control --reload-rules
sudo udevadm trigger
```

After that, build and run the application with Go:

```bash
go run .
```
