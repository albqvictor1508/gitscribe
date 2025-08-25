package main

import (
	"fmt"
	"log"

	"github.com/albqvictor1508/gitscribe/cmd/internal"
)

func main() {
	// cmd := exec.Command("git", "diff")

	/*
		 output, err := cmd.Output()
			if err != nil {
				log.Fatal("Error to exec git diff")
			}
	*/

	context := "who wins the 2002 world cup"
	reply, err := internal.SendPrompt(context)
	if err != nil {
		log.Fatalf("error to get reply: %v", err)
	}

	fmt.Print(reply)
}
