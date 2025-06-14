package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := args[1]
	fmt.Printf("Fetching recent activity for user: %s\n", username)

	events, err := fetchUserActivity(username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printEvent(events)
}

func fetchUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("user not found or invalid username")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch activity. HTTP Status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %w", err)
	}

	return events, nil
}

func printEvent(events []Event) {
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed code to %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("- Opened a pull request on %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Created an issue on %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Started watching %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s on %s\n", event.Type, event.Repo.Name)
		}
	}
}
