package normal

import (
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Formatter "github.com/DeKal/costa-rewrite/formatter"
	Parser "github.com/DeKal/costa-rewrite/parser"
	Reader "github.com/DeKal/costa-rewrite/reader"
	Writer "github.com/DeKal/costa-rewrite/writer"
)

var csvHeader = [][]string{
	{
		"Label",
		"Orginal Search Term",
		"Original Corrected Term",
		"Search Term",
		"Correct Term",
		"Count",
		"Correction",
	},
}

// RunRewriteAndProduceReports Normal mode
func RunRewriteAndProduceReports(params DataFormat.CommandLineArgs) {
	autoCorrects := Reader.ReadSearchTermsFromExcel(params.CsvInput, params.Country)
	csvContent := [][]string{}
	for _, autoCorrect := range autoCorrects {
		searchURL := Formatter.FormatLinkPattern(params.RewriteHost, autoCorrect.OriginSearchTerm)
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
	csvRow := Formatter.FormatCsvRow(response, autoCorrectRow)
	return append(csvContent, csvRow...)
}
