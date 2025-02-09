#!/bin/bash
set -e

# Set environment variables for WebAssembly target.
export GOOS=js
export GOARCH=wasm

# Source file and output name (adjust if needed)
OUTPUT="main.wasm"

echo "Compiling for WebAssembly..."
go build -o "${OUTPUT}" .
echo "Compilation complete: ${OUTPUT} generated."

# Locate wasm_exec.js in your Go installation and copy it.
WASM_EXEC="$(go env GOROOT)/misc/wasm/wasm_exec.js"
if [ -f "${WASM_EXEC}" ]; then
    cp "${WASM_EXEC}" .
    echo "Copied wasm_exec.js from $(go env GOROOT)/misc/wasm."
else
    echo "Error: wasm_exec.js not found in $(go env GOROOT)/misc/wasm. Please check your Go installation."
fi
