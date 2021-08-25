package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports TeCP1.0.
 */
type TCP1 struct{}

func (h *TCP1) Run(domain string) *ProbeResult {
	enabled := false

	response, err := irm.SendTCP1Request(domain)
	if response != nil {
		response.Body.Close()
	}

	if err == nil {
		enabled = true
	} else {
		log.Println("tcp1 failed")
	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "TCP 1/0 successful",
	}
}
