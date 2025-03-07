#!/bin/bash

set -e

# Test variables
TEST_DIR=$(mktemp -d)
export HOME=$TEST_DIR

# Mock curl
curl() {
    echo "https://github.com/aicraftlab-dev/aicraft-cli/releases/download/v0.1.0/aicraft-cli-ubuntu-latest-alpha"
}

# Mock sudo
sudo() {
    "$@"
}

# Source the install script
source ../install.sh

# Verify installation
if ! command -v aicraft &> /dev/null; then
    echo "Test failed: aicraft not found in PATH"
    exit 1
fi

echo "All tests passed!"
rm -rf $TEST_DIR