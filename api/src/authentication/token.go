package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	return token.SignedString([]byte(config.SecretKey)) // secret
}

// ValidationToken verify if token in request is valid
func ValidationToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// ExtractUserID return userID from token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		/*
			No caso o primeiro parâmetro que o ParseUint recebe é uma string,
			porém o que esta retornando do permissions é uma interface.

			Então, é preciso converter o valor para string para depois
			converte-lo a um uint.

			OBS: A funçao CreateToken recebe um "userID uint64" no parâmetro,
			porém por padrão ele ficara salvo como um Float.
			Então, é preciso converter de Float para String e então passa-la
			ao ParseUint.
		*/
		userID, erro := strconv.ParseUint(
			fmt.Sprintf("%.0f", permissions["userId"]), 10, 64,
		)
		if erro != nil {
			return 0, erro
		}

		return userID, nil
	}

	return 0, errors.New("invalid token")
}

// extractToken if token in correct format, extract from request and return
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnVerificationKey verify if signature method is correct
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
