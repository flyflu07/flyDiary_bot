package services

import (
	"tg_bot/internal/user/model"
)

func IsTimeZoneInMap(userID int64) bool {
	_, ok := model.TimeZone[userID]
	return ok
}
