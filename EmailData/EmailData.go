package EmailData

import (
	"GenesisTask/BitcoinPrice"
	"bufio"
	"log"
	"os"
	"time"

	gomail "gopkg.in/gomail.v2"
)

// This file will contain all subscribed emails
var path = "emails.txt"

func CreateFile() {
	// Check if file exists
	_, err := os.Stat(path)

	// Create it, if not
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func AddEmail(email string) int {
	// Open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if email == scanner.Text() {
			return 409 // File already in database
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write email to the file database
	_, err = file.WriteString(email + "\n")
	if err != nil {
		log.Fatal(err)
	}

	// Save file changes
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	return 200
}

func SendEmails() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Sender data
	from := "genesisbriefingnm@gmail.com"
	// Using 16 characters special gmail password
	password := "kszosdngchzoszqr" 

	dialer := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// Message content
	price, code := BitcoinPrice.GetBitcoinPrice()
	if code != 200 {
		log.Fatal("Unable to get bitcoin price!")
	}

	subject := "BTC Price"
	body := "Bitcoin price on " + time.Now().String() + " : " + price + " UAH"

	// Compose message
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	// Add some fanciness to the message
	msg.SetBody("text/html", "<p style=\"font: 20px Times New Roman, italic\">"+body+"</p>")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		to := scanner.Text()
		msg.SetHeader("To", to)

		// Send the email
		if err := dialer.DialAndSend(msg); err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
