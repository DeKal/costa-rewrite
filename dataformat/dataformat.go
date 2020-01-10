package dataformat

// AutoCorrectRow format
type AutoCorrectRow struct {
	OriginSearchTerm  string
	OriginCorrectTerm string
	Rating            string
}

// RewriteResponse format
type RewriteResponse struct {
	SearchTerm  string `json:"search_term"`
	CorrectTerm string `json:"correct_term"`
	Count       int    `json:"count"`
}

// CommandLineArgs format
type CommandLineArgs struct {
	Mode         string
	CompareFile1 string
	CompareFile2 string
	CsvInput     string
	CsvOutput    string
	RewriteHost  string
	Country      string
}

// ReportRow format
type ReportRow struct {
	Label                string
	OrginalSearchTerm    string
	OrginalCorrectedTerm string
	SearchTerm           string
	CorrectTerm          string
}
