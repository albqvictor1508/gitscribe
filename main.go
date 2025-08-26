package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

/*
Copyright © 2025 Victor Albuquerque albq.victor@gmail.com
*/

func main() {
	var message, branch, filepath string
	rootCmd := &cobra.Command{Use: "gs"}

	cmd := &cobra.Command{
		Use: "cmt",
		Run: func(cmd *cobra.Command, args []string) {
			if len(branch) == 0 {
				branch = "main"
			}

			if len(message) == 0 {
				diff := exec.Command("git", "diff")
				res, err := diff.Output() // TODO: jogar esse git diff na IA
				if err != nil {
					panic(err)
				}

				message = fmt.Sprintf("mensagem e parara, aqui está a diferença no código: %v", string(res))
			}

			r := exec.Command("git", "add", filepath)
			response, err := r.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(response))

			raw := exec.Command("git", "commit", "-m", message)
			t, err := raw.Output()

			fmt.Println(string(t))
			fmt.Printf("message: %v \n", message)
			fmt.Printf("branch: %v \n", branch)
		},
	}

	cmd.Flags().StringVarP(&filepath, "", "", ".", "")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Messagem do commit")
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
