Internet Measurement Research
=============================

yeah the repo should probably be "imr", but I didn't realize it and I like "irm" better.

This repo will eventually be a system that collects data about various protocol usage on the internet.

Project layout is based on https://github.com/golang-standards/project-layout

Questions to answer
- [ ] How many HTTP services have HTTP/1.0 enabled?
- [ ] How many HTTP services have HTTP/2 enabled?
- [ ] Are there HTTP services that do _not_ have HTTP/1.1 enabled?
- [ ] How many HTTP services support HTTP/3?
- [ ] How many HTTP services are served over IPv6? IPv4? Both?
- [ ] How many TLS-enabled HTTP services support TLS 1.0?
- [ ] How many TLS-enabled HTTP services support TLS 1.1?
- [ ] How many TLS-enabled HTTP services support TLS 1.2?
- [ ] How many TLS-enabled HTTP services support TLS 1.3?
- [ ] What percentage of HTTP services are served over Akamai?
- [ ] What percentage of HTTP services are served over Cloudflare?
- [ ] What percentage of HTTP services are served over Fastly?
- [ ] What percentage of HTTP services are served directly from AWS?
- [ ] What percentage of HTTP services are served directly from Google Cloud?
- [ ] What percentage of HTTP services are served directly from Azure?

# Tools

## Analyze
`analyze` is a command-line tool that accepts a domain as input, and runs a protocol inspection against it.
It is the basic building block for this internet measurement research.

Run the Go as a "Go script":
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
