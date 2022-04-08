package probes

import (
	"fmt"
	"net"

	"github.com/likexian/whois"
)

/**
 * Checks if the domain supports cloudflare.
 */

type dns_NS struct {
	Ipv4_cidr []*net.IPNet
	Ipv6_cidr []*net.IPNet
}

//verify http request with ip
func (h *dns_NS) Run(domain string) *ProbeResult {
	//nameserver, err := net.LookupNS(domain)
	result, err := whois.Whois(domain)

	if err != nil {
		fmt.Println("cant return name server")
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	return &ProbeResult{
		Supported: !false, //if false, means dns didnt respond to query
		Err:       err,
		Name:      "cloudflare not supported",
	}
}
