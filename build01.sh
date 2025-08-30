#!/bin/bash

# Define your binary name
BINARY_NAME="exe.digital_picture_frame"

# List of platforms (OS/ARCH pairs)
PLATFORMS=("linux/amd64"  "linux/arm64"  "linux/arm/7"  "linux/arm/6"  "windows/amd64")

# Compile for each platform
for platform in "${PLATFORMS[@]}"; do
    IFS='/' read -r os arch goarm <<< "$platform"
    output="${BINARY_NAME}_${os}_${arch}"
    if [ "$os" = "windows" ]; then
        output="${output}.exe" # Windows binaries typically have .exe extension
    fi
    if [ "$arch" = "arm" ]; then
        output="${output}v${goarm}"
        GOOS=$os GOARCH=$arch GOARM=$goarm go build -o "$output" .
        echo "#1 Built for ${os}/${arch}/v${goarm}"
    else
        GOOS=$os GOARCH=$arch go build -o "$output" .
        echo "#2 Built for ${os}/${arch}"
    fi
done

# Output example:
# -----------------------------
#2 Built for linux/amd64
#2 Built for linux/arm64
#1 Built for linux/arm/v7
#1 Built for linux/arm/v6
#2 Built for windows/amd64

# Output files:
# -----------------
# exe.digital_picture_frame_linux_amd64
# exe.digital_picture_frame_linux_arm64
# exe.digital_picture_frame_linux_armv6
# exe.digital_picture_frame_linux_armv7
# exe.digital_picture_frame_windows_amd64.exe
