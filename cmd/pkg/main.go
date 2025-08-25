package main

import (
	"fmt"

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

	context := "baseado nos dados desse git diff, me fala oque foi alterado"
	reply, err := internal.SendPrompt(context)
	if err != nil {
		panic(err)
	}

	fmt.Print(reply)
}
