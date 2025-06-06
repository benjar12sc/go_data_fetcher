#!/bin/zsh
set -e

INPUT_DIR="input_data"
OUTPUT_DIR="extracted_data"
BIN_PATH="target/excel_data_extractor"

mkdir -p "$OUTPUT_DIR"

for file in "$INPUT_DIR"/*; do
  if [[ -f "$file" && "$file" == *.xlsx ]]; then
    fname=$(basename "$file" .xlsx)
    out_subdir="$OUTPUT_DIR/$fname"
    mkdir -p "$out_subdir"
    "$BIN_PATH" "$file" "$out_subdir"
  fi
done
