package internal

import (
	"github.com/pterm/pterm"
)

func ConfirmAction(msg string) bool {
	pterm.DefaultBox.WithTitle("Sugestão de Commit").Println(msg)
	pterm.Println()

	confirmed, _ := pterm.DefaultInteractiveConfirm.
		Show()

	return confirmed
}

