#!/bin/bash

# Set environment variables for WebAssembly target.
export GOOS=js
export GOARCH=wasm

rm *.wasm *.gz

# Source file and output name (adjust if needed)
OUTPUT="main.wasm"

echo "Compiling for WebAssembly..."
# Strip debug information to reduce size
go build -ldflags="-s -w" -o "${OUTPUT}" .
echo "Compilation complete: ${OUTPUT} generated."

# Optimize the WASM binary using wasm-opt if available.
if command -v wasm-opt >/dev/null 2>&1; then
    echo "Optimizing ${OUTPUT} with wasm-opt..."
    # Attempt to enable bulk memory operations.
    wasm-opt --enable-bulk-memory -O3 "${OUTPUT}" -o "${OUTPUT}"
    echo "Optimization complete: ${OUTPUT} optimized."
else
    echo "wasm-opt not found, skipping WASM optimization."
fi

# Locate wasm_exec.js in your Go installation and copy it.
WASM_EXEC="$(go env GOROOT)/misc/wasm/wasm_exec.js"
if [ -f "${WASM_EXEC}" ]; then
    cp "${WASM_EXEC}" .
    echo "Copied wasm_exec.js from $(go env GOROOT)/misc/wasm."
else
    echo "Error: wasm_exec.js not found in $(go env GOROOT)/misc/wasm. Please check your Go installation."
fi

# Check if gzip is available and compress the output files.
if command -v gzip -9 >/dev/null 2>&1; then
    echo "Compressing ${OUTPUT} with gzip..."
    gzip -k -f "${OUTPUT}"
    echo "Compressed file created: ${OUTPUT}.gz"
    
    if [ -f "wasm_exec.js" ]; then
        echo "Compressing wasm_exec.js with gzip..."
        gzip -k -f "wasm_exec.js"
        echo "Compressed file created: wasm_exec.js.gz"
    fi
else
    echo "gzip command not found, skipping compression."
fi

rm *.wasm