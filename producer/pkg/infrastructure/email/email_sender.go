package email

import (
	"fmt"
	"os"
	"producer/config"
	"producer/pkg/application"
	"producer/pkg/domain/models"
	"strconv"

	"gopkg.in/gomail.v2"
)

type GomailSender struct {
	logger application.Logger
}

func NewGomailSender(logger application.Logger) application.EmailSender {
	return &GomailSender{logger}
}

func (g *GomailSender) SendRatesEmail(rates []models.CurrencyRate,
	email models.EmailAddress) {

	dialer := gomail.NewDialer(config.SMTPHost, config.SMTPPort,
		config.EmailAddress, config.EmailPassword)
	g.logger.LogDebug(fmt.Sprintf("Initialized new gomail dialer: {Host: %s, Port: %d}",
		config.SMTPHost, config.SMTPPort))
	msg := composeMessage(rates, email)
	g.logger.LogInfo("Message composed")
	if err := dialer.DialAndSend(msg); err != nil {
		g.logger.LogError(err)
		os.Exit(1)
	}
}

func composeMessage(rates []models.CurrencyRate, email models.EmailAddress) *gomail.Message {
	var subject string
	if len(rates) == 1 {
		subject = "Currency Rate"
	} else {
		subject = "Currency Rates"
	}
	body := composeBody(rates)

	msg := gomail.NewMessage()
	msg.SetHeader("To", email.ToString())
	msg.SetHeader("From", config.EmailAddress)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", "<p style=\"font: 20px Times New Roman, italic\">"+body+"</p>")

	return msg
}

func composeBody(rates []models.CurrencyRate) string {
	var body string
	for _, rate := range rates {
		pair := rate.GetCurrencyPair()
		rateInfo := pair.GetBase() + config.CurrencyPairSeparator + pair.GetQuote() +
			" rate on " + rate.GetTimestamp().String() + " : " +
			strconv.FormatFloat(rate.GetPrice(), 'f', -1, 64)
		body += rateInfo + "<br />"
	}
	return body
}
