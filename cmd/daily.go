package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Downloads daily sample batches",
	Long: `Downloads the prior day's malware batches into a .zip file.

The collection or daily batches can be found here: https://mb-api.abuse.ch/downloads/`,
	Args: cobra.NoArgs,
	//ValidArgs: []string{"time", "100"},
	Run: func(cmd *cobra.Command, args []string) {
		yesterday := time.Now().AddDate(0, 0, -1)
		fileName := yesterday.Format("2006-01-02") + ".zip"
		fileUrl := "https://mb-api.abuse.ch/downloads/" + fileName
		err := DownloadFile(fileName, fileUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + fileUrl)
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
	dailyCmd.SetUsageTemplate("Usage:\n  malbaz daily\n\nDoes not take arguments. This command simply downloads yesterday's collection of samples.\n")
}
