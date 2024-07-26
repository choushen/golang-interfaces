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

	if key == "" || token == "" {
		fmt.Println("API_KEY and API_TOKEN are required")
		return
	}

	url := fmt.Sprintf("https://api.trello.com/1/members/%s/boards?key=%s&token=%s&filter=open", memberID, key, token)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll((resp.Body))
	if err != nil {
		a
		fmt.Println(err)
		return
	}

	var boards BoardList
	err = json.Unmarshal(body, &boards)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, board := range boards {
		fmt.Println(board.Name)
	}

} // Path: go.mod
