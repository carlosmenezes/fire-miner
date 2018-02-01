package models

type BMinerResult struct {
	Stratum Stratum     `json:"stratum"`
	Miners  interface{} `json:"miners"`
}

type Stratum struct {
	AcceptedShares int `json:"accepted_shares"`
	RejectedShares int `json:"rejected_shares"`
}

type Miner struct {
	Id     string
	Solver Solver
	Device Device
}

type Solver struct {
	SolutionRate float64 `json:"solution_rate"`
}

type Device struct {
	Temperature float64
	Power       float64
	CoreClock   float64
	MemoryClock float64
}

func (b *BMinerResult) MinersList() []Miner {

	var miners []Miner
	minersMap := b.Miners.(map[string]interface{})

	for key, value := range minersMap {
		miner := Miner{}
		createMiner(key, value.(map[string]interface{}), &miner)

		miners = append(miners, miner)
	}

	return miners
}

func createMiner(id string, attributes map[string]interface{}, miner *Miner) {
	miner.Id = id

	for key, value := range attributes {
		switch key {
		case "solver":
			miner.Solver = createSolver(value.(map[string]interface{}))
		case "device":
			miner.Device = createDevice(value.(map[string]interface{}))
		}
	}
}

func createSolver(attributes map[string]interface{}) Solver {

	solver := Solver{}
	solver.SolutionRate = attributes["solution_rate"].(float64)

	return solver
}

func createDevice(attributes map[string]interface{}) Device {

	device := Device{}
	device.Temperature = attributes["temperature"].(float64)
	device.Power = attributes["power"].(float64)
	device.CoreClock = attributes["clocks"].(map[string]interface{})["core"].(float64)
	device.MemoryClock = attributes["clocks"].(map[string]interface{})["memory"].(float64)

	return device
}
