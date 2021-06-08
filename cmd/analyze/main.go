package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString = "✅ %s supports HTTP/2\n"
)

func checkArgLength() {
	// Exclude the first "arg" used for the program name, e.g. `./analyze`.
	numArgs := len(os.Args) - 1

	if numArgs < 1 {
		log.Fatalf(errNumArgsMsgString, numArgs)
	}
}

func sendHTTP2Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: &http2.Transport{}}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://" + domain, nil)
	return client.Do(request)
}

func main() {
	checkArgLength()
	domain := os.Args[1]

	response, err := sendHTTP2Request(domain)
	// If an error occurred, then the website probably does not support HTTP/2. This error-checking
	// should be more advanced to account for things like network errors. It currently just assumes
	// _any_ error implies HTTP/2 incompatability which is not necessarily true.
	if err != nil {
		fmt.Printf(http2NoSupportMsgString, domain)
	} else {
		fmt.Printf(http2SupportMsgString, domain)
	}

	if response != nil {
		response.Body.Close()
	}
}
