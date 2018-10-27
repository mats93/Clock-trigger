/*
  File: main.go
  Contains the clock trigger program.
  Assignment 2: IGC track viewer extended - IMT2681-2018 (Cloud Technologies)

  By Mats Ove Mandt Skjærstein.
*/

package main

import "time"

// Holds the track information.
type trackInfo struct {
	TimeLatest int64         `json:"t_latest"`
	TimeStart  int64         `json:"t_start"`
	TimeStop   int64         `json:"t_stop"`
	Tracks     []int         `json:"tracks"`
	Processing time.Duration `json:"processing"`
}

// Stores information about the Slack webhook.
type webhook struct {
}

// Stored the latest timestamp the clock trigger knows about.
var latestTimestamp string

// Gets the latest timestamp from the paragliding API.
func getLatestTimestamp() {

}

// Gets the timestamp information newer then the timestamp provided.
func getTrackInfoNewerThen(timestamp string) {

}

// Notifies the webhook about the new tracks.
func notifyWebhook() {

}

func main() {
	// Sjekk om latest timestmap er høyere enn den som er lagret her.
	// Hvis den er høyere, nye tracks er lagt til. (/api/ticker/latest)

	// Sjekk så alle nyere tracks enn det programmet vet om.
	// Bruk lagret timestmap her og kjør mot (/api/ticker/<timestamp>)
	// Denne json informasjonen skal sendes til webhook.
	// Denne nye "siste" timestmapen skal lagres in-memory.

}
