package models

type EwbfResult struct {
	Error  string   `json:"error"`
	Result []Result `json:"result"`
}

type Result struct {
	GpuID          int    `json:"gpuid"`
	Name           string `json:"name"`
	Temperature    int    `json:"temperature"`
	PowerUsage     int    `json:"gpu_power_usage"`
	SolsPerSecond  int    `json:"speed_sps"`
	AcceptedShares int    `json:"accepted_shares"`
	RejectedShares int    `json:"rejected_shares"`
}
