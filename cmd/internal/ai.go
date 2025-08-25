package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResponseMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Messages []ResponseMessages `json:"messages"`
}

type APIResponse struct {
	/*
		id       string   `json:"id"`
		chatcmpl string   `json:"chatcmpl"`
		object   string   `json:"object"`
		created  string   `json:"created"`
	*/
	choices []Choice `json:"choices"`
}

func SendPrompt(ctx string) (APIResponse, error) {
	data, err := json.Marshal(map[string]any{
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": ctx,
			},
		},
	})
	if err != nil {
		log.Fatal("error to create body")
	}

	reply, err := http.Post("https://ai.hackclub.com/chat/completions", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return APIResponse{}, fmt.Errorf("error to request to hackclub: %v", err)
	}
	defer reply.Body.Close()

	body, err := io.ReadAll(reply.Body)
	if err != nil {
		return APIResponse{}, fmt.Errorf("error to read data: %v", err)
	}

	if reply.StatusCode != http.StatusOK {
		log.Printf("API retornou um status de erro: %s", reply.Status)
		log.Printf("Corpo da resposta de erro: %s", body)
		return APIResponse{}, fmt.Errorf("a API retornou um status não-OK: %s", reply.Status)
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatal("error to parse reply body to json")
	}
	return apiResponse, nil
}
