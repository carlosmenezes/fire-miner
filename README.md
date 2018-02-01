# FireMiner

A simple telegram bot for monitoring mining rigs.

## Commands Available

  - `/status` - Shows informations about the rig GPU's (temperatures, hashrate, power usage, power efficiency, etc.).
  - `/reboot SECRET` - Reboots the rig if the secret is correct.
  - `/shutdown SECRET` - Shutdown the rig if the secret is correct.
  - `/startMiner` - Start the mining software (useful after a reboot).

## Environment Variables

  In order to this bot work corretly some enviroment variables must be set:

  - __WORKER_ID:__ A identifier for your rig, example: `MyAwesomeRig`
  - __TOKEN:__ Your telegram api secret token, example: `12122334:SOME_HASH`
  - __TARGET:__ The ip of the rig to be monitored, example: `http://localhost:42000/getstat`
  - __MINER_COMMAND:__ The command used to start the miner software, example: `miner.exe --server zec-eu1.nanopool.org --user YOUR_WALLET_ADDRESS.YOUR_WORKER_NAME/YOUR_EMAIL --pass z --port 6666
`
  - __SECRET:__ A password required to use the commands `/shutdown` and `/restart`
  - __MINER:__ The you actually using, example: (`ewbf` or `bminer`)

## Miner Softwares Supported

  At this moment only [EWBF Cuda Miner](https://github.com/nanopool/ewbf-miner) and [BMiner](https://www.bminer.me), but contributions are welcome!
