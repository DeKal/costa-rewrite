package reader

import (
	"encoding/csv"
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Utils "github.com/DeKal/costa-rewrite/utils"
	"io"
	"os"
	"strings"
)

// ReadSearchTermsFromExcel from file
func ReadSearchTermsFromExcel(inputCsvFileName string) []DataFormat.AutoCorrectRow {
	file, _ := os.Open(inputCsvFileName)
	defer file.Close()

	csvr := csv.NewReader(file)

	autoCorrects := []DataFormat.AutoCorrectRow{}
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			return autoCorrects
		}
		if isSG(row) {
			autoCorrect := DataFormat.AutoCorrectRow{}
			autoCorrect.Rating = row[0]

			searchAndCorrectedTerm := strings.Split(row[1], ">")
			autoCorrect.OriginSearchTerm = Utils.TrimLeftRightSpace(searchAndCorrectedTerm[0])

			autoCorrect.OriginCorrectTerm = ""
			if len(searchAndCorrectedTerm) == 2 {
				autoCorrect.OriginCorrectTerm = Utils.TrimLeftRightSpace(searchAndCorrectedTerm[1])
			}

			autoCorrects = append(autoCorrects, autoCorrect)
		}

	}
}

func isSG(row []string) bool {
	return (len(row[6]) > 0)
}
