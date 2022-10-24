package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerUrl   string
	DatabaseUrl string
)

func init() {
	loadEnv()
	setupVariables()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func setupVariables() {
	ServerUrl = os.Getenv("SERVER_URL")
	DatabaseUrl = os.Getenv("MYSQL_DSN")
}
