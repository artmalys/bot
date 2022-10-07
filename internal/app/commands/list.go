package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) List(inputMessage *tg.Message) {
	var keyboardButtons []tg.InlineKeyboardButton
	for _, p := range c.productService.List() {
		keyboardButtons = append(keyboardButtons,
			tg.NewInlineKeyboardButtonData(p.Title, p.TextPrice()))
	}
	msg := tg.NewMessage(
		inputMessage.Chat.ID,
		"Hello, please choose a product!")

	msg.ReplyMarkup = tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(keyboardButtons...))

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Panic(err)
		return
	}
}
