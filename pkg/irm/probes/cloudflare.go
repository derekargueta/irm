package probes

import (
	"bufio"
	"log"
	"net"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Cloudflareprobe struct {
	ipv4 []string
	ipv6 []string
}

func (h *Cloudflareprobe) Run(domain string) *ProbeResultcloudfast {
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

	ips, err2 := net.LookupIP(domain)
	if err2 != nil {
		log.Println("nope on lookupIP")
	}
	enabledtotal := false
	ipv4 := false
	ipv6 := false
	for _, x := range ips {
		if x.To4() != nil {
			for cidrsipv4.Scan() {
				//log.Println("scanning for ipv4")
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv4.Text())
				if cidrsparse.Contains(x) {
					log.Println("went thru ipv4", x)
					enabledtotal = true
					ipv4 = true
				}

			}
		} else {
			for cidrsipv6.Scan() {
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv6.Text())
				if cidrsparse.Contains(x) {
					log.Println("went thru ipv6", x)
					ipv4 = false
					enabledtotal = true
					ipv6 = true
					log.Println(ipv6)

				}

			}
		}
	}

	//ipv6 returns false (good) but ipv4 doesn't
	//fmt.Println(enabledtotal, ipv4, ipv6)
	return &ProbeResultcloudfast{
		Supported:     enabledtotal,
		Supportedipv4: ipv4,
		Supportedipv6: ipv6,
		Err:           err,
		Name:          "cloudflare supported",
	}
}
