package jwt

import (
	"github.com/golang-jwt/jwt"
	"os"
)

var KEY = os.Getenv("KEY")

func GenerateJWT(uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid
	secretToken, err := token.SignedString(KEY)
	if err != nil {
		return "", err
	}

	return secretToken, nil
}


