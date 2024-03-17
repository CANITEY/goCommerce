package jwt

import (
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = uuid
	secretToken, err := token.SigningString()
	if err != nil {
		return "", err
	}

	return secretToken, nil
}


