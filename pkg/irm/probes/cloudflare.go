package probes

import (
	"log"
	"net"
	"net/http"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Cloudflareprobe struct {
	Ipv4_cidr []*net.IPNet
	Ipv6_cidr []*net.IPNet
}

//verify http request with ip
func (h *Cloudflareprobe) Run(domain string) *ProbeResultCDN {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope on lookupIP in cloudflare probe")
	}
	enabledtotal := false
	ipv4 := false
	ipv6 := false
	_, httperr := http.NewRequest("GET", domain, nil)
	if httperr != nil {
		return &ProbeResultCDN{
			Supported:     false,
			Supportedipv4: false,
			Supportedipv6: false,
			Err:           err,
			Name:          "cloudflare not supported",
		}
	}

	for _, x := range ips { //sep dns

		if x.To4() != nil { //is ipv4
			for _, cidrsparse := range h.Ipv4_cidr {
				if cidrsparse.Contains(x) {
					//log.Println("went thru ipv4", x)
					ipv4 = true
					ipv6 = false
					enabledtotal = true
				}
			}

		} else {
			for _, cidrsparse := range h.Ipv6_cidr {
				if cidrsparse.Contains(x) {
					//log.Println("went thru ipv6", x)
					ipv6 = true
					ipv4 = false
					enabledtotal = true
					//log.Println(ipv6)

				}

			}
		}
	}

	//ipv6 returns false (good) but ipv4 doesn't
	//fmt.Println(enabledtotal, ipv4, ipv6)
	return &ProbeResultCDN{
		Supported:     enabledtotal,
		Supportedipv4: ipv4,
		Supportedipv6: ipv6,
		Err:           err,
		Name:          "cloudflare supported",
	}
}
