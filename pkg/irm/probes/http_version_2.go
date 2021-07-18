package probes

import (
	"log"
	"strings"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports HTTP/2.
 */
type HTTP2Probe struct{}

func (h *HTTP2Probe) Run(domain string) *ProbeResult {
	enabled := false
	response, err := irm.SendHTTP2Request(domain)

	if response != nil {
		response.Body.Close()
	}

	if err == nil {
		enabled = true
	} else {
		// There are certain "errors" that are normal indicators of lacking HTTP/2 support. We're not
		// interested in those - but if it's a different error, truly exceptional in that we don't
		// expect it, then let's log it to investigate and understand the failure mode.
		errOtherThanHTTP2Support := !strings.Contains(err.Error(), "unexpected ALPN protocol")
		if errOtherThanHTTP2Support {
			log.Println(err.Error() + " - request error for http2")
		}
	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "HTTP/2 supported",
	}
}
