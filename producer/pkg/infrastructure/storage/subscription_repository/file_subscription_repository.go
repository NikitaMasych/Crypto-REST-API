package storage

import (
	"bufio"
	"os"
	"producer/config"
	"producer/pkg/application"
	"producer/pkg/domain/models"
	"producer/pkg/errors"
	"strings"
)

type SubscriptionFileRepository struct {
	path   string
	logger application.Logger
}

func NewSubscriptionFileRepository(logger application.Logger) application.SubscriptionStorage {
	path := config.StorageFile
	return &SubscriptionFileRepository{path, logger}
}

func (r *SubscriptionFileRepository) IsSaved(subscription models.Subscription) bool {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		r.logger.LogError(err)
		os.Exit(1)
	}
	r.logger.LogDebug("Opened storage file: " + r.path)
	defer func() {
		file.Close()
		r.logger.LogDebug("Closed storage file: " + r.path)
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if subscription.ToString() == scanner.Text() {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		r.logger.LogError(err)
		os.Exit(1)
	}

	return false
}

func (r *SubscriptionFileRepository) AddSubscription(subscription models.Subscription) error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	r.logger.LogDebug("Opened storage file: " + r.path)
	defer func() {
		file.Close()
		r.logger.LogDebug("Closed storage file: " + r.path)
	}()

	_, err = file.WriteString(subscription.ToString())
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return err
}

func (r *SubscriptionFileRepository) GetAll() []models.Subscription {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		r.logger.LogError(err)
		os.Exit(1)
	}
	r.logger.LogDebug("Opened storage file: " + r.path)
	defer func() {
		file.Close()
		r.logger.LogDebug("Closed storage file: " + r.path)
	}()

	scanner := bufio.NewScanner(file)
	var subscriptions []models.Subscription
	for scanner.Scan() {
		line := scanner.Text()
		subscription, err := ComposeSubscription(line)
		if err != nil {
			r.logger.LogError(err)
			os.Exit(1)
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := scanner.Err(); err != nil {
		r.logger.LogError(err)
		os.Exit(1)
	}

	return subscriptions
}

func ComposeSubscription(line string) (models.Subscription, error) {
	emailSepIndex := strings.Index(line, config.EmailAddressSeparator)
	currencySepIndex := strings.Index(line, config.CurrencyPairSeparator)
	if emailSepIndex == -1 || currencySepIndex == -1 {
		emptyEmail := *models.NewEmailAddress("")
		emptyPair := *models.NewCurrencyPair("", "")
		emptySubscription := *models.NewSubscription(emptyEmail, emptyPair)
		return emptySubscription, errors.ErrCouldNotMarshallSubscription
	}
	email := models.NewEmailAddress(line[:emailSepIndex])
	base := line[emailSepIndex+1 : currencySepIndex]
	quote := line[currencySepIndex+1:]
	pair := models.NewCurrencyPair(base, quote)
	return *models.NewSubscription(*email, *pair), nil
}
