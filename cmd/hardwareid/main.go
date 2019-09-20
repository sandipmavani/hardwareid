// Package main provides the command line app for reading the unique hardware id of most OSs.
//
// Usage: hardwareid [options]
//
// Options:
//   --appid    <AppID>    Protect hardware id by hashing it together with an app id.
//
// Try:
//   hardwareid
//   hardwareid --appid MyAppID
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sandipmavani/hardwareid"
)

const usageStr = `
Usage: hardwareid [options]

Options:
  --appid    <AppID>    Protect hardware id by hashing it together with an app id.

Try:
  hardwareid
  hardwareid --appid MyAppID
`

func usage() {
	log.Fatalln(usageStr)
}

func main() {
	var appID string
	flag.StringVar(&appID, "appid", "", "Protect hardware id by hashing it together with an app id.")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	var id string
	var err error
	if appID != "" {
		id, err = hardwareid.ProtectedID(appID)
	} else {
		id, err = hardwareid.ID()
	}
	if err != nil {
		log.Fatalf("Failed to read hardware id with error: %s\n", err)
	}
	fmt.Println(id)
}
