package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "github.com/derekargueta/irm/pkg/util"

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

/*
1. scan each line (instantiate jobs,results) add to jobs
2. send jobs thru main
*/
func fileEntry(filepath string) {

	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)
	count := 0

	chunkwebs := []string{}

	jobs := make(chan []string, 100)
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

	newcount := 0
	for domain.Scan() {
		if newcount < 1000 {
			chunkwebs = append(chunkwebs, domain.Text())
			if count == 10 {
				log.Println(chunkwebs)
				jobs <- chunkwebs
				chunkwebs = []string{}
				count = 0
			}
			count++
		}
		newcount++

	}

	for x := 0; x < 10; x++ {
		go Worker(jobs, results)
		log.Println("workers started")
	}

	for a := 1; a <= 1000; a++ {
		<-results
	}
}

func Worker(jobs <-chan []string, results chan<- string) {
	for x := range jobs {
		for y := range x {
			log.Println(filepathHTTP(x[y]))
			results <- filepathHTTP(x[y])
		}

	}

}

func filepathHTTP(myURL string) string {

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
			fmt.Println("broken")
			fmt.Println("")
		} else {
			return fmt.Sprintf(http1xSupportMsgString, myURL)

		}
	}
	return fmt.Sprintf(http2SupportMsgString, myURL)

	// results.domainsTested = countoverall
	// results.http2enabled = counthttp2

}

func filepathHTTP2OLD(filepath string) {
	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)
	limit := 0
	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}
	for domain.Scan() {
		limit++
		if limit <= 10 {
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
					fmt.Println("")
				} else {
					fmt.Println(http1xSupportMsgString, domain.Text())

				}
			} else {
				fmt.Println(http2SupportMsgString, domain.Text())

			}
		} else {
			break
		}

	}

	// results.domainsTested = countoverall
	// results.http2enabled = counthttp2

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

	// fmt.Println(util.Http10Request("https://www.google.com")) Google does
	// fmt.Println(util.Http10Request("https://www.facebook.com")) Facebook does not

	/**
	$ analyze www.twitter.com # base case, probe one domain
	$ analyze domains.txt -f  # -f makes the input a file name instead of URL
	$ analyze domains.txt -f -o results.csv # same as above but write results to results.csv
	*/

	var filepath string
	var urlInput string

	urlInput = os.Args[1]
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepath, "f -o", "", "export to csv")
	flag.Parse()

	if filepath != "" {
		fileEntry(filepath)
		//filepathHTTP2OLD(filepath)

	} else if urlInput != "" {
		websitepathHTTP2(urlInput)
	}

	/*
		response.Body.Close()
		-f : read file domains
		-o : write to file csv

	*/

}
