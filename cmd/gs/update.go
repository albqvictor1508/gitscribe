package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
			latest := CompareVersion(v)

			if latest == nil {
				log.Println("No release info found")
				return
			}

			fmt.Print("Do you want to update to", latest.Version, "? (y/n): ")
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal("Error to get input:", err.Error())
				return
			}

			answer := strings.TrimSpace(strings.ToLower(input))
			if answer != "y" {
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

func CompareVersion(v semver.Version) (latest *selfupdate.Release) {
	latest, found, err := selfupdate.DetectLatest("albqvictor1508/gitscribe")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		os.Exit(1)
		return
	}

	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	pterm.DefaultBox.WithTitle("Update Available").Printf("version v%s is available to download, exec gs update if you want!", v)
	return latest
}
