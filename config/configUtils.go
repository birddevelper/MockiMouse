package config

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (scenario Scenario) GetResponse() (string, error) {
	const FILE_PREFIX = "file://"
	const URL_PREFIX = "http"
	const RESPONSE_FILE_FOLDER = "responses/"
	if strings.HasPrefix(scenario.Response, FILE_PREFIX) {
		response, err := readFromFile(RESPONSE_FILE_FOLDER + scenario.Response[7:])
		if err != nil {
			return "", err
		}
		return string(response), nil
	} else if strings.HasPrefix(scenario.Response, URL_PREFIX) {
		response, err := readFromUrl(scenario.Response)
		if err != nil {
			return "", err
		}
		// return response body
		return response, nil
	}
	// if the response is a simple text return it
	return scenario.Response, nil

}

func readFromUrl(url string) (string, error) {
	// call remote address
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	// close response body after reading whole data
	defer response.Body.Close()

	return string(responseBody), err
}

func readFromFile(filePath string) (string, error) {
	filename, _ := filepath.Abs(filePath)
	response, err := os.ReadFile(filename)

	return string(response), err
}
