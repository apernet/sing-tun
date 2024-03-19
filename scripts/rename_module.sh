#!/bin/bash

OLD_MODULE_NAME="github.com/sagernet/sing-tun"
NEW_MODULE_NAME="github.com/apernet/sing-tun"

go mod edit -module "$NEW_MODULE_NAME"

find . -path '.git' -prune \
  -o -type f -name "*.go" \
  -exec sed -i "s|$OLD_MODULE_NAME|$NEW_MODULE_NAME|g" \
  '{}' '+'

