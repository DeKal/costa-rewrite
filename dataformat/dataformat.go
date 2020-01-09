package dataformat

import (
	"strconv"
)

// AutoCorrectRow format
type AutoCorrectRow struct {
  OriginSearchTerm  string
  OriginCorrectTerm string
  Rating            string
}

// RewriteResponse format
type RewriteResponse struct {
  SearchTerm  string `json:“search_term”`
  CorrectTerm string `json:“correct_term”`
  Count       int    `json:“count”`
}

// FormatCsvRow return rewrite response as row
func FormatCsvRow(response RewriteResponse, autoCorrectRow AutoCorrectRow) [][]string {
	return [][]string{
		{
			autoCorrectRow.Rating,
			autoCorrectRow.OriginSearchTerm,
			autoCorrectRow.OriginCorrectTerm,
			response.SearchTerm,
			response.CorrectTerm,
			strconv.FormatInt(int64(response.Count), 10),
		},
	}
}




