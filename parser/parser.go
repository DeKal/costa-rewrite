package parser

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
)

// Get Request to an Url
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}
	return "", nil
}

// Parse resonse from string
func Parse(response string) DataFormat.RewriteResponse {
	data := DataFormat.RewriteResponse{}
	json.Unmarshal([]byte(response), &data)
	return data
}

// ParseCommandLineParams parse args from cmd
func ParseCommandLineParams() DataFormat.CommandLineArgs {
	const (
		defaultCsvInput           = "example_input.csv"
		defaultCsvOutput          = "output.csv"
		defaultRewriteLinkPattern = "http://localhost:9999/_c/v1/search/rewrite/?q=%s&lang=en&segment=women"
	)

	var rewriteURL string
	flag.StringVar(&rewriteURL, "rewriteUrl", defaultRewriteLinkPattern, "rewrite url")

	var csvInput string
	flag.StringVar(&csvInput, "inputName", defaultCsvInput, "an Input name for reading data")

	var csvOutput string
	flag.StringVar(&csvOutput, "outputName", defaultCsvOutput, "an Output name for writing data")

	flag.Parse()

	return DataFormat.CommandLineArgs{
		CsvInput:           csvInput,
		CsvOutput:          csvOutput,
		RewriteLinkPattern: rewriteURL,
	}
}
