package utils

import (
	"errors"
	"time"

	"github.com/gabereiser/blog/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrValidationError      = errors.New("Invalid token")
	ErrInvalidSigningMethod = errors.New("Invalid signing method")
)

func GenerateJWT(username string, userID uint) (string, error) {
	// Generate a JWT token for the user
	// This is where you would typically use a JWT library to create a token
	// For example:
	claims := jwt.MapClaims{
		"username": username,
		"user_id":  userID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get("SECRET")))

}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// Validate the JWT token
	// This is where you would typically use a JWT library to parse and validate a token
	// For example:
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(config.Get("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ParseJWT(tokenString string) (string, uint, error) {
	// Parse the JWT token and return the user ID
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return "", 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", 0, ErrValidationError
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", 0, ErrValidationError
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", 0, ErrValidationError
	}
	return username, uint(userID), nil
}
