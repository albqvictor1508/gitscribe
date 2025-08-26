package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/albqvictor1508/gitscribe/internal"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
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
			pterm.DefaultBigText.WithLetters(
				putils.LettersFromStringWithStyle("Git", pterm.NewStyle(pterm.FgCyan)),
				putils.LettersFromStringWithStyle("Scribe", pterm.NewStyle(pterm.FgLightMagenta))).
				Render()
			pterm.Info.Println("Seu assistente de commits com IA.")
			pterm.Println()
			time.Sleep(time.Second)

			files := args
			if len(files) == 0 {
				files = append(files, ".")
			}
			addSpinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Adicionando arquivos: %s...", strings.Join(files, " ")))

			for _, file := range files {
				r := exec.Command("git", "add", file)
				if _, err := r.Output(); err != nil {
					addSpinner.Fail(fmt.Sprintf("Falha ao adicionar o arquivo %s: %v", file, err))
					os.Exit(1)
				}
			}
			addSpinner.Success("Arquivos adicionados ao stage!")

			if len(message) == 0 {
				aiSpinner, _ := pterm.DefaultSpinner.Start("Analisando as mudanças e gerando mensagem com IA...")
				diff := exec.Command("git", "diff", "--staged") // Use --staged for better accuracy
				res, err := diff.Output()
				if err != nil {
					aiSpinner.Fail("Falha ao obter o 'git diff': ", err)
					panic(err)
				}

				if len(res) == 0 {
					aiSpinner.Warning("Nenhuma mudança encontrada no stage. Nada para commitar.")
					os.Exit(0)
				}

				context := fmt.Sprintf("aqui está a diferença no código do usuário, em cima dessa diferença, quero que crie uma mensagem de commit que siga os padrões estabelecidos pelo 'Conventional Commit', (chore, feat, entre outros) . Além disso quero que me retorne somente a mensagem de commit, nada além disso, quero que retorne somente a mensagem de commit: %v", string(res))

				msg, err := internal.SendPrompt(context)
				if err != nil {
					aiSpinner.Fail(fmt.Sprintf("Erro ao gerar mensagem com IA: %v", err))
					os.Exit(1)
				}
				message = msg
				aiSpinner.Success("Mensagem de commit gerada!")
			}

			if !internal.ConfirmAction(message) {
				pterm.Warning.Println("Operação cancelada pelo usuário.")
				os.Exit(1)
			}

			commitSpinner, _ := pterm.DefaultSpinner.Start("Realizando commit...")
			raw := exec.Command("git", "commit", "-m", message)
			if output, err := raw.CombinedOutput(); err != nil {
				commitSpinner.Fail(fmt.Sprintf("Erro ao commitar: %s", string(output)))
				os.Exit(1)
			}
			commitSpinner.Success("Commit realizado com sucesso!")

			pushSpinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Executando 'git push origin %s'...", branch))
			pushCmd := exec.Command("git", "push", "origin", branch)

			if output, err := pushCmd.CombinedOutput(); err != nil {
				pushSpinner.Fail(fmt.Sprintf("Erro ao realizar push: %s", string(output)))
				os.Exit(1)
			}
			pushSpinner.Success("Push realizado com sucesso!")
			pterm.Println()
			pterm.Success.Println("Tudo pronto!")
		},
	}

	cmd.Flags().StringVarP(&message, "message", "m", "", "Messagem do commit")
	cmd.Flags().StringVarP(&branch, "branch", "b", "main", "Branch para o push")

	rootCmd.AddCommand(cmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

