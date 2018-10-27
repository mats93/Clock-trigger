/*
  File: main.go
  Contains the clock trigger program.
  Assignment 2: IGC track viewer extended - IMT2681-2018 (Cloud Technologies)

  By Mats Ove Mandt Skjærstein.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// APIURL is the url for the paraglider API.
const APIURL = "https://paragliding-api.herokuapp.com/paragliding/api/"

// Holds the track information.
type trackInfo struct {
	TimeLatest int64         `json:"t_latest"`
	TimeStart  int64         `json:"t_start"`
	TimeStop   int64         `json:"t_stop"`
	Tracks     []int         `json:"tracks"`
	Processing time.Duration `json:"processing"`
}

// Stored the latest timestamp the clock trigger knows about.
var latestTimestamp string

// The webhook url to use.
var webhookURL string

// Gets the latest timestamp from the paragliding API.
func getLatestTimestamp() {
	// Formats the url.
	url := string(APIURL + "ticker/latest")

	// Creates a GET request.
	resp, _ := http.Get(url)

	// Grabs the contents from the GET request.
	output, _ := ioutil.ReadAll(resp.Body)

	// Closes the response body.
	defer resp.Body.Close()

	// Check if the output is empty.
	if string(output) != "" {
		// Stores the timestamp.
		latestTimestamp = string(output)
	}
}

// Gets the timestamp information newer then the timestamp provided.
func getTrackInfoNewerThen(timestamp string) {
	// Formats the url.
	url := string(APIURL + "ticker/" + timestamp)

	// Creates a GET request.
	resp, _ := http.Get(url)

	// Grabs the contents from the GET request.
	output, _ := ioutil.ReadAll(resp.Body)

	// Closes the response body.
	defer resp.Body.Close()

	// Check if there was any new timestamps.
	if string(output) != "" {
		var info trackInfo

		// Converts json to struct.
		json.Unmarshal(output, &info)

		// Notify the webhook about new tracks.
		notifyWebhook(info)
	}
}

// Notifies the webhook about the new tracks.
func notifyWebhook(info trackInfo) {

	// Formats the message:
	newTracks := fmt.Sprintf("New tracks added since last check: %v\n", info.Tracks)
	latest := fmt.Sprintf("Latest added timestamp: %v\n", info.TimeLatest)
	first := fmt.Sprintf("First timestamp of the added tracks: %v\n", info.TimeStart)
	last := fmt.Sprintf("Last timestamp of the added tracks: %v\n", info.TimeStop)
	message := newTracks + latest + first + last

	// Converts the message to json.
	jsonMessage := map[string]string{"text": message}
	body, _ := json.Marshal(jsonMessage)

	// Send the message.
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err == nil {
		defer resp.Body.Close()
	}

}

func main() {
	// Initialise the webhook.
	webhookURL = "https://hooks.slack.com/services/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"

	// Gets the latest timestamp from the API.
	getLatestTimestamp()

	// Loops forever.
	for {
		// Check if any new tracks have been added since last check.
		// If new are added, notify the webhook.
		getTrackInfoNewerThen(latestTimestamp)

		// Gets the latest added timestamp.
		getLatestTimestamp()

		// Wait for 10 min.
		time.Sleep(10 * time.Minute)
	}
}
