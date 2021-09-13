package probes

import (
	"bufio"
	"log"
	"net"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/1.1.
 */

type Cloudflareprobe struct{}

func (h *Cloudflareprobe) Run(domain string) *ProbeResult {
	enabled := false

	cidrsurl, err := irm.Sendcloudflare(domain)
	if err != nil {
		log.Println("nope")
	}
	cidrs := bufio.NewScanner(cidrsurl.Body)

	var arr2 []string
	for cidrs.Scan() {
		arr2 = append(arr2, cidrs.Text())
	}
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope")
	}
	for _, cidr := range arr2 {
		_, cidrsparse, _ := net.ParseCIDR(cidr)

		log.Println(domain)
		for _, x := range ips {

			if cidrsparse.Contains(x) {
				enabled = true
				break
			}
		}
		if enabled == true {
			break

		}

	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "cloudflare supported",
	}
}
