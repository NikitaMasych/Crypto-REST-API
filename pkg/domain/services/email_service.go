package services

type EmailService interface {
	SendEmailsToUsers() error
}
