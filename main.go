package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/albqvictor1508/gitscribe/internal"
	"github.com/spf13/cobra"
)

/*
Copyright © 2025 Victor Albuquerque albq.victor@gmail.com
*/

func main() {
	var message, branch string
	rootCmd := &cobra.Command{Use: "gs"}

	cmd := &cobra.Command{
		Use:   "cmt [files]",
		Args:  cobra.MinimumNArgs(0),
		Short: "realiza 'git add [path]', 'git commit -m [message]', 'git push origin [branch]'",
		Run: func(cmd *cobra.Command, args []string) {
			files := args
			if len(files) == 0 {
				files = append(files, ".")
			}
			if len(message) == 0 {
				diff := exec.Command("git", "diff")
				res, err := diff.Output()
				if err != nil {
					panic(err)
				}

				context := fmt.Sprintf("aqui está a diferença no código do usuário, em cima dessa diferença, quero que crie uma mensagem de commit que siga os padrões estabelecidos pelo 'Conventional Commits'. Além disso quero que me retorne somente a mensagem de commit, nada além disso, quero que retorne somente a mensagem de commit: %v", string(res))

				msg, err := internal.SendPrompt(context)
				if err != nil {
					log.Fatalf("error to get message with ai: %v", err)
				}
				message = msg
			}

			for _, file := range files {
				fmt.Printf("ADDING...\n %v", file)
				r := exec.Command("git", "add", file)
				if _, err := r.Output(); err != nil {
					panic(err)
				}
			}

			fmt.Println("COMMITING...")
			raw := exec.Command("git", "commit", "-m", message)
			if _, err := raw.Output(); err != nil {
				fmt.Println(err.Error())
				log.Fatalf("error to commit: %v", err.Error())
			}

			fmt.Println("CHECKING REMOTE...")
			remoteRaw := exec.Command("git", "remote", "-v")
			res, err := remoteRaw.Output()
			if err != nil {
				panic(err)
			}

			remote := string(res)
			if len(remote) == 0 {
				log.Fatal("Please set a remote branch...")
			}

			fmt.Printf("remote branch: %v\n", remote)

			fmt.Println("PUSHING...")
			rawData := exec.Command("git", "push", "origin", branch)
			if _, err := rawData.Output(); err != nil {
				log.Fatalf("error to push commit: %v", err)
			}

			fmt.Printf("message: %v \n", message)
			fmt.Printf("branch: %v \n", branch)
		},
	}

	cmd.Flags().StringVarP(&message, "message", "m", "", "Messagem do commit")
	cmd.Flags().StringVarP(&branch, "branch", "b", "main", "Branch")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
