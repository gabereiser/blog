package middleware

import (
	"errors"

	"github.com/gabereiser/blog/config"
	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/data/models"
	"github.com/gabereiser/blog/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrMissingOrMalformedJWT = errors.New("Missing or malformed JWT")
	ErrInvalidOrExpiredJWT   = errors.New("Invalid or expired JWT")
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Get("SECRET"))},
		ErrorHandler: NotAuthorizedHandler,
	})
}

func NotAuthorizedHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func GetCurrentUser(c *fiber.Ctx) *models.User {
	// Get the JWT token from the cookie
	token := c.Cookies("auth")
	if token == "" {
		return nil
	}

	// Parse the token
	_, userID, err := utils.ParseJWT(token)
	if err != nil {
		return nil
	}
	// Get the user from the database
	user, err := database.Find(&models.User{}, userID)
	if err != nil {
		return nil
	}
	return user
}
