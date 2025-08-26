package internal

import (
	"github.com/pterm/pterm"
)

func ConfirmAction(msg string) bool {
	pterm.DefaultBox.WithTitle("Sugestão de Commit").Println(msg)
	pterm.Println()

	confirmed, _ := pterm.DefaultInteractiveConfirm.
		WithConfirmText("Deseja usar esta mensagem e criar o commit?").
		Show()

	return confirmed
}

