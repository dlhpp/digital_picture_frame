#!/bin/bash

# Define your binary name
BINARY_NAME="digital_picture_frame"

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
        echo "Built for ${os}/${arch}/v${goarm}"
    else
        GOOS=$os GOARCH=$arch go build -o "$output" .
        echo "Built for ${os}/${arch}"
    fi
done
