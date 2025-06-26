package worker

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/ArtemSafin/osint-bot/internal/leaklookup"
	"github.com/ArtemSafin/osint-bot/internal/queue"
)

func Start(bot *tgbotapi.BotAPI) {
	go func(bot *tgbotapi.BotAPI) {
		for {
			msg, err := queue.Pop()
			if err != nil {
				continue
			}

			var task queue.Task
			err = json.Unmarshal([]byte(msg), &task)
			if err != nil {
				log.Printf("❌ Ошибка декодирования задачи: %v", err)
				continue
			}

			log.Printf("👷 Обрабатываю email: %s для chat_id: %d", task.Email, task.ChatID)

			leaks, err := leaklookup.CheckEmail(task.Email)
			if err != nil {
				log.Printf("❌ LeakLookup error: %v", err)
				continue
			}

			if len(leaks) == 0 {
				log.Printf("✅ %s — утечек не найдено", task.Email)
				if _, err := bot.Send(tgbotapi.NewMessage(task.ChatID, fmt.Sprintf("✅ %s — утечек не найдено", task.Email))); err != nil {
					log.Printf("❌ Ошибка отправки сообщения: %v", err)
				}
			} else {
				log.Printf("💀 %s — найдено утечек: %d", task.Email, len(leaks))

				msg := fmt.Sprintf("💀 %s — найдено утечек: %d\n", task.Email, len(leaks))
				for domain, details := range leaks {
					msg += fmt.Sprintf("🔎 %s — %d записей\n", domain, len(details))
				}

				if _, err := bot.Send(tgbotapi.NewMessage(task.ChatID, msg)); err != nil {
					log.Printf("❌ Ошибка отправки сообщения: %v", err)
				}
			}
		}
	}(bot)
}
