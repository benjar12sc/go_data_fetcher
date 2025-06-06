param()
$ErrorActionPreference = 'Stop'
if (-not (Test-Path target)) { New-Item -ItemType Directory -Path target | Out-Null }
Push-Location excel_extractor
go build -o ../target/excel_data_extractor main.go
Pop-Location
