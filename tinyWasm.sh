#!/usr/bin/env bash
#set -euo pipefail

# Define variables.
# Change these as needed for your project.
WASM_OUTPUT="output.wasm"
OPTIMIZED_WASM="output-opt.wasm"

echo "Compiling to WebAssembly using TinyGo..."
tinygo build -o "${WASM_OUTPUT}" -target wasm .

# Optimize the WASM output using wasm-opt, if available.
if command -v wasm-opt >/dev/null 2>&1; then
    echo "Optimizing ${WASM_OUTPUT} with wasm-opt..."
    # Using -Oz for aggressive size optimizations; adjust as needed.
    wasm-opt -Oz "${WASM_OUTPUT}" -o "${OPTIMIZED_WASM}"
    # Optionally, replace the original WASM file with the optimized version.
    mv "${OPTIMIZED_WASM}" "${WASM_OUTPUT}"
else
    echo "Warning: wasm-opt not found. Skipping optimization step."
fi

echo "Compressing ${WASM_OUTPUT} to ${WASM_OUTPUT}.gz..."
gzip -kf "${WASM_OUTPUT}"  # -k: keep the original file, -f: force compression if necessary

echo "Build complete! Generated files:"
echo "  - ${WASM_OUTPUT}"
echo "  - ${WASM_OUTPUT}.gz"
