#!/bin/zsh
set -e

EXTRACTED_DIR="extracted_data"
BIN_PATH="target/database_populator"
CONFIG_PATH="database_populator/mongo_config.yaml"

for folder in "$EXTRACTED_DIR"/*; do
  if [ -d "$folder" ]; then
    echo "Processing $folder ..."
    "$BIN_PATH" --folder_path "$folder" --config "$CONFIG_PATH"
  fi
done
