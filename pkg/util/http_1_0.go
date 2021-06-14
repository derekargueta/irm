package util

import (
	"strings"
	"os/exec"
)

func Http10Request(url string) bool {
	out, err := exec.Command("curl", "-svo", "/dev/null", "--http1.0", url).CombinedOutput()
	if err != nil {
		panic(err)
	}

	strOut := string(out)
	return strings.Contains(strOut, "< HTTP/1.0")
}
