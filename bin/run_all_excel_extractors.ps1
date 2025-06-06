param()
$ErrorActionPreference = 'Stop'
$INPUT_DIR = "input_data"
$OUTPUT_DIR = "extracted_data"
$BIN_PATH = "target/excel_data_extractor"
if (-not (Test-Path $OUTPUT_DIR)) { New-Item -ItemType Directory -Path $OUTPUT_DIR | Out-Null }
Get-ChildItem $INPUT_DIR -Filter *.xlsx | ForEach-Object {
    $fname = $_.BaseName
    $out_subdir = Join-Path $OUTPUT_DIR $fname
    if (-not (Test-Path $out_subdir)) { New-Item -ItemType Directory -Path $out_subdir | Out-Null }
    & $BIN_PATH $_.FullName $out_subdir
}
