package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvType struct {
	AppPort        string
	MainDB         string
	MainDBUser     string
	MainDBPassword string
	MainDBPort     string
	MainDBHost     string
	MainDBSSLMode  string
}

var Env EnvType

func InitEnv() {
	godotenv.Load()
	Env.AppPort = os.Getenv("APP_PORT")
	Env.MainDB = os.Getenv("MAIN_DB")
	Env.MainDBUser = os.Getenv("MAIN_DB_USER")
	Env.MainDBPassword = os.Getenv("MAIN_DB_PASSWORD")
	Env.MainDBHost = os.Getenv("MAIN_DB_HOST")
	Env.MainDBSSLMode = os.Getenv("MAIN_DB_SSLMODE")
	Env.MainDBPort = os.Getenv("MAIN_DB_PORT")
}
