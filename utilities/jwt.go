package utilities

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

var hmacSecret = []byte(os.Getenv("HMAC_SECRET"))

func ParseToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt: unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("jwt: unexpected parsing problem")
}

func SignToken(claims jwt.MapClaims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err = token.SignedString(hmacSecret)
	if err != nil {
		return "", fmt.Errorf("jwt: unexpected signing problem")
	}
	return tokenString, nil
}

func VerifyTokenFromHeader(bearerToken string) (claims jwt.MapClaims, err error) {
	if len(strings.Split(bearerToken, " ")) == 2 {
		return ParseToken(strings.Split(bearerToken, " ")[1])
	}
	return nil, fmt.Errorf("jwt: no token provided")

}
