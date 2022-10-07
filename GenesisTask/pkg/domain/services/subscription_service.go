package services

import "GenesisTask/pkg/domain/models"

type SubscriptionService interface {
	Subscribe(user *models.User) error
}
