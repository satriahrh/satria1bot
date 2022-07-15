package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/satriahrh/satria1bot/usecase"
)

func main() {
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	fmt.Println("Starting bot...")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			replyText := update.Message.Text
			command := update.Message.Command()
			switch command {
			case "add":
				args := update.Message.CommandArguments()

				re := regexp.MustCompile(`(\d+)\+(\d+)`)
				matches := re.FindStringSubmatch(args)
				a, _ := strconv.Atoi(matches[1])
				b, _ := strconv.Atoi(matches[2])

				c := usecase.Calculate(a, b)
				replyText = strconv.Itoa(c)
			}
			log.Printf("[%s] %s", update.Message.From.UserName, replyText)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
