package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/storage"
	"tg_bot/internal/user/model"
	"tg_bot/internal/user/services"
	"tg_bot/internal/utils/utilsUpdate"
)

func YesButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	password := model.InfoAbUser[utilsUpdate.ExtractUserID(update)].Password
	password = services.Makemd5(password)
	storage.CreateProfile(utilsUpdate.ExtractUserID(update), password)
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    utilsUpdate.ExtractMessageID(update),
		MessageID: utilsUpdate.ExtractMessageID(update),
		Text: "Пароль успешно изменён! " +
			"\n *Напоминаем вам,что пароль больше изменить нельзя!* " +
			"\n (Для повышения безопастности вы можете удалить сообщение " +
			"с паролем, Владелец бота и сам бот не знает пароль. *Пароль храниться только у вас*)",
		ParseMode: models.ParseModeMarkdownV1,
	})
	if err != nil {
		log.Println("0x83f35 -> ", err)
	}

}

func NoButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    utilsUpdate.ExtractChatID(update),
		MessageID: utilsUpdate.ExtractMessageID(update),
		Text:      "Придумайте пароль",
	})
	if err != nil {
		log.Println("0xcd39d -> ", err)
	}
	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    utilsUpdate.ExtractChatID(update),
		MessageID: model.SaveMessageIDToUser[utilsUpdate.ExtractUserID(update)],
	})
	if err != nil {
		log.Println("0xc848e -> ", err)
	}
}
