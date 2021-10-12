package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
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

type ProbeResult struct {
	http10enabled     bool
	http11enabled     bool
	http2enabled      bool
	errorhttp1occured bool
	errorhttp2occured bool
	tls10enabled      bool
	tls11enabled      bool
	tls12enabled      bool
	tls13enabled      bool
	cloudflare        bool
	cloudflareipv4    bool
	cloudflareipv6    bool
	fastlyprobe       bool
	fastlyprobeipv4   bool
	fastlyprobeipv6   bool
}
type TotalTestResult struct {
	domainsTested     int
	http11enabled     int
	http2enabled      int
	http10enabled     int
	errorhttp1occured int
	errorhttp2occured int
	erroroccured      int
	tls10enabled      int
	tls11enabled      int
	tls12enabled      int
	tls13enabled      int
	cloudflare        int
	cloudflareipv4    int
	cloudflareipv6    int
	fastlyprobe       int
	fastlyprobeipv4   int
	fastlyprobeipv6   int
}

type cdn_fast struct {
	Ipv4_addresses []string `json:"addresses"`
	Ipv6_addresses []string `json:"ipv6_addresses"`

	Ipv4_addresses_cidr []*net.IPNet
	Ipv6_addresses_cidr []*net.IPNet
}

func (t *TotalTestResult) AddResult(result ProbeResult) {
	t.domainsTested += 1

	if result.errorhttp1occured && result.errorhttp2occured {
		t.erroroccured += 1
	}

	if result.errorhttp1occured {
		t.errorhttp1occured += 1
	}

	if result.errorhttp2occured {
		t.errorhttp2occured += 1
	}

	if result.http2enabled {
		t.http2enabled += 1
	}

	if result.http11enabled {
		t.http11enabled += 1
	}

	if result.tls10enabled {
		t.tls10enabled += 1
	}
	if result.tls11enabled {
		t.tls11enabled += 1
	}
	if result.tls12enabled {
		t.tls12enabled += 1
	}
	if result.tls13enabled {
		t.tls13enabled += 1
	}
	if result.cloudflare {
		t.cloudflare += 1
	}
	if result.cloudflareipv4 {
		t.cloudflareipv4 += 1
	}
	if result.cloudflareipv6 {
		t.cloudflareipv6 += 1
	}
	if result.fastlyprobe {
		t.fastlyprobe += 1
	}
	if result.fastlyprobeipv4 {
		t.fastlyprobeipv4 += 1
	}
	if result.fastlyprobeipv6 {
		t.fastlyprobeipv6 += 1
	}

}

const (
	errNumArgsMsgString = "Incorrect number of arguments, expecting 1 but received %d. Usage: ./analyze <domain>\n"

	http2NoSupportMsgString = "🚫 %s does not support HTTP/2\n"
	http2SupportMsgString   = "✅ %s supports HTTP/2\n"

	http1xSupportMsgString = "✅ %s supports HTTP/1.1\n"
	http10SupportMsgString = "✅ %s supports HTTP/1.0\n"
)

func worker(input chan string, output chan ProbeResult, cdn_fast probes.Fastlyprobe, cdn_cloud probes.Cloudflareprobe) {
	for x := range input {
		output <- filepathHTTP(x, &cdn_fast, &cdn_cloud)

	}
}

/*
create listener to prevent tcp error
instantiate before starting workers
*/
func fileEntry(filepath string, workers int, cdn_fast probes.Fastlyprobe, cdn_cloud probes.Cloudflareprobe) TotalTestResult {
	domains, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}

	domain := bufio.NewScanner(domains)

	jobs := make(chan string, 300)
	results := make(chan ProbeResult, 1000000)

	log.Printf("Running with %d goroutine workers\n", workers)

	for x := 0; x < workers; x++ {
		go func() {
			worker(jobs, results, cdn_fast, cdn_cloud)
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
		totalresults.AddResult(result)

		if count == resultCount {
			close(results)
		}
	}

	return totalresults
}

func filepathHTTP(myURL string, cdn_fast *probes.Fastlyprobe, cdn_cloud *probes.Cloudflareprobe) ProbeResult {
	result := ProbeResult{}

	http2Result := (&probes.HTTP2Probe{}).Run(myURL)
	result.errorhttp2occured = http2Result.Err != nil
	result.http2enabled = http2Result.Supported

	http1Result := (&probes.HTTP11Probe{}).Run(myURL)
	result.errorhttp1occured = http1Result.Err != nil
	result.http11enabled = http1Result.Supported

	TLS10Result := (&probes.TLS{}).Run(myURL, 0)
	result.tls10enabled = TLS10Result.Err != nil
	result.tls10enabled = TLS10Result.Supported

	TLS11Result := (&probes.TLS{}).Run(myURL, 1)
	result.tls11enabled = TLS11Result.Err != nil
	result.tls11enabled = TLS11Result.Supported

	TLS12Result := (&probes.TLS{}).Run(myURL, 2)
	result.tls12enabled = TLS12Result.Err != nil
	result.tls12enabled = TLS12Result.Supported

	TLS13Result := (&probes.TLS{}).Run(myURL, 3)
	result.tls13enabled = TLS13Result.Err != nil
	result.tls13enabled = TLS13Result.Supported

	Cloudflare := (cdn_cloud).Run(myURL)
	result.cloudflare = Cloudflare.Supported
	result.cloudflareipv4 = Cloudflare.Supportedipv4
	result.cloudflareipv6 = Cloudflare.Supportedipv6

	Fastly := (cdn_fast).Run(myURL)
	result.fastlyprobe = Fastly.Supported
	result.fastlyprobeipv4 = Fastly.Supportedipv4
	result.fastlyprobeipv6 = Fastly.Supportedipv6
	return result
}

func myfast() *probes.Fastlyprobe {
	fastly, err := irm.Sendfastlyprobe()

	if err != nil {
		fmt.Println("no fastly")
	}
	reader, err := ioutil.ReadAll(fastly.Body)
	if err != nil {
		fmt.Println("cant read line")
	}
	fast := probes.Fastlyprobe{}

	err = json.Unmarshal(reader, &fast)
	if err != nil {
		fmt.Println("cant marshal")
	}
	for _, x := range fast.Ipv4_addresses {
		//log.Println("scanning for ipv4")
		_, cidrsparse, _ := net.ParseCIDR(x)
		fast.Ipv4_addresses_cidr = append(fast.Ipv4_addresses_cidr, cidrsparse)

	}

	for _, x := range fast.Ipv6_addresses {
		//log.Println("scanning for ipv4")
		_, cidrsparse, _ := net.ParseCIDR(x)
		fast.Ipv6_addresses_cidr = append(fast.Ipv6_addresses_cidr, cidrsparse)

	}
	fmt.Println("run fast")
	return &fast

}
func mycloud() *probes.Cloudflareprobe {
	cidrsurlipv6, err := irm.Sendcloudflareipv6()
	cidrsurlipv4, err2 := irm.Sendcloudflareipv4()
	if err != nil {
		log.Println("nope on Sendcloudflareipv6")
	}
	if err2 != nil {
		log.Println("nope on Sendcloudflareipv4")
	}

	cidrsipv6 := bufio.NewScanner(cidrsurlipv6.Body)
	cidrsipv4 := bufio.NewScanner(cidrsurlipv4.Body)

	cloud := probes.Cloudflareprobe{}

	for cidrsipv4.Scan() {
		//log.Println("scanning for ipv4")
		_, cidrsparse, _ := net.ParseCIDR(cidrsipv4.Text())
		cloud.Ipv4_cidr = append(cloud.Ipv4_cidr, cidrsparse)
		//fmt.Println("run")
	}
	for cidrsipv6.Scan() {
		//log.Println("scanning for ipv4")
		_, cidrsparse, _ := net.ParseCIDR(cidrsipv6.Text())
		cloud.Ipv6_cidr = append(cloud.Ipv6_cidr, cidrsparse)

	}
	fmt.Println("cloud run")
	return &cloud
}
func main() {
	// Try to minimize filesystem usage and avoid lingering connections.
	http.DefaultTransport.(*http.Transport).DisableKeepAlives = true

	var filepath string
	var filepathexport string
	var numWorkers int
	var timebetrun int
	var enableGit int
	var singleDomain string
	flag.StringVar(&filepath, "f", "", "file path")
	flag.StringVar(&filepathexport, "o", "", "export to csv")
	flag.IntVar(&numWorkers, "w", runtime.NumCPU()*2, "number of workers")
	flag.IntVar(&timebetrun, "d", 10, "time between runs")
	flag.IntVar(&enableGit, "git", 0, "enable (1) git or disable (0)")
	flag.StringVar(&singleDomain, "domain", "", "test single domain")
	flag.Parse()

	cdn_fast := myfast()
	cdn_cloud := mycloud()
	if singleDomain != "" {
		test := filepathHTTP(singleDomain, cdn_fast, cdn_cloud)
		fmt.Printf("HTTP/1.1: %t \nHTTP/1.2: %t \nTLSv1.0: %t \nTLSv1.1: %t \nTLSv1.2: %t \nTLSv1.3: %t\n", test.http11enabled, test.http11enabled, test.tls10enabled, test.tls11enabled, test.tls12enabled, test.tls13enabled)
		os.Exit(0)
	}

	if filepath == "" {
		log.Fatal("Must specify -filepath or -domain")
	} else {
		for {
			timer := time.Now()
			totalresults := fileEntry(filepath, numWorkers, *cdn_fast, *cdn_cloud)
			testDuration := time.Since(timer).Seconds()
			domainsTested := totalresults.domainsTested
			data := [][]string{
				{time.Now().Format("2006-01-02 15:04:05"),
					fmt.Sprintf("%d", totalresults.domainsTested),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.http2enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.http11enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.erroroccured, domainsTested)),
					fmt.Sprintf("%.2fs", testDuration),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.tls10enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.tls11enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.tls12enabled, domainsTested)),
					fmt.Sprintf("%.2f%%", util.Percent(totalresults.tls13enabled, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.cloudflare, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.cloudflareipv4, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.cloudflareipv6, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.fastlyprobe, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.fastlyprobeipv4, domainsTested)),
					fmt.Sprintf("%.2f%%\n", util.Percent(totalresults.fastlyprobeipv6, domainsTested)),
				}}
			//added timer
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

			if enableGit == 1 {
				os.Setenv("SSH_KNOWN_HOSTS", "/home/.ssh/known_hosts")
				publicKeys, err := ssh.NewPublicKeysFromFile("git", "/home/.ssh/id_ed25519", "") //
				if err != nil {
					log.Printf("generate publickeys failed: %s\n", err.Error())
				}
				checkFile, err := os.Open(filepathexport)
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

				file, err := os.OpenFile(filepathexport, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
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
					log.Println("doesn't exist")
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
					log.Println("commit not working properly")
				}
				mrr = repo.Push(&git.PushOptions{
					RemoteName: "origin",
					Auth:       publicKeys,
				})
				log.Printf("errors that happened: %s", mrr)
			} else {
				if filepathexport != "" {
					file, err := os.OpenFile(filepathexport, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
					if err != nil {
						log.Println(err.Error() + "cant open results in tempirmdata")
					}

					writer := csv.NewWriter(file)

					for _, value := range data {
						writer.Write(value)
						log.Println(value)
					}
					writer.Flush()
					if err := writer.Error(); err != nil {
						log.Println(err.Error())
					}
					file.Close()
					fmt.Printf("test results written to %s\n", filepathexport)
				} else {
					fmt.Printf("Test Duration: %.2fs\n", testDuration)
					fmt.Printf("Success Rate: %.2f%%\n", util.Percent((domainsTested-totalresults.erroroccured), domainsTested))
					fmt.Printf("HTTP/1.1 enabled: %.2f%% \n", util.Percent(totalresults.http11enabled, domainsTested))
					fmt.Printf("HTTP/1.2 enabled: %.2f%%\n", util.Percent(totalresults.http2enabled, domainsTested))
					fmt.Printf("TLSv1.0 enabled: %.2f%%\n", util.Percent(totalresults.tls10enabled, domainsTested))
					fmt.Printf("TLSv1.1 enabled: %.2f%%\n", util.Percent(totalresults.tls11enabled, domainsTested))
					fmt.Printf("TLSv1.2 enabled: %.2f%%\n", util.Percent(totalresults.tls12enabled, domainsTested))
					fmt.Printf("TLSv1.3 enabled: %.2f%%\n", util.Percent(totalresults.tls13enabled, domainsTested))
					fmt.Printf("cloudflares enabled: %.2f%%\n", util.Percent(totalresults.cloudflare, domainsTested))
					fmt.Printf("cloudflares ipv4 enabled:  %.2f%%\n", util.Percent(totalresults.cloudflareipv4, domainsTested))
					fmt.Printf("cloudflares ipv6 enabled: %.2f%%\n", util.Percent(totalresults.cloudflareipv6, domainsTested))
					fmt.Printf("fastlyprobe enabled: %.2f%%\n", util.Percent(totalresults.fastlyprobe, domainsTested))
					fmt.Printf("fastlyprobe ipv4 enabled:  %.2f%%\n", util.Percent(totalresults.fastlyprobeipv4, domainsTested))
					fmt.Printf("fastlyprobe ipv6 enabled: %.2f%%\n", util.Percent(totalresults.fastlyprobeipv6, domainsTested))
				}
				os.Exit(0)
			}

			log.Printf("sleeping for %d seconds\b", timebetrun)
			time.Sleep(time.Duration(timebetrun) * time.Second)
		}
	}
}
