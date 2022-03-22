package probes

import (
	"fmt"
	"os/exec"
	"strings"
)

type Http10probe struct{}

func (h *Http10probe) Run(domain string) *ProbeResult {
	out, err := exec.Command("curl", "-svo", "/dev/null", "--no-keepalive", "--http1.0", domain).CombinedOutput()
	enabled := false
	if err != nil {
		fmt.Println("error in http1.0 " + err.Error())

	} else {
		if strings.Contains(string(out), "HTTP/1.0 200 OK") {
			enabled = true

		} else {
			out2, err := exec.Command("curl", "-svo", "/dev/null", "--no-keepalive", "--http1.0", "https://"+domain).CombinedOutput()

			if err != nil {
				fmt.Println("error in http1.0 " + err.Error())

			}

			if strings.Contains(string(out2), "HTTP/1.0 200 OK") {
				enabled = true

			}
		}
	}

	return &ProbeResult{
		Supported: enabled,
		Err:       err,
		Name:      "HTTP/1.0 supported",
	}
}
