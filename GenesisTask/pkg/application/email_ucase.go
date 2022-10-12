package application

import (
	"GenesisTask/pkg/domain/models"
)

type EmailSenderRepository struct {
	storage   SubscriptionStorage
	sender    EmailSender
	exchanger RateRepository
	logger    Logger
}

func NewEmailSenderRepository(storage SubscriptionStorage,
	sender EmailSender, exchanger RateRepository, logger Logger) *EmailSenderRepository {
	return &EmailSenderRepository{storage, sender, exchanger, logger}
}

func (r *EmailSenderRepository) SendEmailsToUsers() error {
	subscriptions := r.storage.GetAll()
	r.logger.LogDebug("Fetched subscriptions from storage")
	users := matchSubscriptionsToUsers(subscriptions)
	for _, user := range users {
		err := r.sendEmailToUser(user)
		if err != nil {
			r.logger.LogError(err)
			return err
		}
		r.logger.LogInfo("Emails sent to " + user.GetEmailAddress().ToString())
	}
	return nil
}

func (r *EmailSenderRepository) sendEmailToUser(user models.User) error {
	var rates []models.CurrencyRate
	for _, pair := range user.GetSubscribedPairs() {
		rate, err := r.exchanger.GetRate(pair)
		if err != nil {
			r.logger.LogError(err)
			return err
		}
		r.logger.LogDebug("Got " + pair.ToString() + " rate")
		rates = append(rates, rate)
	}
	r.sender.SendRatesEmail(rates, *user.GetEmailAddress())
	return nil
}

func matchSubscriptionsToUsers(subscriptions []models.Subscription) []models.User {
	userComposer := make(map[models.EmailAddress][]models.CurrencyPair)
	for _, subscription := range subscriptions {
		email := *subscription.GetEmailAddress()
		pair := *subscription.GetCurrencyPair()
		userComposer[email] = append(userComposer[email], pair)
	}
	var usersRepo []models.User
	for email, pairs := range userComposer {
		usersRepo = append(usersRepo, *models.NewUser(email, pairs))
	}
	return usersRepo
}
