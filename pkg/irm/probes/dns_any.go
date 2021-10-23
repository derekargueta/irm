package probes

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Dns_any struct {
	Ipv4_cidr []*net.IPNet
	Ipv6_cidr []*net.IPNet
}

//verify http request with ip
func (h *Dns_any) Run(domain string) *ProbeResult {
	client := new(dns.Client)
	ah := false
	nameserver, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println("cant return name server")
	}
	for _, x := range nameserver {
		m := new(dns.Msg)
		m.SetQuestion(dns.Fqdn(domain), dns.TypeANY)
		in, time, err := client.Exchange(m, x.Host+":53")
		fmt.Println(x.Host)
		if err != nil {
			fmt.Println("on exchange", err)
		} else {
			yes := in.String()
			fmt.Println(yes)
			fmt.Println(time)
			/*
				if true, means dns query is blocked
			*/
			if ah == true {
				break
			}

		}
	}
	fmt.Println(ah)

	return &ProbeResult{
		Supported: !ah, //if false, means dns didnt respond to query
		Err:       err,
		Name:      "cloudflare not supported",
	}
}
