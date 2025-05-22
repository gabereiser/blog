package controllers

import (
	"time"

	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/data/models"
	"github.com/gabereiser/blog/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}
func (ac *AuthController) RegisterRoutes(app *fiber.App) {
	// Register the routes
	// These routes require authentication
	app.Post("/auth/refresh", ac.refreshHandler)
	app.Get("/auth/me", ac.profileHandler)
}
func (ac *AuthController) RegisterAnonymousRoutes(app *fiber.App) *AuthController {
	// Register the anonymous routes
	// These routes do not require authentication
	app.Post("/auth/login", ac.loginHandler)
	app.Post("/auth/register", ac.registerHandler)
	app.Get("/auth/logout", ac.logoutHandler)
	return ac
}

// AuthLoginHandler handles user login
func (ac *AuthController) loginHandler(c *fiber.Ctx) error {
	// Get the request body
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Find the user in the database
	var dbUser models.User
	if tx := database.DB().Where("username = ?", loginRequest.Username).First(&dbUser); tx.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Unknown username",
		})
	}

	// Check the password
	if !utils.CheckPasswordHash(loginRequest.Password, dbUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	// Generate a JWT token
	token, err := utils.GenerateJWT(dbUser.Username, dbUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{
		"token": token,
	})
}

// AuthLogoutHandler handles user logout
// It clears the JWT token from the cookie
// and returns a success message
func (ac *AuthController) logoutHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

// AuthRegisterHandler handles user registration
// It checks if the username and email already exist
// and creates a new user in the database
// It returns a success message or an error message
// depending on the result of the operation
func (ac *AuthController) registerHandler(c *fiber.Ctx) error {
	// Get the request body
	var registerRequest struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var user models.User
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	// Check if the username already exists
	if _, err := database.Where(&user, "username = ?", registerRequest.Username); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username already exists",
		})
	}
	// Check if the email already exists
	if _, err := database.Where(&user, "email = ?", registerRequest.Email); err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}
	user = models.User{
		Username:  registerRequest.Username,
		Password:  registerRequest.Password,
		Email:     registerRequest.Email,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
	}
	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}
	user.Password = hashedPassword

	// Create the user in the database
	if _, err := database.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// AuthRefreshHandler handles JWT token refresh
// It checks if the JWT token is valid
// and generates a new token if it is
// It returns the new token or an error message
func (ac *AuthController) refreshHandler(c *fiber.Ctx) error {
	// Get the JWT token from the cookie
	token := c.Cookies("auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	// Parse the token
	username, userID, err := utils.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
	}

	// Generate a new JWT token
	newToken, err := utils.GenerateJWT(username, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    newToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{
		"token": newToken,
	})
}

func (ac *AuthController) profileHandler(c *fiber.Ctx) error {
	// Get the JWT token from the cookie
	token := c.Cookies("auth")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	// Parse the token
	_, userID, err := utils.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
	}

	// Find the user in the database
	var dbUser models.User
	if _, err := database.Where(&dbUser, "id = ?", userID); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}
	// Remove the password from the user object
	var profile = struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		Username:  dbUser.Username,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
	}

	return c.JSON(profile)
}
