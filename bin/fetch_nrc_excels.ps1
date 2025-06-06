param()
$ErrorActionPreference = 'Stop'
$INPUT_DIR = "input_data"
$TMP_DIR = "/tmp/nrc_excels"
if (-not (Test-Path $INPUT_DIR)) { New-Item -ItemType Directory -Path $INPUT_DIR | Out-Null }
if (-not (Test-Path $TMP_DIR)) { New-Item -ItemType Directory -Path $TMP_DIR | Out-Null }
$urls = @(
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
foreach ($url in $urls) {
  $fname = Split-Path $url -Leaf
  $tmpfile = Join-Path $TMP_DIR $fname
  Write-Host "Fetching $fname ..."
  try {
    Invoke-WebRequest -Uri $url -OutFile $tmpfile -ErrorAction Stop
    Move-Item $tmpfile (Join-Path $INPUT_DIR $fname)
  } catch {
    Write-Host "Failed to fetch $url"
    if (Test-Path $tmpfile) { Remove-Item $tmpfile }
  }
}
