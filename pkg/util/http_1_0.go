package util

import (
	"strings"
	"os/exec"
)

// Seems that the Go standard library does not support HTTP/1.0 requests?
// https://github.com/golang/go/blob/d4f34f8c63b753160716e9f90ca530016ce019d7/src/net/http/transport.go#L74-L78
func Http10Request(url string) bool {
	out, err := exec.Command("curl", "-svo", "/dev/null", "--http1.0", url).CombinedOutput()
	if err != nil {
		panic(err)
	}

	strOut := string(out)
	return strings.Contains(strOut, "< HTTP/1.0")
}
