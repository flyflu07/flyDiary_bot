package services

import (
	"github.com/go-telegram/bot/models"
	internal "tg_bot/internal/user/keyboards"
	"tg_bot/internal/user/model"
)

func IsUnlockDiary(userID int64) [][]models.InlineKeyboardButton {
	_, ok := model.InfoAbUser[userID]
	if ok {
		return internal.MenuWithUnlock()
	} else {
		return internal.MenuWithLock()
	}
}
