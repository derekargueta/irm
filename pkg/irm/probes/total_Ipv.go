package probes

import (
	"log"
	"net"
	"net/http"
	"strings"
)

/**
 * Checks if the domain supports TeCP1.0.
 */
type Total_ipv struct{}

func (h *Total_ipv) Run(domain string) *TotalIpvProbe {
	enabled := false
	ipv4 := false
	ipv6 := false
	request, err := net.LookupIP(domain)
	_, httperr := http.NewRequest("GET", domain, nil)
	if httperr != nil {
		return &TotalIpvProbe{
			Dualstack:     false,
			Supportedipv4: false,
			Supportedipv6: false,
			Err:           err,
			Name:          "not supported",
		}
	}
	for _, x := range request {
		if x.To4() != nil {
			ipv4 = true
		} else if strings.Count(x.String(), ":") >= 2 {
			ipv6 = true
		}
	}

	if ipv4 && ipv6 {
		enabled = true
	}
	if err != nil {
		log.Println("cant get dns")
	}

	return &TotalIpvProbe{
		Dualstack:     enabled,
		Supportedipv4: ipv4,
		Supportedipv6: ipv6,
		Err:           nil,
		Name:          "nil",
	}
}
