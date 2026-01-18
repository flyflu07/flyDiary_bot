package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/user/keyboards"
	"tg_bot/internal/utils/utilsUpdate"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		Text:        "Пошел нахуй",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.StartButton()},
	})
	if err != nil {
		log.Println("0xfdb4d -> ", err)
	}
}
