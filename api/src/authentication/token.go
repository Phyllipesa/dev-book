package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken returns a signed token with user permissions
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("DevBookGO")) // secret
}
