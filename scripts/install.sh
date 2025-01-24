#!/bin/bash

# If no version is specified as a command line argument, fetch the latest version.
if [ -z "$1" ]; then
    VERSION=$(curl -s https://api.github.com/repos/ry0y4n/azurenv/releases/latest | grep -o '"tag_name": *"[^"]*"' | sed 's/"tag_name": *"//' | sed 's/"//')
    if [ -z "$VERSION" ]; then
        echo "Failed to fetch the latest version"
        exit 1
    fi
else
    VERSION=$1
fi

OS=$(uname -s)
ARCH=$(uname -m)
URL="https://github.com/ry0y4n/azurenv/releases/download/${VERSION}/azurenv_${OS}_${ARCH}.tar.gz"

echo "Start to install."
echo "VERSION=$VERSION, OS=$OS, ARCH=$ARCH"
echo "URL=$URL"

TMP_DIR=$(mktemp -d)
curl -L $URL -o $TMP_DIR/azurenv.tar.gz
tar -xzvf $TMP_DIR/azurenv.tar.gz -C $TMP_DIR
sudo mv $TMP_DIR/azurenv /usr/local/bin/azurenv
sudo chmod +x /usr/local/bin/azurenv

rm -rf $TMP_DIR

if [ -f "/usr/local/bin/azurenv" ]; then
  echo "[SUCCESS] azurenv $VERSION has been installed to /usr/local/bin"
else
  echo "[FAIL] azurenv $VERSION is not installed."
fi