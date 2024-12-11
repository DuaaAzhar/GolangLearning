package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userid": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		fmt.Println("secretKey==> ", []byte(secretKey))
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("unable to parse token")
	}
	tokenIsValid := parsedToken.Valid
	fmt.Println("tokenIsValid==>>> ", tokenIsValid)
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println("claims ==>>", claims)
	if !ok {
		return 0, errors.New("invalid token class")
	}

	value, ok := claims["userid"]
	if !ok {
		return 0, fmt.Errorf("userId not found in token claims")
	}
	fmt.Println("userId in verify token===> ", value)

	floatValue, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("userId in token claims is not a valid float64")
	}
	userId := int64(floatValue)

	// email := claims["email"].(string)
	// userId := int64(claims["userId"].(float64))

	return userId, nil

}
