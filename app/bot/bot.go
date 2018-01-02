package bot

import (
	"fire-miner/app/api"
	"fire-miner/app/messageCreator"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	teleBot "gopkg.in/tucnak/telebot.v2"
)

func Init() {
	bot, err := teleBot.NewBot(teleBot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &teleBot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Println(err)
		return
	}

	bot.Handle("/status", func(message *teleBot.Message) {

		result, err := api.Status(os.Getenv("TARGET"))
		var replyMessage string
		if err == nil {
			replyMessage = messageCreator.Create(result)
		} else {
			log.Println("["+os.Getenv("WORKER_ID")+"]Error retrieving data from API, seems it is offline.", err)
			replyMessage = "[" + os.Getenv("WORKER_ID") + "] Error retrieving data from API, seems it is offline."
		}

		bot.Send(message.Sender, replyMessage, &teleBot.SendOptions{
			ParseMode: "Markdown",
		})
	})

	bot.Handle("/shutdown", func(message *teleBot.Message) {

		contents := strings.Split(message.Payload, " ")
		if len(contents) >= 2 && contents[0] == os.Getenv("WORKER_ID") && contents[1] == os.Getenv("SECRET") {
			bot.Send(message.Sender, "Shutdown started...")

			var command *exec.Cmd
			if runtime.GOOS == "linux" {
				command = exec.Command("sudo", "systemctl", "reboot")
			} else if runtime.GOOS == "windows" {
				command = exec.Command("shutdown", "-s", "-f")
			}

			command.Start()
		}
	})

	bot.Handle("/reboot", func(message *teleBot.Message) {

		contents := strings.Split(message.Payload, " ")
		if len(contents) >= 2 && contents[0] == os.Getenv("WORKER_ID") && contents[1] == os.Getenv("SECRET") {
			bot.Send(message.Sender, "Reboot started...")

			var command *exec.Cmd
			if runtime.GOOS == "linux" {
				log.Println("Rebooting now...")
				command = exec.Command("sudo", "systemctl", "reboot")
			} else if runtime.GOOS == "windows" {
				command = exec.Command("shutdown", "-r", "-f")
			}

			command.Start()
		}
	})

	bot.Handle("/startMiner", func(message *teleBot.Message) {

		command := exec.Command(os.Getenv("MINER_COMMAND"))
		command.Run()

		bot.Send(message.Sender, "Miner started...")
	})

	bot.Start()
}
