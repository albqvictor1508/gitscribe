package internal

import (
	"github.com/pterm/pterm"
)

/*
* Copyright (c) 2025 Victor Albuquerque Arruda. All Rights Reserved.
* */

func ConfirmAction(msg string) bool {
	pterm.DefaultBox.WithTitle("Commit Suggestion").Println(msg)
	pterm.Println()

	confirmed, _ := pterm.DefaultInteractiveConfirm.
		Show()

	return confirmed
}
