package entity

import (
	"crudexample/utils"
	"github.com/golang-jwt/jwt"
	"github.com/goonode/mogo"
	"time"
)

//User struct is to handle user data
type User struct {
	mogo.DocumentModel `bson:",inline" coll:"users"`
	Email              string `idx:"{email},unique" json:"email" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Name               string `json:"name"`
	VerifiedAt         *time.Time
}

//GetJwtToken returns jwt token with user email claims
func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})
	secretKey := utils.GetEnv("TOKEN_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func init() {
	mogo.ModelRegistry.Register(User{})
}
