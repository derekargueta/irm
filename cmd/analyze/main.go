package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

type results struct { //go data type
	domainsTested int
	http2enabled  int //increment when http2
	http1enabled  int
	http11enabled int
}

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString   = "✅ %s supports HTTP/2\n"
)

func sendHTTP1Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: http.DefaultTransport}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "http://"+domain, nil)
	return client.Do(request)
}

func sendHTTP2Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: &http2.Transport{}}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://"+domain, nil)
	return client.Do(request)
}

func filepathHTTP2(filepath string) {
	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)

	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}

	for domain.Scan() {
		time.Sleep(100 * time.Millisecond)
		response, err := sendHTTP2Request(domain.Text())

		if response != nil {
			response.Body.Close()
		}
		if err != nil {
			fmt.Printf(http2NoSupportMsgString, domain.Text())
		} else {
			fmt.Printf(http2SupportMsgString, domain.Text())
		}

	}
}

func websitepathHTTP2(urlInput string) {
	time.Sleep(100 * time.Millisecond)
	response, err := sendHTTP2Request(urlInput)

	if response != nil {
		response.Body.Close()
	}
	if err != nil {
		fmt.Printf(http2NoSupportMsgString, urlInput)
	} else {
		fmt.Printf(http2SupportMsgString, urlInput)
	}
}

func main() {
	var filepath string
	var urlInput string

	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&urlInput, "o", "", "Url")
	flag.Parse()

	if filepath != "" {
		filepathHTTP2(filepath)
	} else if urlInput != "" {
		websitepathHTTP2(urlInput)
	}

	/*
		-f : read file domains
		-o : write to file csv

	*/

}
