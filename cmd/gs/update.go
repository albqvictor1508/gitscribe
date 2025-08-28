package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/pterm/pterm"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

var (
	latest *selfupdate.Release
	v      semver.Version
)

func UpdateCli(version string) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update gitscribe to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			v = semver.MustParse(version)
			err := CheckForUpdate()
			if err != nil && latest != nil {
				log.Fatal(err)
			}

			if latest == nil {
				pterm.Info.Println("Current version is the latest")
				return
			}

			fmt.Println("Do you want to update to ", latest.Version, "?")
			pterm.Println()
			confirmed, _ := pterm.DefaultInteractiveConfirm.
				Show()

			if !confirmed {
				log.Println("Update canceled")
				return
			}

			exe, err := os.Executable()
			if err != nil {
				log.Println("Could not locate executable path")
				return
			}
			if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
				log.Println("Error occurred while updating binary:", err)
				return
			}
			log.Println("Successfully updated to version", latest.Version)
		},
	}
	return updateCmd
}

func CheckForUpdate() (err error) {
	l, found, err := selfupdate.DetectLatest("albqvictor1508/gitscribe")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		os.Exit(1)
		return err
	}

	if !found || latest.Version.LTE(v) {
		pterm.Info.Println("Current version is latest")
		return nil
	}
	latest = l

	return nil
}

func ShowUpdate(v semver.Version) {
	pterm.DefaultBox.WithTitle("Update Available").Println(fmt.Sprintf("The version v%s is new e parara", v))
	pterm.Println()
}
