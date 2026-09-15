#!/usr/bin/env bash
set -euo pipefail
# Use the standard context type without changing generated method signatures.
for file in pkg/*.go; do
    sed 's|"golang.org/x/net/context"|"context"|g' "${file}" > "${file}.tmp"
    mv "${file}.tmp" "${file}"
done
gofmt -w pkg/*.go

# Replace the historical documentation URL supplied by the schema.
sed 's|Wodby Developer Documentation https://wodby.com/docs/1.0/docs/dev|Wodby 1.0 API reference: https://wodby.com/docs/1.0/api/|g' pkg/README.md > pkg/README.md.tmp
mv pkg/README.md.tmp pkg/README.md
