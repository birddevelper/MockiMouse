package config

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (scenario Scenario) GetResponse() (string, error) {

	responseCount := len(scenario.Response)

	if responseCount == 0 {
		return "No response", nil
	}

	if responseCount == 1 {
		return scenario.Response[0], nil
	}

	// select random index in response array
	rand.Seed(time.Now().UnixNano())
	selectedResponseIndex := rand.Intn(responseCount)

	//return the randomly selected response
	return scenario.resolveResponse(selectedResponseIndex)

}

func (scenario Scenario) resolveResponse(responseIndex int) (string, error) {
	const FILE_PREFIX = "file://"
	const URL_PREFIX = "http"
	const RESPONSE_FILE_FOLDER = "responses/"
	if strings.HasPrefix(scenario.Response[responseIndex], FILE_PREFIX) {
		response, err := readFromFile(RESPONSE_FILE_FOLDER + scenario.Response[0][7:])
		if err != nil {
			return "", err
		}
		return string(response), nil
	} else if strings.HasPrefix(scenario.Response[responseIndex], URL_PREFIX) {
		response, err := readFromUrl(scenario.Response[responseIndex])
		if err != nil {
			return "", err
		}
		// return response body
		return response, nil
	}
	// if the response is a simple text return it
	return scenario.Response[responseIndex], nil

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
