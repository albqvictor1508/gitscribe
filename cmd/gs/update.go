package main

import (
	"fmt"
	"os"

	"github.com/albqvictor1508/gitscribe/internal"
	"github.com/pterm/pterm"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

func UpdateCli() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update gitscribe to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			pterm.Info.Println("Checking for updates...")
			latest, found, err := selfupdate.DetectLatest("albqvictor/gitscribe")
			if err != nil {
				pterm.Error.Printf("Error occurred while detecting version: %v\n", err)
				return
			}

			v, err := selfupdate.NewVersion(version)
			if err != nil {
				pterm.Error.Printf("Error occurred while parsing current version: %v\n", err)
				return
			}

			if !found || latest.Version.LTE(v) {
				pterm.Success.Printf("Current version %s is the latest.\n", version)
				return
			}

			pterm.Warning.Printf("New version available: %s\n", latest.Version)
			if !internal.ConfirmAction(fmt.Sprintf("Do you want to update to version %s?", latest.Version)) {
				return
			}

			exe, err := os.Executable()
			if err != nil {
				pterm.Error.Println("Could not locate executable path.")
				return
			}

			updateSpinner, _ := pterm.DefaultSpinner.Start("Updating...")
			if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
				updateSpinner.Fail(fmt.Sprintf("Error occurred while updating: %v", err))
				return
			}

			updateSpinner.Success(fmt.Sprintf("Successfully updated to version %s.", latest.Version))
		},
	}

	return updateCmd
}
