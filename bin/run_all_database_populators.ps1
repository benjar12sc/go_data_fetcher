param()
$ErrorActionPreference = 'Stop'
$EXTRACTED_DIR = "extracted_data"
$BIN_PATH = "target/database_populator"
$CONFIG_PATH = "database_populator/mongo_config.yaml"
Get-ChildItem $EXTRACTED_DIR -Directory | ForEach-Object {
    Write-Host "Processing $($_.FullName) ..."
    & $BIN_PATH --folder_path $_.FullName --config $CONFIG_PATH
}
