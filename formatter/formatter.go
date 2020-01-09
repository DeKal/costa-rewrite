package formatter

import (
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	"strconv"
)

// FormatCsvRow return rewrite response as row
func FormatCsvRow(response DataFormat.RewriteResponse, autoCorrectRow DataFormat.AutoCorrectRow) [][]string {
	return [][]string{
		{
			autoCorrectRow.Rating,
			autoCorrectRow.OriginSearchTerm,
			autoCorrectRow.OriginCorrectTerm,
			response.SearchTerm,
			response.CorrectTerm,
			strconv.FormatInt(int64(response.Count), 10),
			evaluationLabel(autoCorrectRow.OriginCorrectTerm, response.CorrectTerm),
		},
	}
}

func evaluationLabel(originCorrectTerm string, correctTerm string) string {
	isExpectedEqualAnalyzing := originCorrectTerm == correctTerm
	if len(correctTerm) == 0 {
		return ""
	}
	if isExpectedEqualAnalyzing {
		return "Correct Label"
	}
	return "Different Label"
}
