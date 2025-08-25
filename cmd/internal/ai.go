package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Message ResponseMessage `json:"message"`
}

type APIResponse struct {
	id       string   `json:"id"`
	chatcmpl string   `json:"chatcmpl"`
	object   string   `json:"object"`
	created  string   `json:"created"`
	choices  []Choice `json:"choices"`
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
		log.Fatal("salve")
	}

	reply, err := http.Post("", "application/json", bytes.NewBuffer(data))

	body, err := io.ReadAll(reply.Body)
	if err != nil {
		panic(err)
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatal("error to parse reply body to json")
	}
	return apiResponse, nil
}
