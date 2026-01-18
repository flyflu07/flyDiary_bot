package storage

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	model2 "tg_bot/internal/user/model"
	"time"
)

func findMessagesByUserIDAndDate(userID int64, date time.Time) []model2.DiaryEntry {
	nextday := date.Add(24 * time.Hour)
	query := "SELECT d.id, d.text, d.recording_time + (interval '1 hour' * p.user_timezone) FROM diary d JOIN profile p ON d.id = p.id WHERE p.id = @id AND d.recording_time >= @date  AND d.recording_time < @nextday"
	args := &pgx.NamedArgs{
		"id":      userID,
		"date":    date,
		"nextday": nextday}

	rows, err := pool.Query(ctx, query, args)
	if err != nil {
		fmt.Println("Заплыв не удался")
	}

	var allMessagesByDate []model2.DiaryEntry

	for rows.Next() {
		var message model2.DiaryEntry
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
	queryExec := "insert into diary values(@id, @message, @user_time, @date);"
	_, err := pool.Exec(ctx, queryExec, args)
	if err != nil {
		fmt.Println("Карась уплыл")
	}
}

func getUniqueYearsOfDiaryEntries(UserID int64) []string {
	query := "select distinct extract(year from recording_time ) from diary where id=@id order by 1;"
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
	query := "select distinct extract(month from recording_time ) from diary where id=@id and extract(year from recording_time )=@year order by 1;"
	var args = &pgx.NamedArgs{"id": UserID, "year": model2.Date[UserID].Year}
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
	query := "select distinct extract(day from recording_time ) from diary where id=@id and extract(year from recording_time )=@year and extract(month from recording_time)=@month order by 1;"
	var args = &pgx.NamedArgs{"id": UserID, "year": model2.Date[UserID].Year, "month": model2.Date[UserID].Month}
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
