package probes

type ProbeResult struct {
	Supported bool
	Err				error
	Name      string
}

type Probe interface {
	Run() *ProbeResult
}
