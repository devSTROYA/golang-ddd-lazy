package main

import (
	"lazy/infrastructure/config"
	"log"
)

func main() {
	app := InitializeApp()
	config := config.NewEnv()
	log.Fatal(app.Start(config.Port))
}
