package messageCreator

import (
	"fire-miner/app/models"
	"fmt"
	"os"
)

const baseMessage = `*GPU%d:* %s
  *Temperature:* %d C
  *Power usage:* %d W
  *Performance:* %d Sol/s
  *Efficiency:* %.2f Sol/W
  *Shares:*
    %d accepted, %d rejected
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

		message = message + fmt.Sprintf(baseMessage,
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
	return message
}
