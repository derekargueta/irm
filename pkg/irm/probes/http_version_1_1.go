package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/1.1.
 */

type HTTP11Probe struct {
	Domain string
}

func (h *HTTP11Probe) Run() *ProbeResult {
	enabled := false
	response1, err := irm.SendHTTP1Request(h.Domain)
	if response1 != nil {
		response1.Body.Close()
	}
	if err == nil {
		enabled = true
	} else {
		log.Println(err.Error() + " by http1.1 request")
	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "HTTP/1.1 supported",
	}
}
