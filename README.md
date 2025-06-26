# OSINT Bot 🕵️‍♂️

Telegram-бот для выполнения OSINT-задач по email/IP/имени с использованием Go, Redis, очередей и API.

## Используемый стек

- Go
- Redis
- PostgreSQL (в будущем)
- Docker + Docker Compose
- Telegram Bot API
- HaveIBeenPwned + Epieos

## Установка

1. Клонируй репозиторий
2. Создай `.env` на основе `.env.example`
3. Запусти `docker-compose up`

## Структура проекта

- `cmd/bot` — точка входа Telegram-бота
- `internal/epieos` — работа с Epieos
- `internal/hibp` — работа с HIBP API
- `internal/queue` — логика очередей Redis
- `internal/worker` — воркеры, обрабатывающие задачи
