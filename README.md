Internet Measurement Research
=============================

yeah the repo should probably be "imr", but I didn't realize it and I like "irm" better.

This repo will eventually be a system that collects data about various protocol usage on the internet.

Project layout is based on https://github.com/golang-standards/project-layout

# Questions to answer
## HTTP
- [ ] How many HTTP services have HTTP/1.0 enabled?
- [x] How many HTTP services have HTTP/1.1 enabled?
- [x] How many HTTP services have HTTP/2 enabled?
- [x] How many HTTP services support HTTP/3?
- [ ] How many HTTP services support UDP on :443? (prerequisite to QUIC/HTTP3)
- [ ] How many HTTP services support plaintext HTTP on port 80?

## IP
Use DNS answers to check this (A records vs AAAA records), then connect over that IP address to confirm it is valid.
- [x] How many HTTP services are served over IPv6?
- [x] How many HTTP services are served over IPv4?
- [x] How many HTTP services are served over both?

## TLS
- [x] How many TLS-enabled HTTP services support TLS 1.0? (officially deprecated via [RFC 8996](https://datatracker.ietf.org/doc/rfc8996/))
- [x] How many TLS-enabled HTTP services support TLS 1.1? (officially deprecated via [RFC 8996](https://datatracker.ietf.org/doc/rfc8996/))
- [x] How many TLS-enabled HTTP services support TLS 1.2?
- [x] How many TLS-enabled HTTP services support TLS 1.3?
- [x] What is the distribution of root CA providers? (Digicert, Verisign, Komodo, etc.)

## CDNs/Datacenters
Use the connecting IP address to approximate where the response is originating from, for example checking against [Cloudflare's public IP range](https://www.cloudflare.com/ips/).
In order for this to be accurate, the probes will need to be run from multiple regions to account for multi-CDN architectures.
Initial version will just be a single region to start with.
- [ ] What percentage of HTTP services are served over Akamai?
- [x] What percentage of HTTP services are served over Cloudflare?
- [x] What percentage of HTTP services are served over Fastly?
- [x] What percentage of HTTP services are served over MaxCDN?
- [ ] What percentage of HTTP services are served directly from AWS?
- [ ] What percentage of HTTP services are served directly from Google Cloud?
- [ ] What percentage of HTTP services are served directly from Azure?

## DNS
- [ ] What percentage of domains use NS1 as an authoritative server?
- [ ] What percentage of domains use Dyn/Oracle as an authoritative server?
- [x] What percentage of domains use Cloudflare as an authoritative server?
- [x] What percentage of domains respond to `ANY` queries? (notorious for [DNS amplification attacks](https://www.cloudflare.com/learning/ddos/dns-amplification-ddos-attack/))
- [ ] How long are CNAME chains? (broken into percentiles)

## Misc.
- [ ] Websocket use in the wild?
- [ ] How many websites use Google Analytics?
- [ ] Use of new HTML features? e.g. the `<picture>` tag, `<details>` tag, `<summary>` tag, etc.

# Tools

## Analyze
`analyze` is a command-line tool that accepts a domain as input, and runs a protocol inspection against it.
It is the basic building block for this internet measurement research.

Run the analyze tool as a "Go script":
```
$ go run ./cmd/analyze news.ycombinator.com
🚫 news.ycombinator.com does not support HTTP/2
$ go run ./cmd/analyze www.pinterest.com
✅ www.pinterest.com supports HTTP/2
```

or compile the executable:
```
$ go build ./cmd/analyze
$ ./analyze news.ycombinator.com
🚫 news.ycombinator.com does not support HTTP/2
$ ./analyze www.pinterest.com
✅ www.pinterest.com supports HTTP/2
```

# TODO
Derek
- [ ] lab homepage
- [ ] server deployment tooling
