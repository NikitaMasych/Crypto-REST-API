package logger

import (
	"GenesisTask/config"
	"log"
	"net/http"
	"os"
)

func AddProviderResponseToLog(resp *http.Response) {
	loggerPath := config.Get().LoggerFile
	file, err := os.OpenFile(loggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Println(resp)
}
