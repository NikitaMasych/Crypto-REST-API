package application

import (
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/errors"
)

type SubscriptionRepository struct {
	storage SubscriptionStorage
}

func NewSubscriptionRepository(storage SubscriptionStorage) *SubscriptionRepository {
	return &SubscriptionRepository{storage}
}

func (r *SubscriptionRepository) Subscribe(user models.User) error {
	for _, pair := range user.SubscribedPairs {
		subscription := models.NewSubscription(*user.GetEmailAddress(), pair)
		if r.storage.IsSaved(*subscription) {
			return errors.ErrAlreadySubscribed
		} else {
			if err := r.storage.AddSubscription(*subscription); err != nil {
				return err
			}
		}
	}
	return nil
}
