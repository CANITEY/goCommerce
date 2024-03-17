package jwt

import (
	"github.com/golang-jwt/jwt"
	"os"
)

var key = []byte(os.Getenv("KEY"))

func GenerateJWT(uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid
	secretToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return secretToken, nil
}


