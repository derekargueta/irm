package probes

type ProbeResult struct {
	Supported bool
	Err       error
	Name      string
}
type ProbeResultCDN struct {
	Supported     bool
	Supportedipv4 bool
	Supportedipv6 bool
	Err           error
	Name          string
}
type TotalIpvProbe struct {
	Dualstack     bool
	Supportedipv4 bool
	Supportedipv6 bool
	Err           error
	Name          string
}
type TotalTlsCertify struct {
	Digicert bool
	Comodo   bool
	Encrypt  bool
	Amazon   bool
	Err      error
}

type Probe interface {
	Run(string) *ProbeResult
}
