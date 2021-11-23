package probes

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type MaxCDN struct {
	Ipv4_cidr []*net.IPNet
}

/*
Doesnt work atm, supported Max CDN websites
dont return true when ran through cidr list

*/

func (h *MaxCDN) Run(domain string) *ProbeResult {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope on lookupIP")
	}

	enabledtotal := false
	_, httperr := http.NewRequest("GET", domain, nil)
	if httperr != nil {
		return &ProbeResult{

			Supported: false,
			Err:       httperr,
			Name:      "MaxCDN not supported",
		}
	}

	for _, x := range ips {

		if x.To4() != nil {
			for _, cidr := range h.Ipv4_cidr {
				if cidr.Contains(x) {

					enabledtotal = true
					fmt.Println("enabled true on maxcdn")
				}
			}

		}
	}

	return &ProbeResult{
		Supported: enabledtotal,
		Err:       err,
		Name:      "supported",
	}
}
