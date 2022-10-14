package services

import "producer/pkg/domain/models"

type SubscriptionService interface {
	Subscribe(user *models.User) error
}
