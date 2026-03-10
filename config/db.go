package config

import (
	"be-test/ent"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *ent.Client

func InitMainDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Env.MainDBHost,
		Env.MainDBPort,
		Env.MainDBUser,
		Env.MainDB,
		Env.MainDBPassword,
		Env.MainDBSSLMode,
	)
	client, err := ent.Open(
		"postgres",
		dsn,
	)
	if err != nil {
		log.Fatalf(
			"failed opening connection to postgres: %v",
			err,
		)
	}
	DB = client
}
