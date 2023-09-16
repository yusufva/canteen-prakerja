package user_repository

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
)

type UserRepository interface {
	CreateNewUser(userPayload entity.User) custerrs.MessageErr
	GetUserById(userId int) (*entity.User, custerrs.MessageErr)
	GetUserByUsername(username string) (*entity.User, custerrs.MessageErr)
}
