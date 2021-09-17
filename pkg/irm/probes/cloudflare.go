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

	cidrsurl, err := irm.Sendcloudflare()
	if err != nil {
		log.Println("nope on sendcloudflare")
	}

	cidrs := bufio.NewScanner(cidrsurl.Body)

	var arr2 []string
	for cidrs.Scan() {
		arr2 = append(arr2, cidrs.Text())
	}
	ips, err2 := net.LookupIP(domain)
	if err2 != nil {
		log.Println("nope on lookupIP")
	}
	for _, cidr := range arr2 {
		_, cidrsparse, err3 := net.ParseCIDR(cidr)
		if err3 != nil {
			log.Println("cant parse")
		}
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
