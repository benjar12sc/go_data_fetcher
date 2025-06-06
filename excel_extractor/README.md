# Excel Data Extractor

This tool converts each sheet in a given Excel spreadsheet into a separate CSV file. The output CSV files are saved in a specified directory. The program will create the output directory if it does not exist.

## Usage

```
go run main.go <input_excel_file> <output_directory>
```

Or, after building:

```
./bin/excel_data_extractor <input_excel_file> <output_directory>
```

- `<input_excel_file>`: Path to the Excel file (e.g., `input_data/advanced-reactor-integrated-schedule.xlsx`)
- `<output_directory>`: Path to the directory where CSV files will be saved

## Build

To build the program and place the binary in the `bin` directory:

```
./bin/build_excel_extractor.sh
```

## Requirements

- Go 1.18 or newer

## Notes

- Each sheet in the Excel file will be exported as a separate CSV file named after the sheet.
- The program will skip sheets that cannot be converted to CSV and print an error message.
