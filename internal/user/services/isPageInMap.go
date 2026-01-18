package services

import (
	"tg_bot/internal/user/model"
)

func IsPageInMap(userID int64) bool {
	_, ok := model.Date[userID]
	return ok
}
