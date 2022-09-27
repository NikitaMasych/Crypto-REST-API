package logger

import (
	"GenesisTask/config"
	"log"
	"os"

	"gopkg.in/resty.v0"
)

func LogProviderResponse(provider string, resp *resty.Response) {
	loggerPath := config.Get().LoggerFile
	file, err := os.OpenFile(loggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(provider, "- Response:", resp.Status(), string(resp.Body))
}
