package services

import (
	"tg_bot/internal/user/model"
)

func IsPasswordInMap(userID int64) bool {
	_, ok := model.InfoAbUser[userID]
	return ok
}
