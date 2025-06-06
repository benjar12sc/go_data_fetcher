# Database Populator

This tool loads all CSV files from a specified folder into a MongoDB database. The database name is derived from the folder name, and each CSV file is loaded into a collection named after the file (without extension, spaces/dashes replaced with underscores, and parentheses removed).

## Usage

```
go run main.go --folder_path <csv_folder> [--config <mongo_config.yaml>]
```

Or, after building:

```
./bin/database_populator --folder_path <csv_folder> [--config <mongo_config.yaml>]
```

- `--folder_path`: Path to the folder containing CSV files to load
- `--config`: Path to the MongoDB config YAML file (default: `mongo_config.yaml`)

## MongoDB Config Example (`mongo_config.yaml`)

```
uri: "mongodb://localhost:27017"
username: ""
password: ""
```

## Requirements

- Go 1.18 or newer
- MongoDB instance accessible from the config

## Notes

- The database name is the last folder name, sanitized.
- Each CSV file is loaded into a collection named after the file, sanitized.
- Only `.csv` files are processed.
