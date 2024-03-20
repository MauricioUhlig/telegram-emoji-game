package main

import (
	"log"

	"github.com/MauricioUhlig/telegram-emoji-game/models"
	"github.com/MauricioUhlig/telegram-emoji-game/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	util.FetchYAML()
	bot, err := tgbotapi.NewBotAPI(util.GetTelegramKey())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "Send /new to recive your emoji list to play. Or you can send /hi or /status as well"
		case "hi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "new":
			msg.Text = models.GetEmojis()
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
