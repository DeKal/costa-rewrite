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
  CsvInput  string 
  CsvOutput string 
	RewriteLinkPattern string  
	Country string
}
