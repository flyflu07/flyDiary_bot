package services

import (
	"tg_bot/internal/storage"
	model2 "tg_bot/internal/user/model"
	"time"
)

func GetMessage(userID int64) ([]model2.DiaryEntry, error) {
	text := model2.Date[userID].Year + "-" + model2.Date[userID].Month + "-" + model2.Date[userID].Day + ".000000"
	theTime, err := time.Parse("2006-1-2.000000", text)
	messages := storage.FindMessagesByDate(userID, theTime)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
