package main

import (
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/ArtemSafin/osint-bot/internal/queue"
	"github.com/ArtemSafin/osint-bot/internal/worker"
)

func main() {

	queue.InitRedis()

	_ = godotenv.Load()

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	worker.Start(bot)

	bot.Debug = true
	log.Printf("Bot authorized on account %s", &bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() && update.Message.Command() == "check" {
			email := update.Message.CommandArguments()
			if email == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Введите email после команды"))
				continue
			}

			task := queue.Task{
				ChatID: update.Message.Chat.ID,
				Email:  email,
			}

			data, err := json.Marshal(task)
			if err != nil {
				log.Printf("❌ Ошибка сериализации задачи: %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при формировании задачи"))
				return
			}

			err = queue.Push(string(data))
			if err != nil {
				log.Printf("❌ Ошибка при постановке в очередь: %v", err)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при постановке в очередь"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "✅ Email поставлен в очередь. Ожидайте результат."))
			}

		}
	}
}
