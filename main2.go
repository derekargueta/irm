package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// A := "172.17.0.0/16"
	// B := "172.17.0.2"
	oneurl("cloudflare.com")

}
func yes()
{
	ips, err2 := net.LookupIP("cloudflare.com")
	if err2 != nil {
		log.Println("nope on lookupIP")
	}

	for _, x := range ips {
		if x.To4() != nil {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
}
func oneurl(myurl string) {
	enabled := false
	cidrsv4, err := http.Get("https://www.cloudflare.com/ips-v4")
	cidrsv6, err := http.Get("https://www.cloudflare.com/ips-v6")
	if err != nil {
		log.Println("nope")
	}
	cidrsipv4 := bufio.NewScanner(cidrsv4.Body)
	cidrsipv6 := bufio.NewScanner(cidrsv6.Body)
	//defer cidrsurl.Body.Close()
	ips, err := net.LookupIP(myurl)
	if err != nil {
		log.Println("nope")
	}
	for _, x := range ips {
		if x.To4() != nil {
			for cidrsipv4.Scan() {
				//log.Println("scanning for ipv4")
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv4.Text())
				if cidrsparse.Contains(x) {
					enabled = true
					break
				}
			}
		} else {
			for cidrsipv6.Scan() {
				_, cidrsparse, _ := net.ParseCIDR(cidrsipv6.Text())
				if cidrsparse.Contains(x) {

					enabled = true
					log.Println("ipv6")
					break
				}
			}
		}
		if enabled == true {
			break
		}

	}

	// for cidrs.Scan() {
	// 	_, cidrsparse, _ := net.ParseCIDR(cidrs.Text())

	// 	ips, err := net.LookupIP(myurl)
	// 	if err != nil {
	// 		log.Println("nope")
	// 	}
	// 	for _, x := range ips {
	// 		if cidrsparse.Contains(x) {
	// 			log.Println(cidrsparse, " yes")
	// 			enabled = true
	// 			break
	// 		} else {
	// 			log.Println(cidrsparse, "no")
	// 		}

	// 	}
	// 	if enabled == true {
	// 		break
	// 	}

	// }
	fmt.Println(enabled)
}
func listurls() {
	domains, erroropen := os.Open("/Users/Tavo/Documents/irm/domains/me.txt")
	if erroropen != nil {
		log.Println("nah")
	}
	domainlist := bufio.NewScanner(domains)
	cidrsurl, err := http.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		log.Println("nope")
	}
	cidrs := bufio.NewScanner(cidrsurl.Body)

	var arr []string
	for domainlist.Scan() {
		arr = append(arr, domainlist.Text())
	}

	var arr2 []string
	for cidrs.Scan() {
		arr2 = append(arr2, cidrs.Text())
	}
	counter := 0
	on := false
	// on2 := false
	for _, domain := range arr {
		ips, err := net.LookupIP(domain)
		if err != nil {
			log.Println("nope")
		}
		for _, cidr := range arr2 {
			_, cidrsparse, _ := net.ParseCIDR(cidr)
			log.Println("current cidrs parse", cidr)

			log.Println(domain)
			for _, x := range ips {
				fmt.Println("my total is ", counter)
				if cidrsparse.Contains(x) {
					fmt.Println(cidrsparse, " yes")
					counter++
					on = true
					break
				} else {
					log.Println(cidrsparse, " no")

				}
			}
			if on == true {
				on = false
				break

			}

		}

	}
	fmt.Println(counter)
}
