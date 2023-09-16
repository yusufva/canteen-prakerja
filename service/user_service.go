package service

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/repository/user_repository"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, custerrs.MessageErr)
	Login(loginPayload dto.LoginRequest) (*dto.LoginResponse, custerrs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	userEntity := entity.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	err = userEntity.HashPassword()

	if err != nil {
		return nil, err
	}

	err = us.userRepo.CreateNewUser(userEntity)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "user registered successfully",
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

func (us *userService) Login(loginPayload dto.LoginRequest) (*dto.LoginResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(loginPayload)

	if err != nil {
		return nil, err
	}

	var user *entity.User

	user, err = us.userRepo.GetUserByUsername(loginPayload.Username)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, custerrs.NewUnauthenticatedError("invalid username/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(loginPayload.Password)

	if !isValidPassword {
		return nil, custerrs.NewUnauthenticatedError("invalid username/password")
	}

	response := dto.LoginResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "user has been successfully logged in",
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}
