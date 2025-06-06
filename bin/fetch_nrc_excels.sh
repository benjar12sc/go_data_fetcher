#!/bin/zsh
set -e

INPUT_DIR="input_data"
TMP_DIR="/tmp/nrc_excels"
mkdir -p "$INPUT_DIR"
mkdir -p "$TMP_DIR"

urls=(
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/advanced-reactor-integrated-schedule.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-canceled.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-operating.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-operating-under-construction.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-decommissioning.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/commercial-nuclear-power-plant-licensing-history.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/waste-spent-fuel-storage-dry-cask.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/waste-spent-fuel-storage-designs.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/import-export-licenses-imports.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/import-export-licenses-exports.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/emergency-native-american-reservations.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-applications.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/decommissioning-complex-materials-sites.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-research-test-decommissioning.xlsx"
  "https://www.nrc.gov/reading-rm/doc-collections/datasets/reactors-research-test.xlsx"
)

for url in "${urls[@]}"; do
  fname=$(basename "$url")
  tmpfile="$TMP_DIR/$fname"
  echo "Fetching $fname ..."
  if curl -fLo "$tmpfile" "$url"; then
    mv "$tmpfile" "$INPUT_DIR/$fname"
  else
    echo "Failed to fetch $url"
    rm -f "$tmpfile"
  fi
done
