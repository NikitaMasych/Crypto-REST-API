package email

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"log"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

type GomailSender struct{}

func NewGomailSender() application.EmailSender {
	return &GomailSender{}
}

func (g *GomailSender) SendRateEmails(rate models.CurrencyRate,
	emails []models.EmailAddress) {
	cfg := config.Get()
	dialer := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort,
		cfg.EmailAddress, cfg.EmailPassword)

	msg := composeMessage(rate)

	for _, email := range emails {
		to := email.GetAddress()
		msg.SetHeader("To", to)

		if err := dialer.DialAndSend(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func composeMessage(rate models.CurrencyRate) *gomail.Message {
	subject := "Currency Rate"
	pair := rate.GetCurrencyPair()
	body := pair.GetBase() + "-" + pair.GetQuote() + " rate on " + time.Now().String() + " : " +
		strconv.FormatFloat(rate.GetPrice(), 'f', -1, 64)

	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Get().EmailAddress)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", "<p style=\"font: 20px Times New Roman, italic\">"+body+"</p>")

	return msg
}
