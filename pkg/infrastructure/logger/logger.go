package logger

import (
	"GenesisTask/config"
	"log"
	"os"
	"time"

	"gopkg.in/resty.v0"
)

func LogProviderResponse(timestamp time.Time, provider string, resp *resty.Response) {
	loggerPath := config.Get().LoggerFile
	file, err := os.OpenFile(loggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(provider, "- Response at "+timestamp.String()+" :", resp.Status(), string(resp.Body))
}
