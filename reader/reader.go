package reader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Utils "github.com/DeKal/costa-rewrite/utils"
)

const (
	rating           = "Autocorrect Rating"
	labelResult      = "Event Label"
	hk               = "HK"
	id               = "ID"
	my               = "MY"
	ph               = "PH"
	sh               = "SG"
	tw               = "TW"
	label            = "Label"
	searchTerm       = "Search Term"
	correctTerm      = "Correct Term"
	orgSearchTerm    = "Orginal Search Term"
	orgCorrectedTerm = "Original Corrected Term"
)

// ReadSearchTermsFromExcel from file excel
func ReadSearchTermsFromExcel(inputCsvFileName string, country string) []DataFormat.AutoCorrectRow {
	file, _ := os.Open(inputCsvFileName)
	defer file.Close()
	csvr := csv.NewReader(file)
	csvHeader := ReadHeader(csvr)
	if _, ok := csvHeader[country]; !ok {
		log.Fatal("Country code is not existed")
		os.Exit(1)
	}
	return ReadSearchTerms(csvr, csvHeader, country)
}

// ReadHeader readheader from csv file and return a field map
func ReadHeader(csvr *csv.Reader) map[string]int {
	fieldMap := map[string]int{}
	header, err := csvr.Read()
	if err == io.EOF {
		return fieldMap
	}
	for index, field := range header {
		fieldMap[field] = index
	}

	return fieldMap
}

// ReadSearchTerms from file
func ReadSearchTerms(csvr *csv.Reader, fieldMap map[string]int, country string) []DataFormat.AutoCorrectRow {
	autoCorrects := []DataFormat.AutoCorrectRow{}
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			return autoCorrects
		}
		if isTargetedCountry(row, fieldMap[country]) {
			searchTerm, correctedTerm := SplitResult(row[fieldMap[labelResult]])
			autoCorrect := DataFormat.AutoCorrectRow{
				Rating:            row[fieldMap[rating]],
				OriginSearchTerm:  searchTerm,
				OriginCorrectTerm: correctedTerm,
			}
			autoCorrects = append(autoCorrects, autoCorrect)
		}

	}
}

func isTargetedCountry(row []string, countryPos int) bool {
	return len(row[countryPos]) > 0
}

// SplitResult split result by >
func SplitResult(result string) (string, string) {
	searchAndCorrectedTerm := strings.Split(result, ">")
	searchTerm := Utils.TrimLeftRightSpace(searchAndCorrectedTerm[0])
	correctedTerm := ""
	if len(searchAndCorrectedTerm) == 2 {
		correctedTerm = Utils.TrimLeftRightSpace(searchAndCorrectedTerm[1])
	}

	return searchTerm, correctedTerm
}

// ReadSearchCsvReport from file excel
func ReadSearchCsvReport(inputCsvFileName string) []DataFormat.ReportRow {
	file, _ := os.Open(inputCsvFileName)
	defer file.Close()
	csvr := csv.NewReader(file)
	csvHeader := ReadHeader(csvr)
	return ReadCsvReportRow(csvr, csvHeader)
}

// ReadCsvReportRow from file
func ReadCsvReportRow(csvr *csv.Reader, fieldMap map[string]int) []DataFormat.ReportRow {
	csvRows := []DataFormat.ReportRow{}
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			return csvRows
		}

		csvRow := DataFormat.ReportRow{
			Label:                row[fieldMap[label]],
			OrginalSearchTerm:    row[fieldMap[orgSearchTerm]],
			OrginalCorrectedTerm: row[fieldMap[orgCorrectedTerm]],
			SearchTerm:           row[fieldMap[searchTerm]],
			CorrectTerm:          row[fieldMap[correctTerm]],
		}
		csvRows = append(csvRows, csvRow)
	}
}
