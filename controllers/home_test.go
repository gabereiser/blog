package controllers

import (
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewHomeController(t *testing.T) {
	tests := []struct {
		name string
		want *HomeController
	}{
		{
			name: "Test NewHomeController",
			want: &HomeController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHomeController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHomeController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHomeController_RegisterRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		hc   *HomeController
		args args
	}{
		{
			name: "Test RegisterRoutes",
			hc:   &HomeController{},
			args: args{
				app: fiber.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := &HomeController{}
			hc.RegisterRoutes(tt.args.app)
		})
	}
}

func Test_home(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := home(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("home() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_status(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := status(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("status() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
