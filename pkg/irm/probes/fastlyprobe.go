package probes

import (
	"log"
	"net"
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

	ips, err2 := net.LookupIP(domain)
	if err2 != nil {
		log.Println("nope on lookupIP")
	}
	enabledtotal := false
	ipv4 := false
	ipv6 := false
	for _, x := range ips {
		if x.To4() != nil {
			for _, cidr := range h.Ipv4_addresses_cidr {
				if cidr.Contains(x) {
					log.Println("went thru ipv4", x)
					enabledtotal = true
					ipv4 = true
				}
			}

		} else {
			for _, cidr := range h.Ipv6_addresses_cidr {
				if cidr.Contains(x) {
					log.Println("went thru ipv6", x)
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
