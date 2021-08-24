package irm

import (
	"crypto/tls"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

/*
 * HTTP utility stuff.
 */

func SendHTTP1Request(domain string, http_str string) (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 10 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TODO(derekargueta): if TLS fails, try HTTP/1 without TLS.

	request, _ := http.NewRequest("GET", http_str+domain, nil)
	request.Close = true
	return client.Do(request)
}

func SendHTTP2Request(domain string) (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://"+domain, nil)
	request.Close = true
	return client.Do(request)
}

func SendTCP1Request(domain string) (*http.Response, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	//client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://"+domain, nil)
	request.Close = true
	return clientele.Do(request)
}
