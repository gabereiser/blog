package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/data/models"
	"github.com/gabereiser/blog/utils"
	"github.com/gofiber/fiber/v2"
)

func TestNewAuthController(t *testing.T) {
	tests := []struct {
		name string
		want *AuthController
	}{
		{
			name: "Test NewAuthController",
			want: &AuthController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthController_RegisterRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		ac   *AuthController
		args args
	}{
		{
			name: "Test RegisterRoutes",
			ac:   NewAuthController(),
			args: args{
				app: fiber.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewAuthController()
			ac.RegisterRoutes(tt.args.app)
		})
	}
}

func TestAuthController_RegisterAnonymousRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		ac   *AuthController
		args args
		want *AuthController
	}{
		{
			name: "Test RegisterAnonymousRoutes",
			ac:   NewAuthController(),
			args: args{
				app: fiber.New(),
			},
			want: NewAuthController(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AuthController{}
			if got := ac.RegisterAnonymousRoutes(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthController.RegisterAnonymousRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthController_loginHandler(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test loginHandler",
			wantErr: false,
		},
	}
	database.Init(true)
	hashword, err := utils.HashPassword("password")
	tx := database.DB().Save(&models.User{
		Username:  "user",
		Password:  hashword,
		Email:     "test@test.org",
		FirstName: "Test",
		LastName:  "User",
	})
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	if tx.Error != nil {
		t.Fatalf("Failed to create test user: %v", tx.Error)
	}

	app := fiber.New(fiber.Config{
		AppName: "Test",
	})
	ac := NewAuthController()
	ac.RegisterAnonymousRoutes(app).RegisterRoutes(app)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := json.RawMessage(`{"username":"user","password":"password"}`)
			buf := bytes.NewReader(body)
			req, err := http.NewRequestWithContext(context.Background(), "POST", "/auth/login", buf)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			res, err := app.Test(req, 1000)
			if err != nil {
				t.Fatalf("Failed to test request: %v", err)
			}
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected status code 200, got %d", res.StatusCode)
			}
			// Check the response body
			var responseBody map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}
			if _, ok := responseBody["token"]; !ok {
				t.Fatalf("Expected token in response body, got %v", responseBody)
			}
			// Check the response headers
			if authCookie := res.Header.Get("Set-Cookie"); authCookie == "" {
				t.Fatalf("Expected auth cookie in response headers, got %v", res.Header)
			}

		})
	}
}

func TestAuthController_logoutHandler(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name string
	}{
		{
			name: "Test logoutHandler",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			ac := NewAuthController()
			ac.RegisterAnonymousRoutes(app).RegisterRoutes(app)

			req, err := http.NewRequestWithContext(context.Background(), "GET", "/auth/logout", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			res, err := app.Test(req, 1)
			if err != nil {
				t.Fatalf("Failed to test request: %v", err)
			}
			if res.StatusCode != http.StatusOK {
				t.Fatalf("Expected status code 200, got %d", res.StatusCode)
			}
			// Check the response body
			var responseBody map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}
			if _, ok := responseBody["message"]; !ok {
				t.Fatalf("Expected message in response body, got %v", responseBody)
			}
		})
	}
}

func TestAuthController_registerHandler(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		ac      *AuthController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AuthController{}
			if err := ac.registerHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthController.registerHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthController_refreshHandler(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		ac      *AuthController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AuthController{}
			if err := ac.refreshHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthController.refreshHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthController_profileHandler(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		ac      *AuthController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AuthController{}
			if err := ac.profileHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthController.profileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
