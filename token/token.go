package token

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dshills/apix/config"
)

var secretKey []byte

// SetSecret sets the secret key used for Token and Decode
func SetSecret(key []byte) {
	secretKey = key
}

// Config sets the JWT secret
func Config(con *config.Server) error {
	if con.JWTKey == "" {
		return fmt.Errorf("%v is required, quitting", con.Prefix+"_JWT_KEY")
	}
	SetSecret([]byte(con.JWTKey))
	return nil
}

// Decode will return a claim from a token string
func Decode(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Failed to decode token")
}
