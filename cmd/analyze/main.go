package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/http2"
)

type TotalTestResult struct { //go data type
	domainsTested int
	http2enabled  int //increment when http2
	http11enabled int
	http10enabled int
}

type ProbeResult struct {
	http2enabled  bool
	http11enabled bool
	http10enabled bool
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
	request.Close = true
	return client.Do(request)
}

func sendHTTP2Request(domain string) (*http.Response, error) {
	client := &http.Client{Transport: &http2.Transport{}, Timeout: 5 * time.Second}

	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return client.Do(request)
}

func worker(input chan string, output chan ProbeResult) {
	for x := range input {
		output <- filepathHTTP(x)
		fmt.Println(x)
	}
}

/*
1. scan each line (instantiate jobs,results) add to jobs
2. send jobs thru main
*/
func fileEntry(filepath string) TotalTestResult {

	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)

	jobs := make(chan string, 100)
	results := make(chan ProbeResult, 1000000)

	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}

	for x := 0; x < 30; x++ {
		go func() {
			worker(jobs, results)
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
	totalresults := TotalTestResult{}
	for result := range results {
		resultCount += 1
		totalresults.domainsTested += 1

		if result.http2enabled {
			totalresults.http2enabled += 1
		}

		if result.http11enabled {
			totalresults.http11enabled += 1
		}

		if count == resultCount {
			close(results)
		}
	}

	return totalresults
}

func filepathHTTP(myURL string) ProbeResult {
	result := ProbeResult{}
	response, err := sendHTTP2Request(myURL)
	if response != nil {
		response.Body.Close()
	}

	if err == nil { //if http2 request returns error
		result.http2enabled = true
	} else {
		errOtherThanHTTP2Support := !strings.Contains(err.Error(), "unexpected ALPN protocol")
		if errOtherThanHTTP2Support {
			log.Println(err.Error())
		}
	}

	response1, err1 := sendHTTP1Request(myURL)
	if response1 != nil {
		response1.Body.Close()
	}
	if err1 == nil {
		result.http11enabled = true
	} else {
		log.Println(err1.Error())
	}

	return result
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
	// Try to minimize filesystem usage and avoid lingering connections.
	http.DefaultTransport.(*http.Transport).DisableKeepAlives = true

	// fmt.Println(util.Http10Request("https://www.google.com")) Google does
	// fmt.Println(util.Http10Request("https://www.facebook.com")) Facebook does not

	/*
	* use support string for one-off checks, but exclude for mass scans (indicated by use of `-f` for a file of domains)
	 */

	/**
	$ analyze www.twitter.com # base case, probe one domain
	$ analyze domains.txt -f  # -f makes the input a file name instead of URL
	$ analyze domains.txt -f -o results.csv # same as above but write results to results.csv
	*/

	var filepath string
	var urlInput string
	var filepathexport string
	urlInput = os.Args[1]
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepathexport, "o", "", "export to csv")
	flag.Parse()

	if filepathexport != "" && filepath != "" {
		fmt.Println("in both right now")

		totalresults := fileEntry(filepath)
		data := [][]string{

			{"time stamp", "Domain tested", "percent http2", "percent http1.1"},
			{time.Now().String(), fmt.Sprintf("%d", totalresults.domainsTested),
				fmt.Sprintf("%f", (float32(totalresults.http2enabled)/float32(totalresults.domainsTested))*100),
				fmt.Sprintf("%f", (float32(totalresults.http11enabled)/float32(totalresults.domainsTested))*100)},
		}
		/*
			base csv file
			increment when new data incoming
			* If file already exist, only increment/change data not title
		*/
		file, err := os.Create(filepathexport)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, value := range data {
			writer.Write(value)
		}

	} else if filepath != "" {

		totalresults := fileEntry(filepath)
		fmt.Printf("domains tested: %d\n", totalresults.domainsTested)
		fmt.Printf("percent http/2: %.2f%%\n", (float32(totalresults.http2enabled)/float32(totalresults.domainsTested))*100)
		fmt.Printf("percent http/1.1: %.2f%%\n", (float32(totalresults.http11enabled)/float32(totalresults.domainsTested))*100)

	} else if urlInput != "" {
		fmt.Println("in one right now")
		websitepathHTTP2(urlInput)
	}
	/*
		response.Body.Close()
		-f : read file domains
		-o : write to file csv

	*/

}
