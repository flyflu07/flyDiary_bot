package storage

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	q "tg_bot/internal/storage/query"
)

func passwordCheck(id int64) bool {
	query := q.PasswordCheckQuery
	var args = &pgx.NamedArgs{"id": id}
	var isPassword bool
	err := pool.QueryRow(ctx, query, args).Scan(&isPassword)
	if err != nil {
		log.Printf("Scan failed: %v", err)
		return false
	}

	return isPassword
}

func createProfile(id int64, password string) {
	fmt.Println(password)
	query := q.CreateProfileQuery
	args := &pgx.NamedArgs{"id": id, "password": password}
	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		log.Printf("Create profile failed: %v", err)
	}
}

func getPassword(id int64) string {
	query := q.GetPasswordQuery
	var args = &pgx.NamedArgs{"id": id}
	var password string
	err := pool.QueryRow(ctx, query, args).Scan(&password)
	if err != nil {
		log.Printf("Scan failed in compare passwords: %v", err)
	}
	return password
}

func saveTimeZoneInDB(id int64, timezone int) {
	query := q.SaveTimeZone
	args := &pgx.NamedArgs{"id": id, "timezone": timezone}
	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		log.Printf("Save time zone failed: %v", err)
	}
}

func getTimeZone(id int64) string {
	query := q.GetTimeZoneQuery
	var args = &pgx.NamedArgs{"id": id}
	var timezone string
	err := pool.QueryRow(ctx, query, args).Scan(&timezone)
	if err != nil {
		log.Printf("Scan failed in get timezone: %v", err)
	}
	return timezone
}
