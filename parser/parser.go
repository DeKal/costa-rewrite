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
		defaultMode        = "normal"
		defaultFileName1   = "output"
		defaultFileName2   = "output-2"
		defaultCsvInput    = "example_input.csv"
		defaultCsvOutput   = "output.csv"
		defaultRewriteHost = "http://localhost:9999"
		defaultCountryCode = "SG"
	)
	var mode string
	flag.StringVar(&mode, "mode", defaultMode, "mode: [normal, compare]")

	var cmpFile1 string
	flag.StringVar(&cmpFile1, "file1", defaultFileName1, "file name 1")
	var cmpFile2 string
	flag.StringVar(&cmpFile2, "file2", defaultFileName2, "file name 2")

	var rewriteHost string
	flag.StringVar(&rewriteHost, "rewriteHost", defaultRewriteHost, "rewrite host")

	var csvInput string
	flag.StringVar(&csvInput, "inputName", defaultCsvInput, "an Input name for reading data")

	var csvOutput string
	flag.StringVar(&csvOutput, "outputName", defaultCsvOutput, "an Output name for writing data")

	var country string
	flag.StringVar(&country, "country", defaultCountryCode, "country filter for some specific word [HK, ID, MY, PH, SG, TW]")

	flag.Parse()

	return DataFormat.CommandLineArgs{
		Mode:         mode,
		CompareFile1: cmpFile1,
		CompareFile2: cmpFile2,
		CsvInput:     csvInput,
		CsvOutput:    csvOutput,
		RewriteHost:  rewriteHost,
		Country:      country,
	}
}

// IsNormalMode Check if mode is normal
func IsNormalMode(mode string) bool {
	return mode == "normal"
}
