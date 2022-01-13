package main_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
	"time"

	"golang.org/x/net/http2"
	// "golang.org/x/net/http2"
)

func TestMain(t *testing.T) {
	fmt.Println("started")
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS12}
	clientele := &http.Client{Transport: &http2.Transport{TLSClientConfig: tlsConfig}, Timeout: time.Second * 10}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://localhost:8081/status", nil)
	fmt.Println("middle")
	request.Close = true
	fmt.Println("almost")
	response, err := clientele.Do(request)
	fmt.Println("here")
	if err != nil {
		t.Error("err", err)
	} else {
		fmt.Println("response", response)
	}

}
