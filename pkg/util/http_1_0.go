package util

import (
	"fmt"
	"os/exec"
	"strings"
)

// Seems that the Go standard library does not support HTTP/1.0 requests?
// https://github.com/golang/go/blob/d4f34f8c63b753160716e9f90ca530016ce019d7/src/net/http/transport.go#L74-L78
func Http10Request(url string) bool {
	out, err := exec.Command("curl", "-svo", "/dev/null", "--no-keepalive", "--http1.0", url).CombinedOutput()
	if err != nil {
		fmt.Println("error in http1.0 " + err.Error())
	}

	strOut := string(out)
	return strings.Contains(strOut, "< HTTP/1.0")
}
