package cmd

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"io"
)

// function to get the json from malwarebazaar api for samples via http request
func getJson(infoMethod string, infoEndpoint string, infoBody *strings.Reader) Samples {
	
	// create http client, setup req (request)
	client := &http.Client {
	}
	req, err := http.NewRequest(infoEndpoint, infoMethod, infoBody)
		if err != nil {
		fmt.Println(err)
		}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// response stored as req
	res, err := client.Do(req)
		if err != nil {
		fmt.Println(err)
		}
	defer res.Body.Close()
	byteValue, _ := ioutil.ReadAll(res.Body)
	
	// define samples and returned unmarshalled json
	var samples Samples
	json.Unmarshal(byteValue, &samples)
	return samples
}

// function to download individual samples via required POST request
func downloadSample(downloadMethod string, downloadEndpoint string, downloadPath string, downloadBody *strings.Reader) error {

	// create http client, setup req (request)
	client := &http.Client {
	}
	req, err := http.NewRequest(downloadMethod, downloadEndpoint, downloadBody)
		if err != nil {
		fmt.Println(err)
		}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// response stored as req
	res, err := client.Do(req)
		if err != nil {
		fmt.Println(err)
		}
	defer res.Body.Close()
	
	out, err := os.Create(downloadPath + ".zip")
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	return err
}

// function to download file to local path via http GET request
func DownloadFile(dailyPath string, dailyUrl string) error {

	// get request to dailyUrl
	resp, err := http.Get(dailyUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(dailyPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
