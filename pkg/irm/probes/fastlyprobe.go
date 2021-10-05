package probes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports cloudflare.
 */

type fastlyprobe struct {
	addresses      []string `json:"addresses"`
	ipv6_addresses []string `json:"ipv6_addresses"`
}

func (h *fastlyprobe) Run(domain string) *ProbeResult {
	fastly, err := irm.Sendfastlyprobe()
	if err != nil {
		fmt.Println("no fastly")
	}
	reader, err := ioutil.ReadAll(fastly.Body)
	if err != nil {
		fmt.Println("cant read line")
	}
	var inter interface{}
	err = json.Unmarshal(reader, &inter)
	if err != nil {
		fmt.Println("cant marshal")
	}

	fmt.Println(inter)
	return &ProbeResult{
		Supported: true,
		Err:       err,
		Name:      "cloudflare supported",
	}
}
