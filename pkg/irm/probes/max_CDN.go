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
Maxcdn is now stackpath

*/

func (h *MaxCDN) Run(domain string) *ProbeResult {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("nope on lookupIP")
	}

	enabledtotal := false
	_, httperr := http.NewRequest("GET", domain, nil)
	if httperr != nil {
		fmt.Println("http error on maxcdn")
		return &ProbeResult{

			Supported: false,
			Err:       httperr,
			Name:      "MaxCDN not supported",
		}
	}

	for _, x := range ips {

		if x.To4() != nil {

			for _, cidr := range h.Ipv4_cidr {
				if cidr.String() == "151.139.0.0/17" { //for some reason, this stackpath ip doesn't return it's organization when testing with stackpath websites
					_, temp, err := net.ParseCIDR("151.139.0.0/16") //however, changing the port to this returns it
					if err != nil {
						fmt.Println("error on maxcdn on 151.139.0.0/16")
					}
					if temp.Contains(x) {
						enabledtotal = true
						fmt.Println("FOUND TRUE ON MAXCDN")
						break
					}
				}

				if cidr.Contains(x) {
					enabledtotal = true
					fmt.Println("FOUND TRUE ON MAXCDN")
					break
				}

			}
			if enabledtotal == true {
				break
			}

		}
	}

	return &ProbeResult{
		Supported: enabledtotal,
		Err:       err,
		Name:      "supported",
	}
}
