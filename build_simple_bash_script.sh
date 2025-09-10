#!/bin/bash

BINARY_NAME="exe.digital_picture_frame"
PLATFORMS=("linux/amd64"  "linux/arm64"  "linux/arm/7"  "linux/arm/6"  "windows/amd64")

GOOS=windows GOARCH=amd64         go build -o "${BINARY_NAME}_windows_amd64" .
GOOS=linux   GOARCH=amd64         go build -o "${BINARY_NAME}_linux_amd64" .
GOOS=linux   GOARCH=arm64         go build -o "${BINARY_NAME}_linux_arm64" .

GOOS=linux   GOARCH=arm  GOARM=6  go build -o "${BINARY_NAME}_linux_armv6" .
GOOS=linux   GOARCH=arm  GOARM=7  go build -o "${BINARY_NAME}_linux_armv7" .


# Output files:
# -----------------
# exe.digital_picture_frame_linux_amd64
# exe.digital_picture_frame_linux_arm64
# exe.digital_picture_frame_linux_armv6
# exe.digital_picture_frame_linux_armv7
# exe.digital_picture_frame_windows_amd64.exe

