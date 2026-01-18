package handlers

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/user/keyboards"
	model2 "tg_bot/internal/user/model"
	services2 "tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
)

func MenuFindDateWithYearsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if services2.IsPasswordInMap(utilsUpdate.ExtractUserID(update)) {

		fmt.Println(model2.Page[utilsUpdate.ExtractUserID(update)])

		_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:      utilsUpdate.ExtractChatID(update),
			MessageID:   utilsUpdate.ExtractMessageID(update),
			Text:        "Выберите год по которой хотите найти запись",
			ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.FindYears(utilsUpdate.ExtractUserID(update))},
		})
		if err != nil {
			log.Println("0xab763 -> ", err)
		}
	} else {
		_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: utilsUpdate.ExtractCallbackQueryID(update),
			Text:            "Дневнин заблокирован, для доступа к записям введите пароль",
			ShowAlert:       true,
		})
		if err != nil {
			log.Println("0x065f3 -> ", err)
		}
	}
}

func MenuFindDateWithMonthsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		MessageID:   utilsUpdate.ExtractMessageID(update),
		Text:        "Выберите месяц по которой хотите найти запись",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.FindMonths(utilsUpdate.ExtractUserID(update))},
	})
	if err != nil {
		log.Println("0x065f3 -> ", err)
	}

}

func MenuFindDateWithDaysHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println(model2.Page[utilsUpdate.ExtractUserID(update)])

	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      utilsUpdate.ExtractChatID(update),
		MessageID:   utilsUpdate.ExtractMessageID(update),
		Text:        "Выберите день по которой хотите найти запись",
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.FindDays(utilsUpdate.ExtractUserID(update))},
	})
	if err != nil {
		log.Println("0xdd268 -> ", err)
	}
}

// 1
func YearsButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if model2.Date[utilsUpdate.ExtractUserID(update)] == nil {
		model2.Date[utilsUpdate.ExtractUserID(update)] = &model2.DateForBD{}
	}

	model2.Date[utilsUpdate.ExtractUserID(update)].Year = update.CallbackQuery.Data[6:]
	MenuFindDateWithMonthsHandler(ctx, b, update)
}

func MonthsButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	model2.Date[utilsUpdate.ExtractUserID(update)].Month = update.CallbackQuery.Data[7:]
	MenuFindDateWithDaysHandler(ctx, b, update)
}

func DaysButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	model2.Date[utilsUpdate.ExtractUserID(update)].Day = update.CallbackQuery.Data[5:]

	FinalMenuFindHandler(ctx, b, update)
}

func FinalMenuFindHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	messages, err := services2.GetMessage(utilsUpdate.ExtractUserID(update))
	if err != nil {
		log.Println("0xac733 -> ", err)
		return
	}

	for i, msg := range messages {
		var message *models.Message
		if i == len(messages)-1 {
			decryptedText := services2.CryptoDecrypt(msg.Message, model2.InfoAbUser[utilsUpdate.ExtractUserID(update)].Password)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      utilsUpdate.ExtractChatID(update),
				Text:        "*" + msg.Time.Format("15:04:05") + "*" + "\n" + decryptedText,
				ParseMode:   models.ParseModeMarkdownV1,
				ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.ClearAll()},
			})
			if err != nil {
				log.Println("0x35212 -> ", err)
			}
		} else {
			decryptedText := services2.CryptoDecrypt(msg.Message, model2.InfoAbUser[utilsUpdate.ExtractUserID(update)].Password)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    utilsUpdate.ExtractChatID(update),
				Text:      "*" + msg.Time.Format("15:04:05") + "*" + "\n" + decryptedText,
				ParseMode: models.ParseModeMarkdownV1,
			})
			if err != nil {
				log.Println("0x35213 -> ", err)
			}
		}
		model2.GetMessageForClear[utilsUpdate.ExtractUserID(update)] = append(model2.GetMessageForClear[utilsUpdate.ExtractUserID(update)], message.ID)

	}

}

func CloseFDButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch update.CallbackQuery.Data[8:] {
	case "years":
		MenuDiaryButtonHandler(ctx, b, update)
	case "month":
		MenuFindDateWithYearsHandler(ctx, b, update)
	case "days":
		MenuFindDateWithMonthsHandler(ctx, b, update)

	}
}

func ClearAllHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := utilsUpdate.ExtractUserID(update)
	chatID := utilsUpdate.ExtractChatID(update)

	for _, msgID := range model2.GetMessageForClear[userID] {
		_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatID,
			MessageID: msgID,
		})
		if err != nil {
			log.Println("0x8c66c -> ", err)
		}
	}

	delete(model2.GetMessageForClear, userID)

	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: utilsUpdate.ExtractCallbackQueryID(update),
		Text:            "Успешно",
	})
	if err != nil {
		log.Println("0x8c67c -> ", err)
	}
}
