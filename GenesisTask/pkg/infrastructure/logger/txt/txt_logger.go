package logger

import (
	"GenesisTask/pkg/application"
	logger "GenesisTask/pkg/infrastructure/logger/common"
	"GenesisTask/pkg/utils"
	"fmt"
	"log"
	"os"
)

type LoggerFiles struct {
	debugLoggerPath string
	errorLoggerPath string
	infoLoggerPath  string
}

func NewLoggerFiles(d, e, i string) LoggerFiles {
	return LoggerFiles{d, e, i}
}

func EnsureLogFilesExist(f LoggerFiles) {
	utils.EnsureFileExists(f.debugLoggerPath)
	utils.EnsureFileExists(f.errorLoggerPath)
	utils.EnsureFileExists(f.infoLoggerPath)
}

type TxtLogger struct {
	LoggerFiles
}

func NewTxtLogger(files LoggerFiles) application.Logger {
	return &TxtLogger{files}
}

func (l *TxtLogger) LogDebug(v ...any) {
	file, err := os.OpenFile(l.debugLoggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(logger.Debug, fmt.Sprint(v))
}

func (l *TxtLogger) LogError(v ...any) {
	file, err := os.OpenFile(l.errorLoggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(logger.Error, fmt.Sprint(v))
}

func (l *TxtLogger) LogInfo(v ...any) {
	file, err := os.OpenFile(l.infoLoggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(logger.Info, fmt.Sprint(v))
}
