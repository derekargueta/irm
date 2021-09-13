package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/1.1.
 */

type cloudflareprobe struct{}

func (h *cloudflareprobe) Run(domain string) *ProbeResult {
	enabled := false

	response, err := irm.Sendcloudflarerequest(domain)
	if response != nil {
		response.Body.Close()
	}

	if err == nil {
		enabled = true
	} else {

		response, err := irm.Sendcloudflarerequest(domain)
		if response != nil {
			response.Body.Close()
		}

		if err == nil {
			enabled = true
		} else {
			log.Println(err.Error() + " by cloudflare request")
		}

	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "cloudflare supported",
	}
}
