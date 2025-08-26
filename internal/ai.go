package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int             `json:"index"`
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
}

type APIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

func SendPrompt(ctx string) (string, error) {
	data, err := json.Marshal(map[string]any{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Você é um gerador de commits, que gera commits em inglês seguindo o padrão Conventional Commit",
			},
			{
				"role":    "user",
				"content": ctx,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("error to create body: %v", err)
	}

	reply, err := http.Post("https://ai.hackclub.com/chat/completions", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("error to request to hackclub: %v", err)
	}
	defer reply.Body.Close()

	body, err := io.ReadAll(reply.Body)
	if err != nil {
		return "", fmt.Errorf("error to read data: %v", err)
	}

	if reply.StatusCode != http.StatusOK {
		return "", fmt.Errorf("a API retornou um status não-OK: %s", reply.Status)
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", fmt.Errorf("error to parse response body: %v", err)
	}

	// WARN: meio perigoso mas tá pegando
	msg := apiResponse.Choices[0].Message.Content
	lines := strings.Split(msg, "\n")
	commit := lines[len(lines)-1]
	fmt.Println(commit)

	return commit, nil
}
