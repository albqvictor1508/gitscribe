package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/albqvictor1508/gitscribe/internal"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func main() {
	var message, branch string

	rootCmd := &cobra.Command{Use: "gitscribe"}

	cmd := &cobra.Command{
		Use:   "cmt [files]",
		Args:  cobra.MinimumNArgs(0),
		Short: "AI-powered git add, commit, and push",
		Run: func(cmd *cobra.Command, args []string) {
			asciiArt2 := `
           /$$   /$$                                  /$$ /$$                
          |__/  | $$                                 |__/| $$                
  /$$$$$$  /$$ /$$$$$$   /$$$$$$$  /$$$$$$$  /$$$$$$  /$$| $$$$$$$   /$$$$$$ 
 /$$__  $$| $$|_  $$_/  /$$_____/ /$$_____/ /$$__  $$| $$| $$__  $$ /$$__  $$
| $$  \ $$| $$  | $$   |  $$$$$$ | $$      | $$  \__/| $$| $$  \ $$| $$$$$$$$
| $$  | $$| $$  | $$ /$$\____  $$| $$      | $$      | $$| $$  | $$| $$_____/
|  $$$$$$$| $$  |  $$$$//$$$$$$$/|  $$$$$$$| $$      | $$| $$$$$$$/|  $$$$$$$
 \____  $$|__/   \___/ |_______/  \_______/|__/      |__/|_______/  \_______/
 /$$  \ $$                                                                   
|  $$$$$$/                                                                   
 \______/                                                                    
			`
			pterm.DefaultBasicText.Println(pterm.FgGreen.Sprint(asciiArt2))
			pterm.Info.Println("Your AI-powered commit assistant.")
			time.Sleep(time.Second)

			files := args
			if len(files) == 0 {
				files = append(files, ".")
			}
			addSpinner, _ := pterm.DefaultSpinner.WithSequence("|", "/", "-", "\\ ").Start("Staging files...")
			for _, file := range files {
				addCmd := exec.Command("git", "add", file)
				addCmd.Stdout = io.Discard
				addCmd.Stderr = io.Discard
				if err := addCmd.Run(); err != nil {
					addSpinner.Fail(fmt.Sprintf("Failed to stage file %s: %v", file, err))
					os.Exit(1)
				}
			}
			addSpinner.Success("Files staged successfully!")

			if len(message) == 0 {
				aiSpinner, _ := pterm.DefaultSpinner.WithSequence("|", "/", "-", "\\ ").Start("Analyzing changes and generating message with AI...")

				var diffOutput bytes.Buffer
				diffCmd := exec.Command("git", "diff", "--staged")
				diffCmd.Stdout = &diffOutput
				diffCmd.Stderr = &diffOutput
				if err := diffCmd.Run(); err != nil {
					aiSpinner.Fail(fmt.Sprintf("Failed to get git diff: %s", diffOutput.String()))
					os.Exit(1)
				}

				if diffOutput.Len() == 0 {
					aiSpinner.Warning("No changes found in stage. Nothing to commit.")
					os.Exit(0)
				}

				context := fmt.Sprintf("Based on the git diff below, create a commit message that follows the 'Conventional Commits' specification. Return only the commit message, with nothing else: %v", diffOutput.String())

				msg, err := internal.SendPrompt(context)
				if err != nil {
					aiSpinner.Fail(fmt.Sprintf("Error generating message with AI: %v", err))
					os.Exit(1)
				}
				message = msg
				aiSpinner.Success("Commit message generated!")
			}

			if !internal.ConfirmAction(message) {
				// TODO: talvez colocar dps uma frasezinha
				os.Exit(1)
			}

			commitSpinner, _ := pterm.DefaultSpinner.WithSequence("|", "/", "-", "\\ ").Start("Committing...")
			var commitOutput bytes.Buffer
			commitCmd := exec.Command("git", "commit", "-m", message)
			commitCmd.Stdout = &commitOutput
			commitCmd.Stderr = &commitOutput
			if err := commitCmd.Run(); err != nil {
				commitSpinner.Fail(fmt.Sprintf("Error while committing: %s", commitOutput.String()))
				os.Exit(1)
			}
			commitSpinner.Success("Commit successful!")

			pushSpinner, err := pterm.DefaultSpinner.WithSequence("|", "/", "-", "\\").Start(fmt.Sprintf("pushing files into %s", branch))
			if err != nil {
				log.Fatalf("error to initialize push spinner: %v", err.Error())
			}

			var pushOutput bytes.Buffer
			pushCmd := exec.Command("git", "push", "origin", branch)
			pushCmd.Stdout = &pushOutput
			pushCmd.Stderr = &pushOutput
			if err := pushCmd.Run(); err != nil {
				pushSpinner.Fail(fmt.Sprintf("Error while pushing: %s", pushOutput.String()))
				os.Exit(1)
			}
			pterm.Println()
			pterm.Success.Println("All done!")
		},
	}

	cmd.Flags().StringVarP(&message, "message", "m", "", "The commit message")
	cmd.Flags().StringVarP(&branch, "branch", "b", "main", "The branch to push to")

	rootCmd.AddCommand(cmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
