package services

import (
	"tg_bot/internal/storage"
	model "tg_bot/internal/user/model"
	"time"
)

func GetMessage(userID int64) ([]model.DiaryEntry, error) {
	text := model.Date[userID].Year + "-" + model.Date[userID].Month + "-" + model.Date[userID].Day + ".000000"
	theTime, err := time.Parse("2006-1-2.000000", text)
	messages := storage.FindMessagesByDate(userID, theTime)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
