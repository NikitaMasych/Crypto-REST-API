package emails

import (
	"GenesisTask/config"
	"GenesisTask/crypto"
	"GenesisTask/model"
	"log"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmails(users *[]model.User) {
	cfg := config.Get()
	dialer := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort,
		cfg.EmailAddress, cfg.EmailPassword)

	msg := composeMessage()

	for _, user := range *users {
		to := user.GetEmail()
		msg.SetHeader("To", to)

		if err := dialer.DialAndSend(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func composeMessage() *gomail.Message {
	price, err := crypto.GetCryptoRate()
	if err != nil {
		log.Fatal("Unable to get bitcoin price!")
	}
	subject := "Crypto Rate"
	body := config.Get().BaseCurrency + " price on " + time.Now().String() + " : " +
		strconv.FormatFloat(price, 'f', -1, 64) + config.Get().QuotedCurrency

	msg := gomail.NewMessage()
	msg.SetHeader("From", config.Get().EmailAddress)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", "<p style=\"font: 20px Times New Roman, italic\">"+body+"</p>")

	return msg
}
