package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/2.
 */
type HTTP3Probe struct{}

func (h *HTTP3Probe) Run(domain string) *ProbeResult {
	enabled := false

	response, err := irm.SendHTTP3Request(domain)
	if response != nil {
		response.Body.Close()
	}

	if err == nil {
		enabled = true
	} else {

		log.Println(err.Error() + " - request error for http3")

	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "HTTP/3 supported",
	}
}
