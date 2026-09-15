#!/usr/bin/env bash
set -euo pipefail
# Use the standard context type without changing generated method signatures.
for file in pkg/*.go; do
    sed 's|"golang.org/x/net/context"|"context"|g' "${file}" > "${file}.tmp"
    mv "${file}.tmp" "${file}"
done
gofmt -w pkg/*.go
