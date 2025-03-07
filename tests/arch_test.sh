#!/bin/bash

set -e

# Test variables
TEST_DIR=$(mktemp -d)
cd $TEST_DIR

# Test Intel macOS binary
echo "Testing Intel macOS binary..."
curl -LO https://github.com/aicraftlab-dev/aicraft-cli/releases/latest/download/aicraft-cli-macos-intel-alpha
chmod +x aicraft-cli-macos-intel-alpha
./aicraft-cli-macos-intel-alpha --version

# Test ARM macOS binary
echo "Testing ARM macOS binary..."
curl -LO https://github.com/aicraftlab-dev/aicraft-cli/releases/latest/download/aicraft-cli-macos-arm64-alpha
chmod +x aicraft-cli-macos-arm64-alpha
./aicraft-cli-macos-arm64-alpha --version

# Clean up
echo "All architecture tests passed!"
rm -rf $TEST_DIR