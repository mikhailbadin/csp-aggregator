package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mikhailbadin/csp-aggregator/config"
	"github.com/mikhailbadin/csp-aggregator/db"
	"github.com/mikhailbadin/csp-aggregator/server"
)

func main() {
	godotenv.Load()
	if err := config.Init(); err != nil {
		log.Fatalln("cannot initialize config:", err.Error())
	}
	if err := db.Init(); err != nil {
		log.Fatalln("cannot initialize db:", err.Error())
	}
	server.Init()
}
