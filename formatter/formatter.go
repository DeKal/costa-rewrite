package formatter

import (
	"fmt"
	"strconv"

	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
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
			EvaluationLabel(autoCorrectRow.OriginCorrectTerm, response.CorrectTerm),
		},
	}
}

// EvaluationLabel return evaluation label
func EvaluationLabel(originCorrectTerm string, correctTerm string) string {
	isExpectedEqualAnalyzing := originCorrectTerm == correctTerm
	if len(correctTerm) == 0 {
		return ""
	}
	if isExpectedEqualAnalyzing {
		return "Correct Label"
	}
	return "Different Label"
}

// FormatLinkPattern return evaluation label
func FormatLinkPattern(host string, searchTerm string) string {
	rewriteLinkPattern := "%s/_c/v1/search/rewrite/?q=%s&lang=en&segment=women"
	return fmt.Sprintf(rewriteLinkPattern, host, searchTerm)
}
