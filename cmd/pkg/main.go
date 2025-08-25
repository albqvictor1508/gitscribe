package main

import (
	"fmt"
	"io"

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
	var reply io.ReadCloser = internal.SendPrompt(context)
	fmt.Print(reply)
}
