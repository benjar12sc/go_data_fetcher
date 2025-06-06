# NRC Data Extraction and Database Loader

**NOTE: This is a demo project. The code is not production ready and omits best practices for clarity and brevity.**

This project provides tools to automate the download, extraction, and database population of U.S. Nuclear Regulatory Commission (NRC) public datasets. It is designed to fetch Excel files from the NRC, convert each sheet to CSV, and load the resulting data into a MongoDB database for further analysis or integration.

## Features

- **Automated Download**: Fetches the latest NRC Excel datasets from the NRC's public data portal.
- **Excel to CSV Extraction**: Converts each sheet in every Excel file to a separate CSV file.
- **Batch Database Population**: Loads all CSVs into MongoDB, using folder and file names to determine database and collection names.
- **Shell Scripts**: Includes scripts for each step, making the workflow easy to automate or schedule.

## Data Source

All data is sourced from the U.S. Nuclear Regulatory Commission's public datasets:

https://www.nrc.gov/reading-rm/doc-collections/datasets/index.html

> The U.S. Nuclear Regulatory Commission is in the process of rescinding or revising guidance and policies posted on this webpage in accordance with Executive Order 14151 Ending Radical and Wasteful Government DEI Programs and Preferencing, and Executive Order 14168 Defending Women From Gender Ideology Extremism and Restoring Biological Truth to the Federal Government. In the interim, any previously issued diversity, equity, inclusion, or gender-related guidance on this webpage should be considered rescinded that is inconsistent with these Executive Orders.

## Project Structure

- `bin/` — Shell scripts and compiled binaries for automation
- `input_data/` — Downloaded Excel files
- `extracted_data/` — CSVs extracted from Excel sheets
- `excel_extractor/` — Go code for Excel to CSV extraction
- `database_populator/` — Go code for loading CSVs into MongoDB
- `api/` — Go REST API for querying loaded data (demo only)

## Quickstart

1. **Fetch NRC Excel files:**
   ```
   ./bin/fetch_nrc_excels.sh
   ```
2. **Extract all Excel sheets to CSV:**
   ```
   ./bin/run_all_excel_extractors.sh
   ```
3. **Load all CSVs into MongoDB:**
   ```
   ./bin/run_all_database_populators.sh
   ```
4. **Run the API (demo):**
   ```
   ./bin/build_nrc_api.sh && ./bin/run_nrc_api.sh
   ```

See the subproject `README.md` files for more details on configuration and usage.
