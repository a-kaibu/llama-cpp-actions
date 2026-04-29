package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	reqBody, err := json.Marshal(map[string]any{
		"messages": []map[string]string{
			{"role": "user", "content": "こんにちは"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://127.0.0.1:8080/v1/chat/completions", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status %d: %s", resp.StatusCode, body)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	if len(result.Choices) == 0 {
		log.Fatal("no choices returned")
	}

	fmt.Println(result.Choices[0].Message.Content)
}
