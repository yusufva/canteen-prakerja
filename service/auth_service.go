package service

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invalidTokenErr = custerrs.NewUnauthenticatedError("invalid token")
		bearerToken := c.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserByUsername(user.Username)

		if err != nil {
			c.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		c.Set("userData", user)

		c.Next()
	}
}
