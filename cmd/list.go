package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <arg>",
	Short: "List samples for specified range",
	Long: `Allows you to list samples from MalwareBazaar (https://bazaar.abuse.ch/)

There are two options supported. You can specify to download the last 100 uploads using the argument "100"
or all the samples in the last hour using the argument "time"`,
	Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{"time", "100"},
	Run: func(cmd *cobra.Command, args []string) {
		// setup for http request
		url := "https://mb-api.abuse.ch/api/v1/"
		method := "POST"
		payload := strings.NewReader("query=get_recent&selector="+strings.Join(args,""))
		
		// store request in listSamples
		var listSamples Samples = getJson(url, method, payload)

		// check if string array is empty
		if len(listSamples.Samples) == 0 {
			fmt.Println("No samples available within the specified time")
			return
		}

		// print output
		for i := 0; i < len(listSamples.Samples); i++ {
		fmt.Println("Sample Information\nSHA256: " + listSamples.Samples[i].SHA256)
			if listSamples.Samples[i].Signature != "" {
				fmt.Println("Signature: " + listSamples.Samples[i].Signature)
			}
			if strings.Join(listSamples.Samples[i].Tags,", ") != "" {
				fmt.Println("Tags: " + strings.Join(listSamples.Samples[i].Tags,", "))
			}
		fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.SetUsageTemplate("Usage:\n  malbaz list [arg]\n\nArguments:\n  100 - lists last 100 samples\n  time - lists samples uploaded in last hour\n")
}