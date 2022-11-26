package main

type Configuration struct {
	BOT_TOKEN string
	BOT_DEBUG bool
}

type Questions struct {
	Question string
	Answer   string
}

type DeleteMessageConfig struct {
	ChatID    int64
	MessageID int
}
