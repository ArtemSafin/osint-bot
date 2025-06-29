# OSINT Bot

Telegram-бот для выполнения OSINT-задач по email/IP/никнейма с использованием Go, Redis, очередей и API.

## Используемый стек

- Go
- Redis
- PostgreSQL (в будущем)
- Docker + Docker Compose
- Telegram Bot API
- Leaklookup API

## Установка

1. Клонировать репозиторий
2. Создать `.env` на основе `.env.example`
3. Запустить `docker-compose up`

## Структура проекта

- `cmd/bot` — точка входа Telegram-бота
- `internal/leaklookup` — работа с API "leak look up"
- `internal/queue` — логика очередей Redis
- `internal/worker` — воркеры, обрабатывающие задачи
