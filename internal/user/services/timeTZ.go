package services

import (
	"fmt"
	"tg_bot/internal/user/model"
	"time"
)

func TimeWithTimeZone(userID int64) (string, string) {
	utcNow := time.Now().UTC()
	offset := model.TimeZone[userID]
	formatLayout := "15:04"
	formatLayoutUTC := "-0700"
	offsetSeconds := offset * 3600
	loc := time.FixedZone("_", offsetSeconds)
	timeInZone := utcNow.In(loc)
	return fmt.Sprintf("%s\n", timeInZone.Format(formatLayout)), timeInZone.Format(formatLayoutUTC)

}
