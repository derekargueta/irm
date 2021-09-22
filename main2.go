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
	// oneurl("cloudflare.com")
	listurls()
}

func oneurl(myurl string) {
	cidrsurl, err := http.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		log.Println("nope")
	}
	cidrs := bufio.NewScanner(cidrsurl.Body)
	defer cidrsurl.Body.Close()

	for cidrs.Scan() {
		_, cidrsparse, _ := net.ParseCIDR(cidrs.Text())

		ips, err := net.LookupIP(myurl)
		if err != nil {
			log.Println("nope")
		}
		for _, x := range ips {
			if cidrsparse.Contains(x) {
				log.Println(cidrsparse, " yes")
			} else {
				log.Println(cidrsparse, "no")
			}
		}

	}
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
