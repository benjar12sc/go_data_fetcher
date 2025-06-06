param()
$ErrorActionPreference = 'Stop'
if (-not (Test-Path target)) { New-Item -ItemType Directory -Path target | Out-Null }
Push-Location database_populator
go build -o ../target/database_populator main.go
Pop-Location
