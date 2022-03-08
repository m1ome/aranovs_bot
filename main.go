package main

import (
	"flag"
	"github.com/m1ome/aranovs_bot/matcher"
	"gopkg.in/telebot.v3"
	"log"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "token", "", "telegram token")
	flag.Parse()
}

func main() {
	preferences := telebot.Settings{
		Token: token,
	}

	bot, err := telebot.NewBot(preferences)
	if err != nil {
		log.Fatalf("error starting bot: %v", err)
	}

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		if c.Chat() == nil {
			return nil
		}

		if sticker := matcher.Matches(c.Text()); sticker != nil {
			options := &telebot.SendOptions{
				ReplyTo: c.Message(),
			}

			if _, err := bot.Send(c.Chat(), sticker, options); err != nil {
				log.Fatalf("error sending sticker: %v", err)
			}
		}

		return nil
	})

	log.Println("starting telegram bot")
	bot.Start()
}
