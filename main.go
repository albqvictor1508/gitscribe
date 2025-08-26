package main

import (
	"fmt"
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

	spinner := pterm.DefaultSpinner
	spinner.Sequence = []string{"|", "/", "-", "\\"}

	rootCmd := &cobra.Command{Use: "gitscribe"}

	cmd := &cobra.Command{
		Use:   "cmt [files]",
		Args:  cobra.MinimumNArgs(0),
		Short: "AI-powered git add, commit, and push",
		Run: func(cmd *cobra.Command, args []string) {
			asciiArt := `
'||''|.  '||    ..|''||'  '||''|.''||''|. '||'  '||''|.  '||''|'
 ||   ||  ||  .|'    ||   ||   ||  ||   || ||    ||   ||  ||  ||
 ||    || ||  ||      ||  ||''|'   ||   || ||    ||    ||  ||''|'
 ||    |  ||  '|.     ||  ||   |.  ||   || ||    ||    |   ||
.||...|' .||.  ''|...|'  .||.  '| .||.  .||.  .||...|' .||.'
`
			pterm.DefaultBasicText.Println(pterm.FgCyan.Sprint(asciiArt))
			pterm.Println()
			time.Sleep(time.Second)

			files := args
			if len(files) == 0 {
				files = append(files, ".")
			}
			addSpinner, _ := spinner.Start("Staging files...")
			for _, file := range files {
				r := exec.Command("git", "add", file)
				if _, err := r.Output(); err != nil {
					addSpinner.Fail(fmt.Sprintf("Failed to stage file %s: %v", file, err))
					os.Exit(1)
				}
			}
			addSpinner.Success("Files staged successfully!")

			if len(message) == 0 {
				aiSpinner, _ := spinner.Start("Analyzing changes and generating message with AI...")
				diff := exec.Command("git", "diff", "--staged")
				res, err := diff.CombinedOutput()
				if err != nil {
					aiSpinner.Fail(fmt.Sprintf("Failed to get git diff: %s", string(res)))
					os.Exit(1)
				}

				if len(res) == 0 {
					aiSpinner.Warning("No changes found in stage. Nothing to commit.")
					os.Exit(0)
				}

				context := fmt.Sprintf("Based on the git diff below, create a commit message that follows the 'Conventional Commits' specification. Return only the commit message, with nothing else: %v", string(res))

				msg, err := internal.SendPrompt(context)
				if err != nil {
					aiSpinner.Fail(fmt.Sprintf("Error generating message with AI: %v", err))
					os.Exit(1)
				}
				message = msg
				aiSpinner.Success("Commit message generated!")
			}

			if !internal.ConfirmAction(message) {
				pterm.Warning.Println("Operation cancelled by user.")
				os.Exit(1)
			}

			commitSpinner, _ := spinner.Start("Committing...")
			raw := exec.Command("git", "commit", "-m", message)
			if output, err := raw.CombinedOutput(); err != nil {
				commitSpinner.Fail(fmt.Sprintf("Error while committing: %s", string(output)))
				os.Exit(1)
			}
			commitSpinner.Success("Commit successful!")

			pushSpinner, _ := spinner.Start(fmt.Sprintf("Executing 'git push origin %s'வுகளை...", branch))
			pushCmd := exec.Command("git", "push", "origin", branch)
			if output, err := pushCmd.CombinedOutput(); err != nil {
				pushSpinner.Fail(fmt.Sprintf("Error while pushing: %s", string(output)))
				os.Exit(1)
			}
			pushSpinner.Success("Push successful!")
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
