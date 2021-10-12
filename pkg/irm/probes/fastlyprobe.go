package probes

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Fastlyprobe struct {
	Ipv4_addresses []string `json:"addresses"`
	Ipv6_addresses []string `json:"ipv6_addresses"`

	Ipv4_addresses_cidr []*net.IPNet
	Ipv6_addresses_cidr []*net.IPNet
}

func (h *Fastlyprobe) Run(domain string) *ProbeResultcloudfast {

	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope on lookupIP")
	}
	fmt.Println(domain)
	enabledtotal := false
	ipv4 := false
	ipv6 := false
	_, httperr := http.NewRequest("GET", domain, nil)
	if httperr != nil {
		return &ProbeResultcloudfast{
			Supported:     false,
			Supportedipv4: false,
			Supportedipv6: false,
			Err:           err,
			Name:          "fastly not supported",
		}
	}
	for _, x := range ips {
		if x.To4() != nil {
			for _, cidr := range h.Ipv4_addresses_cidr {
				if cidr.Contains(x) {
					//log.Println("went thru ipv4", x)
					enabledtotal = true
					ipv4 = true
					ipv6 = false
				}
			}

		} else {
			for _, cidr := range h.Ipv6_addresses_cidr {
				if cidr.Contains(x) {
					//log.Println("went thru ipv6", x)
					enabledtotal = true
					ipv6 = true
					ipv4 = false
					//log.Println(ipv6)

				}

			}
		}
	}

	return &ProbeResultcloudfast{
		Supported:     enabledtotal,
		Supportedipv4: ipv4,
		Supportedipv6: ipv6,
		Err:           err,
		Name:          "fastly supported",
	}
}
