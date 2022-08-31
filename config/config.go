package config

import (
	"os"
	"strconv"
	"sync"
)

type Config struct {
	ServerURL          string
	BaseCurrency       string
	QuotedCurrency     string
	CryptoApiFormatUrl string
	EmailAddress       string
	EmailPassword      string
	StorageFile        string
	SMTPHost           string
	SMTPPort           int
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		cfg = Config{
			ServerURL:          os.Getenv(ServerUrl),
			CryptoApiFormatUrl: os.Getenv(CryptoApiFormatUrl),
			BaseCurrency:       os.Getenv(BaseCurrency),
			QuotedCurrency:     os.Getenv(QuotedCurrency),
			EmailAddress:       os.Getenv(EmailAddress),
			EmailPassword:      os.Getenv(EmailPassword),
			StorageFile:        os.Getenv(StorageFile),
			SMTPHost:           os.Getenv(SMTPHost),
		}
		cfg.SMTPPort, _ = strconv.Atoi(os.Getenv(SMTPPort))
	})
	return &cfg
}
