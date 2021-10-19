package probes

import (
	"fmt"
	"net"
	"strings"

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
	ah := false
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeANY)

	in, err := dns.Exchange(m, "1.1.1.1:53")
	if err != nil {
		fmt.Println(err)
	} else {
		yes := in.String()
		ah = strings.Contains(yes, "NOTIMP")
		/*
			if true, means dns query is blocked
		*/

	}

	return &ProbeResult{
		Supported: !ah, //if false, means dns didnt respond to query
		Err:       err,
		Name:      "cloudflare not supported",
	}
}
