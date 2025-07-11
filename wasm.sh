#!/bin/bash

# Set environment variables for WebAssembly target.
export GOOS=js
export GOARCH=wasm

rm *.wasm *.gz *.opt

# Source file and output name (adjust if needed)
OUTPUT="main.wasm"
OUTPUTOPT="main.wasm.opt"

echo "Compiling for WebAssembly..."
# Strip debug information to reduce size
GOOS=js GOARCH=wasm go build -o ${OUTPUT} -ldflags="-s -w" -gcflags="-trimpath=${PWD}" -trimpath . 
echo "Compilation complete: ${OUTPUT} generated."

# Optimize the WASM binary using wasm-opt if available.
if command -v wasm-opt >/dev/null 2>&1; then
    echo "Optimizing ${OUTPUT} with wasm-opt..."
    # Attempt to enable bulk memory operations.
    wasm-opt --enable-bulk-memory -O4 "${OUTPUT}" -o "${OUTPUTOPT}"
    echo "Optimization complete: ${OUTPUT} optimized."
    rm ${OUTPUT}
    mv ${OUTPUTOPT} ${OUTPUT}
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
if command -v gzip -1 >/dev/null 2>&1; then
    echo "Compressing ${OUTPUT} with gzip..."
    gzip -k -f "${OUTPUT}"
    echo "Compressed file created: ${OUTPUT}.gz"
else
    echo "gzip command not found, skipping compression."
fi

rm *.wasm *.opt