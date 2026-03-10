package main

import (
	"be-test/config"
	"be-test/router"
)

func main() {
	config.InitEnv()
	config.InitMainDB()

	r := router.HttpRouter()

	r.Run(
		":" +
			config.Env.AppPort,
	)

}
