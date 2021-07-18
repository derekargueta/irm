package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/derekargueta/irm/pkg/irm"
)

type TotalTestResult struct { //go data type
	domainsTested     int
	http2enabled      int //increment when http2
	http11enabled     int
	http10enabled     int
	errorhttp1occured int
	errorhttp2occured int
	erroroccured      int
}

type ProbeResult struct {
	http2enabled      bool
	http11enabled     bool
	http10enabled     bool
	errorhttp1occured bool
	errorhttp2occured bool
}

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString   = "✅ %s supports HTTP/2\n"

	http1xSupportMsgString = "✅ %s supports HTTP/1.1\n"
	http10SupportMsgString = "✅ %s supports HTTP/1.0\n"
)

func worker(input chan string, output chan ProbeResult) {
	for x := range input {
		output <- filepathHTTP(x)
	}
}

/*
create listener to prevent tcp error
instantiate before starting workers
*/
func fileEntry(filepath string, workers int) TotalTestResult {

	domains, erroropen := os.Open(filepath)
	domain := bufio.NewScanner(domains)

	jobs := make(chan string, 300)
	results := make(chan ProbeResult, 1000000)

	if erroropen != nil {
		log.Fatal(erroropen)
		os.Exit(1)
	}

	for x := 0; x < workers; x++ {
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

		if result.errorhttp1occured && result.errorhttp2occured {
			totalresults.erroroccured += 1
		}

		if result.errorhttp1occured {
			totalresults.errorhttp1occured += 1
		}
		if result.errorhttp2occured {
			totalresults.errorhttp2occured += 1
		}

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
	response, err := irm.SendHTTP2Request(myURL)

	if response != nil {
		response.Body.Close()
	}

	if err == nil { //if http2 request returns error
		result.http2enabled = true
	} else {
		errOtherThanHTTP2Support := !strings.Contains(err.Error(), "unexpected ALPN protocol")
		if errOtherThanHTTP2Support {
			log.Println(err.Error() + " - request error for http2")
			result.errorhttp2occured = true
		}

	}

	response1, err1 := irm.SendHTTP1Request(myURL)
	if response1 != nil {
		response1.Body.Close()
	}
	if err1 == nil {
		result.http11enabled = true
	} else {
		log.Println(err1.Error() + " by http1.1 request")
		result.errorhttp1occured = true
	}

	return result
}

func websitepathHTTP2(urlInput string) {
	time.Sleep(100 * time.Millisecond)
	response, err := irm.SendHTTP2Request(urlInput)

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

	var filepath string
	// var urlInput string
	var filepathexport string
	var numWorkers int
	var timebetrun int
	// urlInput = os.Args[1]
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepathexport, "o", "", "export to csv")
	flag.IntVar(&numWorkers, "w", runtime.NumCPU()*2, "number of workers")
	flag.IntVar(&timebetrun, "d", 10, "time between runs")
	flag.Parse()

	log.Printf("Running with %d goroutine workers\n", numWorkers)

	maindr, merr := os.Getwd()
	if merr != nil {
		log.Fatal(merr)
	}
	fullresultdr := strings.Replace(maindr, "\\", "/", -1) + "/cmd/analyze/results/results.csv"
	checkFile, err := os.Stat(fullresultdr)

	if filepath != "" {
		for x := 0; x > -1; x++ {
			fmt.Println("in both right now")

			totalresults := fileEntry(filepath, numWorkers)
			dataHead := [][]string{
				{"time stamp", "Domain tested", "percent http2", "percent http1.1", "percent connection error: "},
			}

			data := [][]string{
				{time.Now().Format("2006-01-02 15:04:05"),
					fmt.Sprintf("%d", totalresults.domainsTested),
					fmt.Sprintf("%.2f%%", (float32(totalresults.http2enabled)/float32(totalresults.domainsTested))*100),
					fmt.Sprintf("%.2f%%", (float32(totalresults.http11enabled)/float32(totalresults.domainsTested))*100),
					fmt.Sprintf("%.2f%%", (float32(totalresults.erroroccured)/float32(totalresults.domainsTested))*100)},
			}

			if os.IsNotExist(err) {
				log.Println(checkFile)
				log.Println("it doesnt exist, making one NOW")
				file, err := os.Create(fullresultdr)
				if err != nil {
					fmt.Println(err.Error())
				}
				writer := csv.NewWriter(file)

				for _, x := range dataHead {
					writer.Write(x)
				}
				writer.Flush()
				file.Close()

			}

			file, err := os.OpenFile(fullresultdr, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				log.Fatal(err.Error())
			}

			writer := csv.NewWriter(file)

			for _, value := range data {
				writer.Write(value)
			}
			writer.Flush()
			file.Close()
			fmt.Println("Done")

			cmd := exec.Command("bash cron.sh")
			cmd.Run()
			time.Sleep(time.Duration(timebetrun) * time.Second)
		}
	}

}
