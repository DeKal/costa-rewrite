package writer

import (
	"encoding/csv"
	"log"
	"os"
)

// WriteCsvFile to write CSV file for recording
func WriteCsvFile(csvFileName string, csvHeader [][]string, csvContent [][]string) {
	rows := append(csvHeader, csvContent...)

	csvfile, err := os.Create(csvFileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	csvfile.Close()
}
