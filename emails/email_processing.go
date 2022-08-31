package emails

import (
	"GenesisTask/config"
	"GenesisTask/crypto"
	"bufio"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmails() {
	cfg := config.Get()
	path := cfg.StorageFile
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dialer := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort,
		cfg.EmailAddress, cfg.EmailPassword)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	msg := composeMessage()

	for scanner.Scan() {
		to := scanner.Text()
		msg.SetHeader("To", to)

		if err := dialer.DialAndSend(msg); err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func composeMessage() *gomail.Message {
	price, err := crypto.GetConfigCurrencyRate()
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
