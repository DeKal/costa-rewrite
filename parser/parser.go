package parser

import (
	"encoding/json"
	DataFormat "github.com/DeKal/costa-rewrite/dataformat"
	"io/ioutil"
	"log"
	"net/http"
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
