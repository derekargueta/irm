package probes

type ProbeResult struct {
	Supported bool
	Err       error
	Name      string
}
type ProbeResultcloudflare struct {
	Supported     bool
	Supportedipv4 bool
	Supportedipv6 bool
	Err           error
	Name          string
}

type Probe interface {
	Run(string) *ProbeResult
}
