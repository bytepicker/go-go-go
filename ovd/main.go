package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func performGetRequest(apiURL string, params map[string]interface{}) ([]byte, error) {
	query := url.Values{}
	for key, value := range params {
		switch v := value.(type) {
		case string:
			query.Add(key, v)
		default:
			jsonValue, err := json.Marshal(value)
			if err != nil {
				return nil, fmt.Errorf("Error encoding JSON for key %s: %s", key, err)
			}
			query.Add(key, string(jsonValue))
		}
	}

	fullURL := fmt.Sprintf("%s?%s", apiURL, query.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	apiURL := "https://api.repression.info/v1/data"
	params := map[string]interface{}{
		"language": "en",
		"filter": map[string]interface{}{
			"persecution_subject_en":   "freedom of assembly",
			"persecution_started_year": []interface{}{2022, 2023},
			"case_group_en":            "Anti-war case",
		},
	}

	// Perform GET request
	response, err := performGetRequest(apiURL, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse and display JSON
	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Total entries:", result["total"])
	fmt.Println("Received JSON:")
	prettyResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error pretty printing JSON:", err)
		return
	}
	fmt.Println(string(prettyResult))

	// Wait for user input before exiting
	fmt.Println("Press Enter to exit.")
	fmt.Scanln()
}
