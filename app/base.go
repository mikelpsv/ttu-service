package app

import (
	"github.com/joho/godotenv"
	"os"
)

type appCfg struct {
	AppAddr    string
	AppPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

var AppVersion string
var ApiVersion string

var Cfg appCfg

func (cfg *appCfg) ReadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg.AppAddr = os.Getenv("APP_ADDR")
	cfg.AppPort = os.Getenv("APP_PORT")

	cfg.DbHost = os.Getenv("POSTGRES_HOST")
	cfg.DbPort = os.Getenv("POSTGRES_PORT")
	cfg.DbName = os.Getenv("POSTGRES_DB")
	cfg.DbUser = os.Getenv("POSTGRES_USER")
	cfg.DbPassword = os.Getenv("POSTGRES_PASSWORD")

	return nil
}
