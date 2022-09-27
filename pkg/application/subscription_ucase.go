package application

import (
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/errors"
)

type SubscriptionRepository struct {
	storage EmailAddressesStorage
}

func NewSubscriptionRepository(storage EmailAddressesStorage) *SubscriptionRepository {
	return &SubscriptionRepository{storage}
}

func (r *SubscriptionRepository) Subscribe(user models.User) error {
	if r.storage.IsSaved(*user.GetEmailAddress()) {
		return errors.ErrAlreadySubscribed
	}
	err := r.storage.AddEmail(*user.GetEmailAddress())
	return err
}
