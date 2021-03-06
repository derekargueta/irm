package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/1.1.
 */

type HTTP11Probe struct{}

func (h *HTTP11Probe) Run(domain string) *ProbeResult {
	enabled := false

	response, err := irm.SendHTTP1Request(domain, "https://")
	if response != nil {
		response.Body.Close()
	}

	if err == nil {
		enabled = true
	} else {

		response, err := irm.SendHTTP1Request(domain, "http://")
		if response != nil {
			response.Body.Close()
		}

		if err == nil {
			enabled = true
		} else {
			log.Println(err.Error() + " by http1.1 request")
		}

	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "HTTP/1.1 supported",
	}
}
