#!/bin/bash
for file in *; do
  if [[ -f "$file" ]]; then
    if grep -q "^//go:build.* with_gvisor" "$file"; then
      rm "$file"
    fi
  fi
done

sed -i 's| -tags with_gvisor||' Makefile

sed -i '/\/\/go:build.* !with_gvisor/d' stack_gvisor_stub.go
sed -i 's|gVisor is not [^`]*|gVisor is not supported in this fork.|' stack_gvisor_stub.go

go mod tidy
