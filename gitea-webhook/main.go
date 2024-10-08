package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebhookPayload struct {
	Secret     string     `json:"secret,omitempty"`
	Ref        string     `json:"ref,omitempty"`
	Before     string     `json:"before,omitempty"`
	After      string     `json:"after,omitempty"`
	CompareURL string     `json:"compare_url,omitempty"`
	Repository Repository `json:"repository"`
	Pusher     Pusher     `json:"pusher"`
	Sender     Sender     `json:"sender"`
	Commits    []Commit   `json:"commits"`
}

type Repository struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    User   `json:"owner"`
	Private  bool   `json:"private"`
	CloneURL string `json:"clone_url"`
	HTMLURL  string `json:"html_url"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Sender struct {
	Login string `json:"login"`
	ID    int64  `json:"id"`
}

type Commit struct {
	ID        string   `json:"id"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    User     `json:"author"`
	Committer User     `json:"committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Check if it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the JSON payload into our struct
	var payload WebhookPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the push details
	fmt.Printf("Received push event on ref: %s\n", payload.Ref)
	fmt.Printf("Pushed by: %s\n", payload.Pusher.Name)
	fmt.Printf("Repository: %s\n", payload.Repository.FullName)
	fmt.Println("Commits:")

	for _, commit := range payload.Commits {
		fmt.Printf("- Commit ID: %s\n", commit.ID)
		fmt.Printf("  Message: %s\n", commit.Message)
		fmt.Printf("  URL: %s\n", commit.URL)
	}

	// Respond to the webhook with a success status
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook received successfully")

	// TODO! build
}

func main() {
	// Register the webhook handler at the "/webhook" endpoint
	http.HandleFunc("/webhook", handleWebhook)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
