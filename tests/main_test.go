package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCovidSummary(t *testing.T) {
	mockJsonRequest, err := ioutil.ReadFile("mockJsonRequest.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/covid/summary", bytes.NewBuffer(mockJsonRequest))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	}
	fmt.Printf("Response: %s\n", body)

	mockJsonResponse, err := ioutil.ReadFile("mockJsonResponse.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	if string(body) != string(mockJsonResponse) {
		t.Errorf("Response should be\n %s but have\n %s", string(mockJsonResponse), string(body))
	}
}
