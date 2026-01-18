package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"tg_bot/config"
	"time"
)

var pool *pgxpool.Pool
var ctx context.Context

func init() {
	err := recover()
	if err != nil {
		log.Panicln("0x00000 -> ", err)
	}
	// Get the current time in the application's local timezone

	// Get the zone name (e.g., "CET", "UTC", "PDT") and offset in seconds

	connTemplate := "postgres://%s:%s@%s:%s/%s?options=-c%20TimeZone%3DAsia/Ekaterinburg"
	connStr := fmt.Sprintf(connTemplate, config.Cfg["Username"], config.Cfg["Password"], config.Cfg["Host"], config.Cfg["Port"], config.Cfg["DBName"])
	ctx = context.Background()

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Println("0x1b600 -> ", err)
	}

	// config values below are chosen by trial and error, feel free to modify them to suit your needs
	poolConfig.ConnConfig.ConnectTimeout = time.Second * 2
	poolConfig.ConnConfig.StatementCacheCapacity = 100
	poolConfig.MaxConnLifetime = 5 * time.Second
	poolConfig.MaxConnLifetimeJitter = 147 * time.Millisecond
	poolConfig.MaxConns = 30
	poolConfig.HealthCheckPeriod = 10 * time.Second

	pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}

}
