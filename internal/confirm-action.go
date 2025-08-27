package internal

import (
	"github.com/pterm/pterm"
)

// ConfirmAction shows the user the generated commit message and asks for confirmation.
func ConfirmAction(msg string) bool {
	pterm.DefaultBox.WithTitle("Commit Suggestion").Println(msg)
	pterm.Println()

	confirmed, _ := pterm.DefaultInteractiveConfirm.
		Show()

	return confirmed
}

