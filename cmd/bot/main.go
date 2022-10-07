package main

import (
	"log"
	"os"

	"github.com/artmalys/bot/internal/app/commands"
	"github.com/artmalys/bot/internal/service/product"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	token := os.Getenv("TOKEN")
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Bot %s started", bot.Self.UserName)
	u := tg.UpdateConfig{Timeout: 60}
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
