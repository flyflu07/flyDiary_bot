package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/storage"
	"tg_bot/internal/user/keyboards"
	model2 "tg_bot/internal/user/model"
	services2 "tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
	"time"
)

func Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		if update.Message.Text != "" {

			if !storage.PasswordCheck(utilsUpdate.ExtractUserID(update)) {
				model2.SaveMessageIDToUser[utilsUpdate.ExtractUserID(update)] = utilsUpdate.ExtractMessageID(update)
				model2.InfoAbUser[utilsUpdate.ExtractUserID(update)] = &model2.User{Password: update.Message.Text}
				text := model2.InfoAbUser[utilsUpdate.ExtractUserID(update)].Password
				_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
					ChatID:      model2.ReplyChatIDConfirmationPassword[utilsUpdate.ExtractChatID(update)],
					MessageID:   model2.ReplyMessageIDConfirmationPassword[utilsUpdate.ExtractChatID(update)],
					Text:        "Вы точно хотите сохранить " + "*" + text + "*" + " как пароль? (Пароль изменить позже нельзя)",
					ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: keyboards.ConfirmationPasswordButtons()},
					ParseMode:   models.ParseModeMarkdownV1,
				})
				if err != nil {
					log.Println("0x22202 -> ", err)
				}
			} else {
				switch {

				case !services2.IsPasswordInMap(utilsUpdate.ExtractUserID(update)):
					if services2.ComparePasswords(update.Message.Text, storage.GetPassword(utilsUpdate.ExtractUserID(update))) {
						logoutTime := 5 * time.Minute
						model2.InfoAbUser[utilsUpdate.ExtractUserID(update)] = &model2.User{Password: update.Message.Text, TimeOfStartSession: time.Now(), TimeOfAutoLogOut: logoutTime}
						MenuDiaryCommandHandler(ctx, b, update)
						go services2.CheckForSessionTime(utilsUpdate.ExtractUserID(update))
					} else {
						_, err := b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID: utilsUpdate.ExtractChatID(update),
							Text:   "Пароль неверный, попробуйте снова",
						})
						if err != nil {
							log.Println("0xca945 -> ", err)
						}
						_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
							ChatID:    utilsUpdate.ExtractChatID(update),
							MessageID: utilsUpdate.ExtractMessageID(update),
						})
						if err != nil {
							log.Println("0xb06bd -> ", err)
						}

					}

				default:

					cryptoText := services2.CryptoEncrypt(update.Message.Text, model2.InfoAbUser[utilsUpdate.ExtractUserID(update)].Password)
					storage.SaveDiary(utilsUpdate.ExtractUserID(update), cryptoText, time.Now().UTC(), time.Now().UTC())
					_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
						ChatID:    utilsUpdate.ExtractChatID(update),
						MessageID: utilsUpdate.ExtractMessageID(update),
					})
					if err != nil {
						log.Println("0x2b84d -> ", err)
					}
				}
			}
		}
	}
}
