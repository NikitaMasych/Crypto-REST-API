package txtlogger

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/infrastructure/logger/logtypes"
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

func CreateTxtLoggerWithConfigSpecs() application.Logger {
	loggerFiles := NewLoggerFiles(config.DebugLogFile,
		config.ErrorsLogFile, config.InfoLogFile)
	EnsureLogFilesExist(loggerFiles)
	logger := NewTxtLogger(loggerFiles)
	return logger
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

	log.Println(logtypes.Debug, fmt.Sprint(v))
}

func (l *TxtLogger) LogError(v ...any) {
	file, err := os.OpenFile(l.errorLoggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(logtypes.Error, fmt.Sprint(v))
}

func (l *TxtLogger) LogInfo(v ...any) {
	file, err := os.OpenFile(l.infoLoggerPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println(logtypes.Info, fmt.Sprint(v))
}
