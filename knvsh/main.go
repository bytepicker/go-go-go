package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	url := "http://knvsh.gov.spb.ru/gosuslugi/svedeniya2/"
	searchTerm := "YOUR_NUMBER"
	startTime := time.Now()

	fmt.Println(time.Now().Format(time.RFC3339), "Start polling results page")

	for {
		// Fetch the webpage
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(time.Now().Format(time.RFC3339), "Error fetching webpage:", err)
			continue // Try again
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(time.Now().Format(time.RFC3339), "Error reading response body:", err)
			continue // Try again
		}

		// Check if the body contains the search term
		if strings.Contains(string(body), searchTerm) {
			fmt.Println("Ready!")
			endTime := time.Now()
			elapsedTime := endTime.Sub(startTime)
			fmt.Println(time.Now().Format(time.RFC3339), "Ready")
			fmt.Println("Elapsed Time:", elapsedTime)
			break // Exit loop if found
		}

		// Sleep for 5 seconds before checking again
		time.Sleep(5 * time.Second)
	}

	// Wait for user input before exiting
	fmt.Println("Press Enter to exit.")
	fmt.Scanln()
}
