#!/bin/bash

BINARY_NAME="slideshow"
PLATFORMS=("linux/amd64"  "linux/arm64"  "linux/arm/7"  "linux/arm/6"  "windows/amd64")

GOOS=linux   GOARCH=arm  GOARM=6  go build -o "${BINARY_NAME}_linux_armv6.exe" .
GOOS=linux   GOARCH=arm  GOARM=7  go build -o "${BINARY_NAME}_linux_armv7.exe" .

GOOS=linux   GOARCH=arm64         go build -o "${BINARY_NAME}_linux_arm64.exe" .
GOOS=linux   GOARCH=amd64         go build -o "${BINARY_NAME}_linux_amd64.exe" .

GOOS=windows GOARCH=amd64         go build -o "${BINARY_NAME}_windows_amd64.exe" .



# Output files:
# -----------------
# slideshow_linux_armv6.exe
# slideshow_linux_armv7.exe
# slideshow_linux_arm64.exe
# slideshow_linux_amd64.exe
# slideshow_windows_amd64.exe

