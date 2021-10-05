package probes

import (
	"log"
	"net"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Fastlyprobe struct {
	Addresses      []string `json:"addresses"`
	Ipv6_addresses []string `json:"ipv6_addresses"`
}

func (h *Fastlyprobe) Run(domain string) *ProbeResultcloudfast {

	ips, err2 := net.LookupIP(domain)
	if err2 != nil {
		log.Println("nope on lookupIP")
	}
	enabledtotal := false
	ipv4 := false
	ipv6 := false
	for _, x := range ips {
		if x.To4() != nil {
			for _, fastlyip := range h.Addresses {
				log.Printf("scanning for ipv4")
				_, cidrsparse, _ := net.ParseCIDR(fastlyip)
				if cidrsparse.Contains(x) {
					log.Println("went thru ipv4 on fastly", x)
					enabledtotal = true
					ipv4 = true
				}

			}
		} else {
			for _, fastlyip := range h.Ipv6_addresses {
				_, cidrsparse, _ := net.ParseCIDR(string(fastlyip))
				if cidrsparse.Contains(x) {
					log.Println("went thru ipv6 on fastly ", x)
					ipv4 = false
					enabledtotal = true
					ipv6 = true
					log.Println(ipv6)

				}

			}
		}
	}

	return &ProbeResultcloudfast{
		Supported:     enabledtotal,
		Supportedipv4: ipv4,
		Supportedipv6: ipv6,
		Err:           err2,
		Name:          "fastly supported",
	}
}
