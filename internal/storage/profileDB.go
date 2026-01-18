package storage

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func passwordCheck(id int64) bool {
	query := "select exists(select 1 from profile where id=@id and ispassword)"
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
	query := "insert into profile(id, ispassword, passmd5) values (@id, true, @password);"
	args := &pgx.NamedArgs{"id": id, "password": password}
	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		log.Printf("Create profile failed: %v", err)
	}
}

func getPassword(id int64) string {
	query := "select passmd5 from profile p  where p.id = @id;"
	var args = &pgx.NamedArgs{"id": id}
	var password string
	err := pool.QueryRow(ctx, query, args).Scan(&password)
	if err != nil {
		log.Printf("Scan failed in compare passwords: %v", err)
	}
	return password
}

func saveTimeZoneInDB(id int64, timezone int) {
	query := "update profile set timezone=@timezone where id=@id"
	args := &pgx.NamedArgs{"id": id, "timezone": timezone}
	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		log.Printf("Save time zone failed: %v", err)
	}
}

func getTimeZone(id int64) string {
	query := "select timezone from profile where id=@id;"
	var args = &pgx.NamedArgs{"id": id}
	var timezone string
	err := pool.QueryRow(ctx, query, args).Scan(&timezone)
	if err != nil {
		log.Printf("Scan failed in get timezone: %v", err)
	}
	return timezone
}
