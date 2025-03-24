# PiDisplay

Controller to manage Rasberry Pis with displays

Features:

- Separte device/display and controller system
- Very simple device/display server for low power consumption
- Controller to manage one or multiple displays based on web content

Support for [Inky Color E-Ink displays](https://shop.pimoroni.com/products/inky-impression-7-3).
More in the future as I buy more random stuff!

## How it works

1. The display server has an endpoint that receives and image and displays it
2. The controller takes screenshots of websites and sends them to the displays
