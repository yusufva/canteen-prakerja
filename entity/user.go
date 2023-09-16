package entity

import (
	"canteen-prakerja/pkg/custerrs"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var invalidTokenErr = custerrs.NewUnauthenticatedError("invalid token")

var secret_key = os.Getenv("Secret_Key")

type User struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"unique;not null;varchar(255)" json:"username"`
	Password  string    `gorm:"not null;type:text;" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp(3)" json:"updated_at"`
}

func (u *User) parseToken(tokenString string) (*jwt.Token, custerrs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) custerrs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.ID = int(id)
	}

	if username, ok := claim["username"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Username = username
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) custerrs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}
	//"Bearer aksdnvokaenovkbnk.kdnvaokn.okanvoknv"

	//[]string{"Bearer", "aksdnvokaenovkbnk.kdnvaokn.okanvoknv"}
	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"exp":      time.Now().Add(time.Minute * 120).Unix(),
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(secret_key))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() custerrs.MessageErr {
	salt := 8

	userPassword := []byte(u.Password)

	bs, err := bcrypt.GenerateFromPassword(userPassword, salt)

	if err != nil {
		return custerrs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
