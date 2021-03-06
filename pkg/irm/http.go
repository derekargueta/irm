package irm

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/lucas-clemente/quic-go/http3"
	"golang.org/x/net/http2"
)

/*
 * HTTP utility stuff.
 */

func SendHTTP1Request(domain string, http_str string) (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TODO(derekargueta): if TLS fails, try HTTP/1 without TLS.

	request, _ := http.NewRequest("GET", http_str+domain, nil)
	request.Close = true
	return client.Do(request)
}

func SendHTTP2Request(domain string) (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	client := &http.Client{Transport: &http2.Transport{}, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return client.Do(request)
}

func SendHTTP3Request(domain string) (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10}
	client := &http.Client{Transport: &http3.RoundTripper{}, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return client.Do(request)
}

func SendTLS10Request(domain string) (*http.Response, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS10, MaxVersion: tls.VersionTLS10}
	//client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 5 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return clientele.Do(request)
}

func SendTLS11Request(domain string) (*http.Response, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS11, MaxVersion: tls.VersionTLS11}
	//client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 5 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return clientele.Do(request)
}

func SendTLS12Request(domain string) (*http.Response, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS12}
	//client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 5 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return clientele.Do(request)
}
func SendTLS13Request(domain string) (*http.Response, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13}
	//client := &http.Client{Transport: &http2.Transport{}, Timeout: 10 * time.Second}
	clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 5 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www."+domain, nil)
	request.Close = true
	return clientele.Do(request)
}

func Sendcloudflareipv6() (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www.cloudflare.com/ips-v6", nil)
	request.Close = true
	return client.Do(request)
}
func Sendcloudflareipv4() (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://www.cloudflare.com/ips-v4", nil)
	request.Close = true
	return client.Do(request)
}
func Sendfastlyprobe() (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://api.fastly.com/public-ip-list", nil)
	request.Close = true
	return client.Do(request)
}
func SendMaxCdnprobe() (*http.Response, error) {
	//tlsConfig := &tls.Config{MinVersion: tls.VersionTLS13, MaxVersion: tls.VersionTLS13}
	client := &http.Client{Transport: http.DefaultTransport, Timeout: 5 * time.Second}
	//clientele := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}, Timeout: 10 * time.Second}
	// TLS is required for public HTTP/2 services, so assume `https`.
	request, _ := http.NewRequest("GET", "https://support.stackpath.com/hc/en-us/article_attachments/360096407372/ipblocks.txt", nil)
	request.Close = true
	return client.Do(request)
}
