package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DtmCoordinatorAddress = os.Getenv("DTM_COORDINATOR")
	CustomersServerURL    = os.Getenv("CUSTOMERS_SERVICE_URL")
	ServerUrl             string
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
}
