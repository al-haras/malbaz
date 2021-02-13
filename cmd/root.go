package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "malbaz",
	Short: "Simple CLI tool made to get information or download the latest samples from MalwareBazaar",
	Long: `Simple CLI tool made to get information or download the latest samples from MalwareBazaar

This main purpose of this tool is to get information or download the latest samples from MalwareBazaar (https://bazaar.abuse.ch/) for testing purposes. 
Supported commands allow you to list or download samples within the last hour or the last 100 samples uploaded to the platform.
Additionally, you have the ability to grab the collection of the yesterday's sample collection.`,
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
