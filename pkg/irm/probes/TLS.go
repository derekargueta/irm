package probes

import (
	"log"

	"github.com/derekargueta/irm/pkg/irm"
)

/**
 * Checks if the domain supports TeCP1.0.
 */
type TLS struct{}

func (h *TLS) Run(domain string, typeTLS int) *ProbeResult {
	enabled := false
	switch typeTLS {
	case 0:
		{
			response, err := irm.SendTLS10Request(domain)
			if response != nil {
				response.Body.Close()
			}

			if err == nil {
				enabled = true
			} else {
				log.Println(err, " tcp1.0 failed")
			}

			return &ProbeResult{
				Supported: enabled,
				Err:       err,
				Name:      "TLS 1/0 successful",
			}
		}

	case 1:
		{
			response, err := irm.SendTLS11Request(domain)
			if response != nil {
				response.Body.Close()
			}

			if err == nil {
				enabled = true
			} else {
				log.Println(err, " tcp1.1 failed")
			}

			return &ProbeResult{
				Supported: enabled,
				Err:       err,
				Name:      "TLS 1/1 successful",
			}
		}
	case 2:
		{
			response, err := irm.SendTLS12Request(domain)
			if response != nil {
				response.Body.Close()
			}

			if err == nil {
				enabled = true
			} else {
				log.Println(err, " tcp1.2 failed")
			}

			return &ProbeResult{
				Supported: enabled,
				Err:       err,
				Name:      "TLS 1/2 successful",
			}
		}

	case 3:
		{
			response, err := irm.SendTLS13Request(domain)
			if response != nil {
				response.Body.Close()
			}

			if err == nil {
				enabled = true
			} else {
				log.Println(err, " tcp1.2 failed")
			}

			return &ProbeResult{
				Supported: enabled,
				Err:       err,
				Name:      "TLS 1/2 successful",
			}
		}

	}
	return &ProbeResult{
		Supported: enabled,
		Err:       nil,
		Name:      "nil",
	}
}
