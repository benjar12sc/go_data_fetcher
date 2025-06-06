param()
$ErrorActionPreference = 'Stop'
if (-not (Test-Path target)) { New-Item -ItemType Directory -Path target | Out-Null }
Push-Location api
go build -o ../target/nrc_api main.go
Pop-Location
