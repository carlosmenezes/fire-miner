package messageCreator

import (
	"fire-miner/app/models"
	"fmt"
	"os"
)

const baseEwbfMessage = `*GPU%d:* %s
  *Temperature:* %d C
  *Power usage:* %d W
  *Performance:* %d Sol/s
  *Efficiency:* %.2f Sol/W
  *Shares:*
    %d accepted, %d rejected
`

const baseBMinerMessage = `*GPU%s:*
  *Temperature:* %.0f C
  *Power usage:* %.0f W
  *Performance:* %.2f Sol/s
  *Efficiency:* %.2f Sol/W
  *Clocks:*
    %0.f Mhz core
    %.0f Mhz memory
`

func Create(jsonData interface{}) string {

	var message string

	switch value := jsonData.(type) {
	case models.EwbfResult:
		message = createEwbfMessage(value)
	case models.BMinerResult:
		message = createBMinerMessage(value)
	default:
		message = "unknown miner"
	}

	return message
}

func createEwbfMessage(ewbfResult models.EwbfResult) string {

	message := ""

	for _, result := range ewbfResult.Result {

		if len(message) > 0 {
			message = message + "\n"
		}

		message = message + fmt.Sprintf(baseEwbfMessage,
			result.GpuID, result.Name, result.Temperature, result.PowerUsage,
			result.SolsPerSecond, float32(result.SolsPerSecond)/float32(result.PowerUsage),
			result.AcceptedShares, result.RejectedShares,
		)
	}
	return makeEwbfHeader(ewbfResult) + "\n" + message
}

func makeEwbfHeader(ewbfResult models.EwbfResult) string {
	header := "*%s - Status*\n"

	if len(ewbfResult.Error) > 0 {
		header = header + "\nError: " + ewbfResult.Error
	}

	return fmt.Sprintf(header, os.Getenv("WORKER_ID"))
}

func createBMinerMessage(bminerResult models.BMinerResult) string {

	message := ""
	for _, result := range bminerResult.MinersList() {
		if len(message) > 0 {
			message = message + "\n"
		}

		message = message + fmt.Sprintf(baseBMinerMessage,
			result.Id, result.Device.Temperature, result.Device.Power,
			result.Solver.SolutionRate, (result.Solver.SolutionRate/result.Device.Power),
			result.Device.CoreClock, result.Device.MemoryClock,
		)
	}

	return fmt.Sprintf("*%s - Status*\n", os.Getenv("WORKER_ID")) + message + createBMinerFooter(bminerResult)
}

func createBMinerFooter(bminerResult models.BMinerResult) string {
	baseFooter := `
*Shares:*
  %d accepted, %d rejected`
	return fmt.Sprintf(baseFooter, bminerResult.Stratum.AcceptedShares, bminerResult.Stratum.RejectedShares)
}
