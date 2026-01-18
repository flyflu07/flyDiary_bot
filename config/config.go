package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	EntranceToken string `env:"DIARY_TG_ENTRANCETOKEN,required,notEmpty"`
	PostgresHost  string `env:"DIARY_PG_HOST,required,notEmpty"`
	PostgresPort  uint16 `env:"DIARY_PG_PORT,required,notEmpty"`
	PostgresDb    string `env:"DIARY_PG_DBNAME,required,notEmpty"`
	PostgresUser  string `env:"DIARY_PG_USERNAME,required,notEmpty"`
	PostgresPwd   string `env:"DIARY_PG_PASSWORD,required,notEmpty"`
	PostgresSSLM  string `env:"DIARY_PG_SSLMODE,required,notEmpty"`
}

var configInstance Config

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = env.Parse(&configInstance)
	if err != nil {
		panic(err)
	}
}

func GetEntranceToken() string {
	return configInstance.EntranceToken
}
func GetPostgresHost() string {
	return configInstance.PostgresHost
}
func GetPostgresPort() uint16 {
	return configInstance.PostgresPort
}
func GetPostgresDb() string {
	return configInstance.PostgresDb
}
func GetPostgresUser() string {
	return configInstance.PostgresUser
}
func GetPostgresPwd() string {
	return configInstance.PostgresPwd
}
func GetPostgresSSLM() string {
	return configInstance.PostgresSSLM
}
