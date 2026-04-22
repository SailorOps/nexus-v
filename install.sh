#!/bin/bash
# NEXUS-V Linux/macOS Install Script

echo "Installing NEXUS-V..."

# Determine OS and Arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH="amd64"
if [[ "$(uname -m)" == "arm64" || "$(uname -m)" == "aarch64" ]]; then
    ARCH="arm64"
fi

# Download binary
URL="https://github.com/geriatric-sailor/nexus-v/releases/latest/download/nexus-v-${OS}-${ARCH}"
curl -L "$URL" -o nexus-v

# Make executable and move to bin
chmod +x nexus-v
sudo mv nexus-v /usr/local/bin/

echo "NEXUS-V installed successfully! Run 'nexus-v version' to verify."
