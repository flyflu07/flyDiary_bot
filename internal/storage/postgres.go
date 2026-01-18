package storage

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	q "tg_bot/internal/storage/query"
	"tg_bot/internal/user/model"
	"time"
)

func findMessagesByUserIDAndDate(userID int64, date time.Time) []model.DiaryEntry {
	nextday := date.Add(24 * time.Hour)
	query := q.FindMessageQuery
	args := &pgx.NamedArgs{
		"id":      userID,
		"date":    date,
		"nextday": nextday}

	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		fmt.Println("Заплыв не удался")
	}

	var allMessagesByDate []model.DiaryEntry

	for rows.Next() {
		var message model.DiaryEntry
		err = rows.Scan(&message.Id, &message.Message, &message.Time)
		if err != nil {
			fmt.Println("Закладку спиздили или не нашли")
		}
		allMessagesByDate = append(allMessagesByDate, message)
	}

	return allMessagesByDate
}

func saveDiary(id int64, message string, user_time time.Time, date time.Time) {

	args := &pgx.NamedArgs{
		"id":        id,
		"message":   message,
		"user_time": user_time,
		"date":      date,
	}
	queryExec := q.SaveDiaryQuery
	_, err := pool.Exec(ctx, queryExec, args)
	if err != nil {
		fmt.Println("Карась уплыл")
	}
}

func getUniqueYearsOfDiaryEntries(UserID int64) []string {
	query := q.GetUniqueYearsQuery
	var args = &pgx.NamedArgs{"id": UserID}
	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		log.Println("0x30220 -> ", err)
	}
	var years []string
	var year int
	for rows.Next() {
		err = rows.Scan(&year)
		if err != nil {
			fmt.Println("Закладку спиздили или не нашли")
		}
		years = append(years, strconv.Itoa(year))
	}

	return years
}

func getUniqueMonthsOfDiaryEntries(UserID int64) []string {
	query := q.GetUniqueMonthsQuery
	var args = &pgx.NamedArgs{"id": UserID, "year": model.Date[UserID].Year}
	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		log.Println("0x11111 -> ", err)
	}
	var months []string
	var month int
	for rows.Next() {
		err = rows.Scan(&month)
		if err != nil {
			fmt.Println("Закладку спиздили или не нашли")
		}
		months = append(months, strconv.Itoa(month))
	}
	return months
}

func getUniqueDaysOfDiaryEntries(UserID int64) []string {
	query := q.GetUniqueDaysQuery
	var args = &pgx.NamedArgs{"id": UserID, "year": model.Date[UserID].Year, "month": model.Date[UserID].Month}
	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		log.Println("0x714e8 -> ", err)
	}
	var days []string
	var day int
	for rows.Next() {
		err = rows.Scan(&day)
		if err != nil {
			fmt.Println("Закладку спиздили или не нашли")
		}
		days = append(days, strconv.Itoa(day))
	}
	return days
}
