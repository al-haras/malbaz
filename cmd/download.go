package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download samples for a specified range",
	Long: `Allows you to download samples from MalwareBazaar (https://bazaar.abuse.ch/)

There are two options supported. You can specify to download the last 100 uploads using the argument "100"
or all the samples in the last hour using the argument "time"`,
	Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{"time", "100"},
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://mb-api.abuse.ch/api/v1/"
		method := "POST"
		payload := strings.NewReader("query=get_recent&selector="+strings.Join(args,""))
		
		var listSamples Samples = getJson(url, method, payload)

		// check if string array is empty
		if len(listSamples.Samples) == 0 {
			fmt.Println("No samples available within the specified time")
			return
		}

		// print output
		for i := 0; i < len(listSamples.Samples); i++ {
			fileBody := strings.NewReader("query=get_file&sha256_hash=" + listSamples.Samples[i].SHA256)
			downloadSample(method, url, listSamples.Samples[i].SHA256, fileBody)
			fmt.Println("Downloaded Sample: " + listSamples.Samples[i].SHA256 + ".zip\n")
			}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.SetUsageTemplate("Usage:\n  malbaz download [arg]\n\nArguments:\n  100 - downloads last 100 samples uploaded\n  time - downloads samples uploaded in last hour\n")
}
