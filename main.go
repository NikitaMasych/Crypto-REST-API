package main

import (
	"GenesisTask/BitcoinPrice"
	"GenesisTask/EmailData"

	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bitcoinPrice(c *gin.Context) {
	value, code := BitcoinPrice.GetBitcoinPrice()
	switch code {
	case 200:
		price, _ := strconv.ParseFloat(value, 64)
		c.JSON(code, gin.H{"description": price})
	case 400:
		c.JSON(code, gin.H{"description": "Invalid status value"})
	}
}

func saveEmail(c *gin.Context) {
	var code int

	email := c.PostForm("email")

	if _, err := mail.ParseAddress(email); err != nil {
		code = 400 // invalid email
	} else {
		code = EmailData.AddEmail(email)
	}

	var status string
	switch code {
	case 200:
		status = "Email added"
	case 400:
		status = "Invalid email"
	case 409:
		status = "Already in database"
	}

	c.JSON(code, gin.H{"description": status})
}

func sendEmails(c *gin.Context) {
	EmailData.SendEmails()
	c.JSON(200, gin.H{"description": "Emails has been sent"})
}

func main() {

	EmailData.CreateFile()

	router := gin.Default()

	router.GET("/api/rate", bitcoinPrice)
	router.POST("/api/subscribe", saveEmail)
	router.POST("/api/sendEmails", sendEmails)

	router.Run(":8080")
}
