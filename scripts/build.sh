#!/bin/bash
# Build script for Nextcloud Compose Generator app
set -e

# Change to project root
dirname=$(dirname "$0")
cd "$dirname/.."

# Build Go binaries (example: build cmd/main.go to bin/nextcloud-compose-gen)
mkdir -p bin
go build -o bin/nextcloud-gen ./cmd/...

echo "Build complete. Binary is in bin/nextcloud-gen."
