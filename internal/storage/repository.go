package storage

import (
	"tg_bot/internal/user/model"
	"time"
)

func FindMessagesByDate(userID int64, date time.Time) []model.DiaryEntry {
	return findMessagesByUserIDAndDate(userID, date)
}

func SaveDiary(userID int64, message string, user_time time.Time, date time.Time) {
	saveDiary(userID, message, user_time, date)
}
func PasswordCheck(userID int64) bool {
	return passwordCheck(userID)
}

func CreateProfile(userID int64, password string) {
	createProfile(userID, password)
}

func GetPassword(userID int64) string {
	return getPassword(userID)
}

func SaveTimeZoneInDB(userID int64, timezone int) {
	saveTimeZoneInDB(userID, timezone)
}

func GetTimeZone(userID int64) string {
	return getTimeZone(userID)
}

func GetUniqueYearsOfDiaryEntries(userID int64) []string {
	return getUniqueYearsOfDiaryEntries(userID)

}

func GetUniqueMonthsOfDiaryEntries(userID int64) []string {
	return getUniqueMonthsOfDiaryEntries(userID)

}

func GetUniqueDaysOfDiaryEntries(userID int64) []string {
	return getUniqueDaysOfDiaryEntries(userID)

}
