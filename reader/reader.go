package reader

import (
	"encoding/csv"
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Utils "github.com/DeKal/costa-rewrite/utils"
	"io"
	"log"
	"os"
	"strings"
)

const (
	rating      = "Autocorrect Rating"
	labelResult = "Event Label"
	hk          = "HK"
	id          = "ID"
	my          = "MY"
	ph          = "PH"
	sh          = "SG"
	tw          = "TW"
)

// ReadSearchTermsFromExcel from file excel
func ReadSearchTermsFromExcel(inputCsvFileName string) []DataFormat.AutoCorrectRow {
	file, _ := os.Open(inputCsvFileName)
	defer file.Close()

	csvr := csv.NewReader(file)
	csvHeader := ReadHeader(csvr)
	log.Println(csvHeader)
	return ReadSearchTerms(csvr, csvHeader)
}

// ReadSearchTerms from file
func ReadSearchTerms(csvr *csv.Reader, fieldMap map[string]int) []DataFormat.AutoCorrectRow {
	autoCorrects := []DataFormat.AutoCorrectRow{}
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			return autoCorrects
		}
		if isSG(row) {

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

func isSG(row []string) bool {
	return (len(row[6]) > 0)
}
