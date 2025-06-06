// ...existing code from main.go...
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input_excel_file> <output_directory>\n", os.Args[0])
		os.Exit(1)
	}
	inputPath := os.Args[1]
	outputDir := os.Args[2]

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	f, err := excelize.OpenFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		log.Fatalf("No sheets found in Excel file.")
	}

	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Printf("Skipping sheet '%s': %v", sheet, err)
			continue
		}
		if len(rows) == 0 {
			log.Printf("Skipping empty sheet '%s'", sheet)
			continue
		}
		csvFileName := fmt.Sprintf("%s.csv", sanitizeFileName(sheet))
		csvPath := filepath.Join(outputDir, csvFileName)
		csvFile, err := os.Create(csvPath)
		if err != nil {
			log.Printf("Failed to create CSV for sheet '%s': %v", sheet, err)
			continue
		}
		w := csv.NewWriter(csvFile)
		for _, row := range rows {
			if err := w.Write(row); err != nil {
				log.Printf("Failed to write row to CSV for sheet '%s': %v", sheet, err)
			}
		}
		w.Flush()
		if err := w.Error(); err != nil {
			log.Printf("Error flushing CSV for sheet '%s': %v", sheet, err)
		}
		csvFile.Close()
		fmt.Printf("Exported sheet '%s' to %s\n", sheet, csvPath)
	}
}

func sanitizeFileName(name string) string {
	// Replace spaces and slashes with underscores, remove problematic characters
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	return name
}
