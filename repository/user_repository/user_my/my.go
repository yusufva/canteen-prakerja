package user_my

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/repository/user_repository"

	"gorm.io/gorm"
)

type userMY struct {
	db *gorm.DB
}

func NewUserMy(db *gorm.DB) user_repository.UserRepository {
	return &userMY{
		db: db,
	}
}

func (u *userMY) CreateNewUser(userPayload *entity.User) custerrs.MessageErr {
	result := u.db.Create(&userPayload)

	if result.Error != nil {
		if result.Error.Error() == gorm.ErrDuplicatedKey.Error() {
			return custerrs.NewConflictError("this username has been used")
		}
		return custerrs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *userMY) GetUserById(userId int) (*entity.User, custerrs.MessageErr) {
	user := entity.User{ID: userId}
	err := u.db.First(&user, "id = ?", userId).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("user not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}

func (u *userMY) GetUserByUsername(username string) (*entity.User, custerrs.MessageErr) {
	user := entity.User{Username: username}
	err := u.db.First(&user, "username = ?", username).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("user not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}
