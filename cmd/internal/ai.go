package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func SendPrompt(ctx *string) ([]byte, error) {
	data, err := json.Marshal(map[string]any{
		"message": []map[string]string{
			{
				"role":    "user",
				"content": *ctx,
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

	if err := json.Unmarshal(body, ""); err != nil {
		log.Fatal("error to parse reply body to json")
	}
	return body, nil
}
