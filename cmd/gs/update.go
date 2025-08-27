package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/pterm/pterm"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

func UpdateCli(version string) *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update gitscribe to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			v := semver.MustParse(version)
			latest, found := CompareVersion(v)
			if !found || latest.Version.LTE(v) {
				log.Println("Current version is the latest")
				return
			}

			fmt.Print("Do you want to update to", latest.Version, "? (y/n): ")
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal("Error to get input:", err.Error())
			}

			if err != nil || (input != "y\n" && input != "n\n") {
				log.Println("Invalid input")
				return
			}

			if input == "n\n" {
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
			return
		},
	}
	return updateCmd
}

func CompareVersion(v semver.Version) (latest *selfupdate.Release, found bool) {
	latest, found, err := selfupdate.DetectLatest("owner/repo")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	pterm.DefaultBox.WithTitle("Commit Suggestion").Printf("version v%v is available to download, exec gs update if you want!")
	return latest, found
}
