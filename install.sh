#!/bin/bash

# Get latest release
LATEST_RELEASE=$(curl -s https://api.github.com/repos/aicraftlab-dev/aicraft-cli/releases/latest | grep "browser_download_url" | grep "$(uname -s | tr '[:upper:]' '[:lower:]')" | cut -d '"' -f 4)

# Download the binary
echo "Downloading AICraft CLI..."
curl -LO $LATEST_RELEASE

# Make it executable
BINARY_NAME=$(basename $LATEST_RELEASE)
chmod +x $BINARY_NAME

# Move to /usr/local/bin
echo "Installing to /usr/local/bin/aicraft..."
sudo mv $BINARY_NAME /usr/local/bin/aicraft

# Verify installation
if command -v aicraft &> /dev/null; then
    echo "AICraft CLI installed successfully!"
    echo "Version: $(aicraft --version)"
else
    echo "Installation failed. Please check permissions."
    exit 1
fi