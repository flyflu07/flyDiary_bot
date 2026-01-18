package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/user/keyboards"
	"tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
)

func MenuDiaryCommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if services.IsPasswordInMap(utilsUpdate.ExtractChatID(update)) {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      utilsUpdate.ExtractChatID(update),
			Text:        "Тут менюшечка",
			ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MenuWithLock()},
		})
		if err != nil {
			log.Println("0x5a4c3 -> ", err)
		}
	} else {

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      utilsUpdate.ExtractChatID(update),
			Text:        "Тут менюшечка",
			ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MenuWithUnlock()},
		})
		if err != nil {
			log.Println("0x9e1f2 -> ", err)
		}
	}

}

func MenuDiaryButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		MessageID:   utilsUpdate.ExtractMessageID(update),
		Text:        "Тут менюшечка",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MenuWithLock()},
	})
	if err != nil {
		log.Println("0x5a4b3 -> ", err)
	}
}
