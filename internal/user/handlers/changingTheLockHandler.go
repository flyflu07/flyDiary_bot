package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/user/keyboards"
	"tg_bot/internal/user/model"
	"tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
)

func LockDiaryButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if services.IsPasswordInMap(utilsUpdate.ExtractUserID(update)) {
		delete(model.InfoAbUser, utilsUpdate.ExtractUserID(update))
	}
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		MessageID:   utilsUpdate.ExtractMessageID(update),
		Text:        "Тут менюшечка",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.MenuWithUnlock()},
	})
	if err != nil {
		log.Println("0xde409 -> ", err)
	}
	_, err = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: utilsUpdate.ExtractCallbackQueryID(update),
		Text:            "Дневник заблокирован",
	})
	if err != nil {
		log.Println("0x97450 -> ", err)
	}
}

func LockDiaryCommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if services.IsPasswordInMap(utilsUpdate.ExtractUserID(update)) {
		delete(model.InfoAbUser, utilsUpdate.ExtractUserID(update))

	}
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: utilsUpdate.ExtractChatID(update),
		Text:   "Дневник заблокирован",
	})
	if err != nil {
		log.Println("0x4f43c -> ", err)
	}
}
