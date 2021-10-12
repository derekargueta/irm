package probes

import (
	"log"
	"net"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Cloudflareprobe struct {
	Ipv4 []*net.IPNet
	Ipv6 []*net.IPNet
}

func (h *Cloudflareprobe) Run(domain string) *ProbeResultcloudfast {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope on lookupIP")
	}

	enabledtotal := false
	ipv4 := false
	ipv6 := false
	for _, x := range ips {
		if x.To4() != nil {
			for _, cidr := range h.Ipv4 {
				if cidr.Contains(x) {
					log.Println("went thru ipv4", x)
					enabledtotal = true
					ipv4 = true
				}
			}

		} else {
			for _, cidr := range h.Ipv6 {
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
