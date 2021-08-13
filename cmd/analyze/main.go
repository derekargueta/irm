package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"

	// "net/http"
	"os"
	"runtime"
	"time"

	"github.com/derekargueta/irm/pkg/irm"
	"github.com/derekargueta/irm/pkg/irm/probes"
	"github.com/derekargueta/irm/pkg/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
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
		log.Println(erroropen)
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

	http2Result := (&probes.HTTP2Probe{}).Run(myURL)
	result.errorhttp2occured = http2Result.Err != nil
	result.http2enabled = http2Result.Supported

	http1Result := (&probes.HTTP11Probe{}).Run(myURL)
	result.errorhttp1occured = http1Result.Err != nil
	result.http11enabled = http1Result.Supported

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
	var filepathexport string
	var numWorkers int
	var timebetrun int
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepathexport, "o", "", "export to csv")
	flag.IntVar(&numWorkers, "w", runtime.NumCPU()*2, "number of workers")
	flag.IntVar(&timebetrun, "d", 10, "time between runs")
	flag.Parse()
	//mainPath := "app/"
	//"/root/"
	// "/Users/Tavo"
	log.Printf("Running with %d goroutine workers\n", numWorkers)

	for {
		if filepath != "" {
			fmt.Println("in both right now")
			timer := time.Now()
			totalresults := fileEntry(filepath, numWorkers)
			domainsTested := totalresults.domainsTested
			data := [][]string{
				{time.Now().Format("2006-01-02 15:04:05"),
					fmt.Sprintf("%d", totalresults.domainsTested),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.http2enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.http11enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.erroroccured, domainsTested)),
					fmt.Sprintf("%.2fs", time.Since(timer).Seconds()),
				}}

			//          TOKEN AUTHENTICATION

			/*
				_, plainerr := git.PlainClone("./tempirmdata/irm-data", false, &git.CloneOptions{
					Auth: &http.BasicAuth{
						Username: "123",
						Password: password,
					},
					URL:      "https://github.com/derekargueta/irm-data",
					Progress: os.Stdout,
				})
			*/

			//	         SHH Authentication
			os.Setenv("SSH_KNOWN_HOSTS", "/home/.ssh/known_hosts")
			publicKeys, err := ssh.NewPublicKeysFromFile("git", "/home/.ssh/id_ed25519", "") //
			if err != nil {
				log.Printf("generate publickeys failed: %s\n", err.Error())
			}
			checkFile, err := os.Open("/app/tempirmdata/results.csv")
			if err != nil {
				_, plainerr := git.PlainClone("/app/tempirmdata", false, &git.CloneOptions{
					Auth:     publicKeys,
					URL:      "git@github.com:derekargueta/irm-data.git",
					Progress: os.Stdout,
				})

				log.Println("in process of cloning")

				if plainerr != nil {
					log.Printf("cant clone : %s", plainerr)
				}
			}
			checkFile.Close()
			// }

			file, err := os.OpenFile("/app/tempirmdata/results.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
			if err != nil {
				log.Println(err.Error() + "cant open results in tempirmdata")
			}

			writer := csv.NewWriter(file)

			for _, value := range data {
				writer.Write(value)
				log.Println(value)
			}
			writer.Flush()
			log.Println(writer.Error())
			file.Close()
			fmt.Println("Done")

			//       patck this value over time.h 4
			repo, mrr := git.PlainOpen("/app/tempirmdata") // checkFile.Close()
			if mrr != nil {
				log.Println("cant open dir")
			}
			tree, mmrr := repo.Worktree()
			fmt.Println(tree.Status())
			if mmrr != nil {
				log.Println(err)
			}

			_, err = tree.Add("results.csv")
			if err != nil {
				log.Println("doesn't exists")
			} else {
				log.Println("exists")
			}
			_, err = tree.Commit(time.Now().Format("2006-01-02 15:04:05"), &git.CommitOptions{All: true,
				Author: &object.Signature{
					Name:  "H",
					Email: "t",
					When:  time.Now(),
				},
			})
			if err != nil {
				log.Println("commit not workig properly")
			}
			mrr = repo.Push(&git.PushOptions{
				RemoteName: "origin",
				Auth:       publicKeys,
			})
			log.Printf("errors that happened: %s", mrr)

		}
		time.Sleep(time.Duration(timebetrun) * time.Second)
	}

}
