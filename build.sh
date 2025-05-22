#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

APP_NAME="$(basename "$SCRIPT_DIR")"

PLATFORMS=("linux/amd64" "windows/amd64" "darwin/amd64" "darwin/arm64")

for PLATFORM in "${PLATFORMS[@]}"
do
    IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
    OUTPUT_NAME="${APP_NAME}-${GOOS}-${GOARCH}"
    if [ "$GOOS" == "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "build/$OUTPUT_NAME"
done

echo "Build complete. Files are in ./build/"
