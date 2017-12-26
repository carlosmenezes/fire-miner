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

func Create(ewbfResult models.EwbfResult) string {

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
	return makeHeader(ewbfResult) + "\n" + message
}

func makeHeader(ewbfResult models.EwbfResult) string {
	header := "*%s - Status*\n"

	if len(ewbfResult.Error) > 0 {
		header = header + "\nError: " + ewbfResult.Error
	}

	return fmt.Sprintf(header, os.Getenv("WORKER_ID"))
}
