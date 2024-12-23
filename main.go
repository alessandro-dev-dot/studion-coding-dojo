package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Testando nossa llm:")
	data := []byte(`{
		"model": "llama3.2",
		"prompt": "Why is the sky blue?",
		"stream": false
	}`)
	resp, err := http.Post("http://localhost:11434/api/generate/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type llmResponse struct {
		CreatedAt string `json:created_at`
		Response  string `json:response`
	}
	var response llmResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response.Response)
}
