package repository

import "GenesisTask/model"

type UserRepository interface {
	IsExist(user *model.User) bool
	Add(user *model.User) error
	GetUsers() *[]model.User
}
