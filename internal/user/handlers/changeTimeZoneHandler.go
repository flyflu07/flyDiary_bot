package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/storage"
	"tg_bot/internal/user/keyboards"
	"tg_bot/internal/user/model"
	"tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
)

func MenuTineZoneButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	model.MessageIDTimeZoneMenu[utilsUpdate.ExtractUserID(update)] = utilsUpdate.ExtractMessageID(update)
	if services.IsTimeZoneInMap(utilsUpdate.ExtractUserID(update)) {
		model.TimeZone[utilsUpdate.ExtractUserID(update)] = 0
	}
	timezone, _ := services.TimeWithTimeZone(utilsUpdate.ExtractUserID(update))
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		Text:        "У вас сейчас " + "*" + timezone + "*?",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.ChangeTimezoneButtons()},
		ParseMode:   models.ParseModeMarkdownV1,
	})
	if err != nil {
		log.Println("0x23d11 -> ", err)
	}
}

func MenuTimeZoneHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	timeZone, _ := services.TimeWithTimeZone(utilsUpdate.ExtractUserID(update))
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		MessageID:   utilsUpdate.ExtractMessageID(update),
		Text:        "У вас сейчас " + "*" + timeZone + "*?",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.ChangeTimezoneButtons()},
		ParseMode:   models.ParseModeMarkdownV1,
	})
	if err != nil {
		log.Println("0x8930a -> ", err)
	}
}

func RightTimeZoneButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	model.TimeZone[utilsUpdate.ExtractUserID(update)]++
	MenuTimeZoneHandler(ctx, b, update)
}

func LeftTimeZoneButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	model.TimeZone[utilsUpdate.ExtractUserID(update)]--
	MenuTimeZoneHandler(ctx, b, update)
}

func SaveTimeZoneButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Сохраняем в бд изменение тайм зоны по userID
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: utilsUpdate.ExtractCallbackQueryID(update),
		Text:            "Ваше время успешно сохранено",
		ShowAlert:       true,
	})
	if err != nil {
		log.Println("0x9202e -> ", err)
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    utilsUpdate.ExtractChatID(update),
		MessageID: model.MessageIDTimeZoneMenu[utilsUpdate.ExtractUserID(update)],
	})
	if err != nil {
		log.Println("0x6e2ca -> ", err)
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    utilsUpdate.ExtractChatID(update),
		MessageID: utilsUpdate.ExtractMessageID(update),
	})
	if err != nil {
		log.Println("0x0fc9f -> ", err)
	}
	delete(model.MessageIDTimeZoneMenu, utilsUpdate.ExtractUserID(update))
	storage.SaveTimeZoneInDB(utilsUpdate.ExtractUserID(update), model.TimeZone[utilsUpdate.ExtractUserID(update)])
	delete(model.TimeZone, utilsUpdate.ExtractUserID(update))
	MenuDiaryButtonHandler(ctx, b, update)
}
