package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"tg_bot/internal/user/model"
	"tg_bot/internal/utils/utilsUpdate"
)

func LetsGoButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	model.ReplyChatIDConfirmationPassword[utilsUpdate.ExtractUserID(update)] = utilsUpdate.ExtractChatID(update)
	model.ReplyMessageIDConfirmationPassword[utilsUpdate.ExtractUserID(update)] = utilsUpdate.ExtractMessageID(update)
	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    utilsUpdate.ExtractChatID(update),
		MessageID: utilsUpdate.ExtractMessageID(update),
		Text:      "Придумайте хороший пароль:",
	})
	if err != nil {
		log.Println("0x7224d -> ", err)
	}

}
