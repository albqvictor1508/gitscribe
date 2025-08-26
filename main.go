package main

import (
	"fmt"
	"log"
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

			fmt.Println("ADDING...")
			r := exec.Command("git", "add", filepath)
			if _, err := r.Output(); err != nil {
				panic(err)
			}

			fmt.Println("COMMITING...")
			raw := exec.Command("git", "commit", "-m", message)
			if _, err := raw.Output(); err != nil {
				panic(err)
			}

			fmt.Println("CHECKING REMOTE...")
			remote := exec.Command("git", "remote", "-v")
			res, err := remote.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(res))

			fmt.Println("PUSHING...")
			rawData := exec.Command("git", "push", "origin", branch)
			if _, err := rawData.Output(); err != nil {
				log.Fatalf("error to push commit: %v", err)
			}

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
