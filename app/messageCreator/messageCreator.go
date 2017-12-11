package messageCreator

import (
	"fire-miner/app/models"
	"fmt"
)

const baseMessage = `*Status*

*GPU%d:* %s
  *Temperature:* %dC
  *Power usage:* %dW
  *Performance:* %dSol/s
  *Efficiency:* %.2f Sol/W
  *Shares:*
    %d accepted, %d rejected
`

func Create(ewbfResult models.EwbfResult) string {

	result := ewbfResult.Result[0]
	return fmt.Sprintf(baseMessage,
		result.GpuID, result.Name, result.Temperature, result.PowerUsage,
		result.SolsPerSecond, float32(result.SolsPerSecond)/float32(result.PowerUsage),
		result.AcceptedShares, result.RejectedShares,
	)
}
