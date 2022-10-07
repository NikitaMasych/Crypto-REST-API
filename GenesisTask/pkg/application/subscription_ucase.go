package application

import (
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/errors"
)

type SubscriptionRepository struct {
	storage SubscriptionStorage
	logger  Logger
}

func NewSubscriptionRepository(storage SubscriptionStorage,
	logger Logger) *SubscriptionRepository {
	return &SubscriptionRepository{storage, logger}
}

func (r *SubscriptionRepository) Subscribe(user models.User) error {
	for _, pair := range user.SubscribedPairs {
		subscription := models.NewSubscription(*user.GetEmailAddress(), pair)
		if r.storage.IsSaved(*subscription) {
			r.logger.LogError(errors.ErrAlreadySubscribed)
			return errors.ErrAlreadySubscribed
		} else {
			if err := r.storage.AddSubscription(*subscription); err != nil {
				r.logger.LogError(err)
				return err
			}
		}
	}
	return nil
}
