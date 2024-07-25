package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	type Board struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	type BoardList []Board

	key := os.Getenv("TRELLO_API_KEY")
	token := os.Getenv("TRELLO_API_TOKEN")
	memberID := os.Getenv("TRELLO_MEMBER_ID")

	fmt.Println("API_KEY: ", key)
	fmt.Println("API_TOKEN: ", token)
	fmt.Println("MEMBER_ID: ", memberID)

	if key == "" || token == "" {
		fmt.Println("API_KEY and API_TOKEN are required")
		return
	}

	url := fmt.Sprintf("https://api.trello.com/1/members/%s/boards?key=%s&token=%s&filter=open", memberID, key, token)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		body, _ := io.ReadAll(io.Reader(resp.Body))
		fmt.Printf("Response body: %s\n", body)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// fmt.Println("Response body:", string(body))

	var boards BoardList
	if err := json.Unmarshal(body, &boards); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the parsed boards
	fmt.Println("Boards:")
	for _, board := range boards {
		fmt.Printf("ID: %s, Name: %s\n", board.ID, board.Name)
	}

}
