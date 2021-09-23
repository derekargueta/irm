package probes

import (
	"bufio"
	"log"
	"net"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports cloudflare.
 */

type Cloudflareprobe struct{}

func (h *Cloudflareprobe) Run(domain string) *ProbeResultcloudflare {
	cidrsurlipv6, err := irm.Sendcloudflareipv6()
	cidrsurlipv4, err2 := irm.Sendcloudflareipv4()
	if err != nil {
		log.Println("nope on Sendcloudflareipv6")
	}
	if err2 != nil {
		log.Println("nope on Sendcloudflareipv4")
	}

	cidrsipv6 := bufio.NewScanner(cidrsurlipv6.Body)
	cidrsipv4 := bufio.NewScanner(cidrsurlipv4.Body)

	ips, err2 := net.LookupIP(domain)
	if err2 != nil {
		log.Println("nope on lookupIP")
	}

	// enabledipv4 := false
	// enabledipv6 := false
	// for _, x := range ips {
	// 	if x.To4() != nil {
	// 		for cidrsipv4.Scan() {
	// 			//log.Println("scanning for ipv4")
	// 			_, cidrsparse, _ := net.ParseCIDR(cidrsipv4.Text())
	// 			if cidrsparse.Contains(x) {
	// 				//log.Println("oiasehjfuhefpweufgwoejfno;hjfvlosodhfpawusehfgawoehgp;owehfgoaeuhfguasdhfoasudhflo;sdfh")
	// 				enabledipv4 = true
	// 			}

	// 		}
	// 	} else {
	// 		for cidrsipv6.Scan() {
	// 			_, cidrsparse, _ := net.ParseCIDR(cidrsipv6.Text())
	// 			if cidrsparse.Contains(x) {
	// 				//log.Println("oiasehjfuhefpweufgwoejfno;hjfvlosodhfpawusehfgawoehgp;owehfgoaeuhfguasdhfoasudhflo;sdfh")
	// 				enabledipv6 = true
	// 			}

	// 		}
	// 	}
	// }

	return &ProbeResultcloudflare{
		Supported:     cidrContains(cidrsipv4, cidrsipv6, ips),
		Supportedipv4: false, //not finished
		Supportedipv6: false, //not finished
		Err:           err2,
		Name:          "cloudflare supported",
	}
}

func cidrContains(cidrsipv4 *bufio.Scanner, cidrsipv6 *bufio.Scanner, ips []net.IP) bool {
	for _, x := range ips {
		if x.To4() != nil {
			for cidrsipv4.Scan() {
				//log.Println("scanning for ipv4")
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv4.Text())
				if cidrsparse.Contains(x) {
					//log.Println("oiasehjfuhefpweufgwoejfno;hjfvlosodhfpawusehfgawoehgp;owehfgoaeuhfguasdhfoasudhflo;sdfh")
					return true
				}

			}
		} else {
			for cidrsipv6.Scan() {
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv6.Text())
				if cidrsparse.Contains(x) {
					//log.Println("oiasehjfuhefpweufgwoejfno;hjfvlosodhfpawusehfgawoehgp;owehfgoaeuhfguasdhfoasudhflo;sdfh")
					return true
				}

			}
		}

	}
	return false
}
