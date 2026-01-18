package internal

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func AllowedUpdates() bot.AllowedUpdates {
	allowedUpdates := bot.AllowedUpdates{
		models.AllowedUpdateMessage,
		models.AllowedUpdateEditedMessage,
		models.AllowedUpdateCallbackQuery,
	}
	return allowedUpdates
}
