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
	http11enabled int
}

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString   = "✅ %s supports HTTP/2\n"

	http1xSupportMsgString = "✅ %s supports HTTP/1.1\n"
)

func sendHTTP1Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 10 * time.Second}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	return client.Do(request)
}

func sendHTTP2Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	return client.Do(request)
}

func filepathHTTP2(filepath string, results results) {
	counthttp2 := 0
	countoverall := 0
	limit := 0
	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)

	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}

	for domain.Scan() {
		if limit < 10 {
			countoverall++

			response, err := sendHTTP2Request(domain.Text())
			if response != nil {
				response.Body.Close()
			}

			if err != nil { //if http2 request returns error

				response1, err1 := sendHTTP1Request(domain.Text())
				if response1 != nil {
					response1.Body.Close()
				}
				if err1 != nil {
					fmt.Println("broken")
				} else {
					fmt.Printf(http1xSupportMsgString, domain.Text())
					fmt.Println("")
				}
			} else {
				fmt.Printf(http2SupportMsgString, domain.Text())
				fmt.Println("")
				counthttp2++
			}

		} else {
			break
		}
		limit++
	}

	results.domainsTested = countoverall
	results.http2enabled = counthttp2

}

/*

 */
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
	var results results

	var filepath string
	var urlInput string

	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&urlInput, "o", "", "Url")
	flag.Parse()

	if filepath != "" {
		filepathHTTP2(filepath, results)
	} else if urlInput != "" {
		websitepathHTTP2(urlInput)
	}

	/*
		response.Body.Close()
		-f : read file domains
		-o : write to file csv

	*/

}
