package commands

import (
	"log"
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMessage *tg.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	product, serviceErr := c.productService.Get(idx)
	if serviceErr != nil {
		log.Panic(serviceErr)
		return
	}
	msg := tg.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)
	_, sendErr := c.bot.Send(msg)
	if sendErr != nil {
		log.Panic(sendErr)
		return
	}
}
