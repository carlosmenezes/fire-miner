package bot

import (
	teleBot "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

func Init() {
	log.Println("TESTE")

	bot, err := teleBot.NewBot(teleBot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &teleBot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/hello", func(m *teleBot.Message) {
		bot.Send(m.Sender, "*hello world* _bla_", &teleBot.SendOptions{
			ParseMode: "Markdown",
		})
	})

	bot.Start()
}
