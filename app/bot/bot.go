package bot

import (
	"fire-miner/app/api"
	"fire-miner/app/messageCreator"
	"log"
	"os"
	"os/exec"
	"time"

	teleBot "gopkg.in/tucnak/telebot.v2"
)

func Init() {
	bot, err := teleBot.NewBot(teleBot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &teleBot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/status", func(message *teleBot.Message) {

		result, _ := api.Status(os.Getenv("TARGET"))
		replyMessage := messageCreator.Create(result)

		bot.Send(message.Sender, replyMessage, &teleBot.SendOptions{
			ParseMode: "Markdown",
		})
	})

	bot.Handle("/shutdown", func(message *teleBot.Message) {

		if message.Payload == os.Getenv("SECRET") {
			bot.Send(message.Sender, "Shutdown started...")

			command := exec.Command("shutdown", "-s", "-f")
			command.Start()
		}
	})

	bot.Handle("/reboot", func(message *teleBot.Message) {

		if message.Payload == os.Getenv("SECRET") {
			bot.Send(message.Sender, "Reboot started...")

			command := exec.Command("shutdown", "-r", "-f")
			command.Start()
		}
	})

	bot.Handle("/startMiner", func(message *teleBot.Message) {

		command := exec.Command(os.Getenv("MINER_COMMAND"))
		command.Start()

		bot.Send(message.Sender, "Miner started...")
	})

	bot.Start()
}
