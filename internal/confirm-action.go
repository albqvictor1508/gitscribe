package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmAction(msg string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------------------ Commit Message ---------------------")
	fmt.Println(msg)
	fmt.Println("-------------------------------------------------------")
	fmt.Print("Deseja continuar? [y/N]: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	fmt.Print(input)

	return input == "y"
}
