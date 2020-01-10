package compare

import (
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	Formatter "github.com/DeKal/costa-rewrite/formatter"
	Reader "github.com/DeKal/costa-rewrite/reader"
	Writer "github.com/DeKal/costa-rewrite/writer"
)

var csvHeader = [][]string{
	{
		"Label",
		"Orginal Search Term",
		"Original Corrected Term",
		"Search Term For Report 1",
		"Correct Term For Report 1",
		"Search Term For Report 2",
		"Correct Term For Report 2",
		"Correction",
	},
}

// Compare2Report Compare 2 reports
func Compare2Report(params DataFormat.CommandLineArgs) {
	csvFile1 := Reader.ReadSearchCsvReport(params.CompareFile1)
	csvFile2 := Reader.ReadSearchCsvReport(params.CompareFile2)
	csvContent := Compare(csvFile1, csvFile2)
	Writer.WriteCsvFile(params.CsvOutput, csvHeader, csvContent)
}

// Compare 2 reports and produce new reports
func Compare(report1 []DataFormat.ReportRow, report2 []DataFormat.ReportRow) [][]string {
	rows := [][]string{}
	for index, reportRow := range report1 {
		searchTerm := reportRow.SearchTerm
		row := []string{
			reportRow.Label,
			reportRow.OrginalSearchTerm,
			reportRow.OrginalCorrectedTerm,
			searchTerm,
			reportRow.CorrectTerm,
			searchTerm,
			report2[index].CorrectTerm,
			Formatter.EvaluationLabel(reportRow.CorrectTerm, report2[index].CorrectTerm),
		}

		rows = append(rows, row)
	}
	return rows
}
