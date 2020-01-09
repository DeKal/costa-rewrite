package main

import (
	"fmt"

	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Parser "github.com/DeKal/costa-rewrite/parser"
	Reader "github.com/DeKal/costa-rewrite/reader"
	Writer "github.com/DeKal/costa-rewrite/writer"
)

var csvHeader = [][]string{
	{"Label", "Orginal Search Term", "Original Corrected Term", "Search Term", "Correct Term", "Count"},
}

func main() {
	params := Parser.ParseCommandLineParams()

	autoCorrects := Reader.ReadSearchTermsFromExcel(params.CsvInput)
	csvContent := [][]string{}
	for _, autoCorrect := range autoCorrects {
		searchTerm := autoCorrect.OriginSearchTerm
		searchURL := fmt.Sprintf(params.RewriteLinkPattern, searchTerm)
		resp, err := Parser.Get(searchURL)
		if err == nil {
			response := Parser.Parse(resp)
			csvContent = AddResultToCsvContent(csvContent, response, autoCorrect)
		}
	}
	Writer.WriteCsvFile(params.CsvOutput, csvHeader, csvContent)
}

// AddResultToCsvContent return new CSV content after appending rewrite response
func AddResultToCsvContent(csvContent [][]string, response DataFormat.RewriteResponse, autoCorrectRow DataFormat.AutoCorrectRow) [][]string {
	csvRow := DataFormat.FormatCsvRow(response, autoCorrectRow)
	return append(csvContent, csvRow...)
}
