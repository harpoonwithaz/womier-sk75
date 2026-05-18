# womier-sk75

## udev rules

For the program to access and write to the keyboard, udev rules must be enabled

```bash
sudo nano /etc/udev/rules.d/99-womier.rules
UBSYSTEM=="usb", ATTR{idVendor}=="320F", ATTR{idProduct}=="5055", MODE="0666"
sudo udevadm control --reload-rules && sudo udevadm trigger
```
