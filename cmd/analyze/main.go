package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/derekargueta/irm/pkg/util"
	"golang.org/x/net/http2"
)

type results struct { //go data type
	domainsTested int
	http2enabled  int //increment when http2
	http11enabled int
	http10enabled int
}

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString   = "✅ %s supports HTTP/2\n"

	http1xSupportMsgString = "✅ %s supports HTTP/1.1\n"
	http10SupportMsgString = "✅ %s supports HTTP/1.0\n"
)

func sendHTTP1Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	return client.Do(request)
}

func sendHTTP2Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: &http2.Transport{}, Timeout: 5 * time.Second}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	return client.Do(request)
}

func worker(input, output chan string, totalresults results) {
	for x := range input {
		output <- filepathHTTP(x, totalresults)
		fmt.Println(x)
	}
}

/*
1. scan each line (instantiate jobs,results) add to jobs
2. send jobs thru main
*/
func fileEntry(filepath string, totalresults results) {

	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)

	jobs := make(chan string, 100)
	results := make(chan string, 1000)

	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}
	// mycount := 0
	// for domain.Scan() {
	// 	mycount++
	// }
	// log.Println()
	for x := 0; x < 30; x++ {
		go func() {
			worker(jobs, results, totalresults)
		}()
	}
	log.Println("workers started")
	count := 0
	for domain.Scan() {
		jobs <- domain.Text()
		count++
	}
	close(jobs)

	resultCount := 0
	for result := range results {
		fmt.Println(result)
		resultCount += 1

		if count == resultCount {
			close(results)
		}
	}
}

func filepathHTTP(myURL string, totalresults results) string {
	fmt.Println(myURL)
	response, err := sendHTTP2Request(myURL)

	if response != nil {
		response.Body.Close()
	}
	if err != nil { //if http2 request returns error
		response1, err1 := sendHTTP1Request(myURL)

		if response1 != nil {
			response1.Body.Close()
		}
		if err1 != nil {
			http10test := util.Http10Request("https://www." + myURL)
			if http10test == true {
				totalresults.http10enabled++
				return fmt.Sprintf(http10SupportMsgString, myURL)
			} else {
				log.Println(err1.Error())
				fmt.Println("")
			}

		} else {
			totalresults.http11enabled++
			return fmt.Sprintf(http1xSupportMsgString, myURL)

		}
	}

	totalresults.http2enabled++
	return fmt.Sprintf(http2SupportMsgString, myURL)

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

	// fmt.Println(util.Http10Request("https://www.google.com")) Google does
	// fmt.Println(util.Http10Request("https://www.facebook.com")) Facebook does not

	/**
	$ analyze www.twitter.com # base case, probe one domain
	$ analyze domains.txt -f  # -f makes the input a file name instead of URL
	$ analyze domains.txt -f -o results.csv # same as above but write results to results.csv
	*/

	var filepath string
	var urlInput string
	var totalresults results
	urlInput = os.Args[1]
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepath, "f -o", "", "export to csv")
	flag.Parse()

	if filepath != "" {
		fileEntry(filepath, totalresults)
		//filepathHTTP2OLD(filepath)

	} else if urlInput != "" {
		websitepathHTTP2(urlInput)
	}
	fmt.Printf("http2: %d", totalresults.http2enabled)
	/*
		response.Body.Close()
		-f : read file domains
		-o : write to file csv

	*/

}
