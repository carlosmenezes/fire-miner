package models

type BMinerResult struct {
	Stratum Stratum `json:"stratum"`
	Miners  []Miner
}

type Stratum struct {
	AcceptedShares int `json:"accepted_shares"`
	RejectedShares int `json:"rejected_shares"`
}

type Miner struct {
	Solver Solver `json:"solver"`
}

type Solver struct {
	SolutionRate float32 `json:"solution_rate"`
}
