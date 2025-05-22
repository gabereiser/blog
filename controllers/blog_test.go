package controllers

import (
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewBlogController(t *testing.T) {
	tests := []struct {
		name string
		want *BlogController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlogController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlogController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlogController_RegisterRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		bc   *BlogController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			bc.RegisterRoutes(tt.args.app)
		})
	}
}

func TestBlogController_getAllPosts(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.getAllPosts(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.getAllPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_getPostByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.getPostByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.getPostByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_createPost(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.createPost(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.createPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_updatePost(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.updatePost(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.updatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_deletePost(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.deletePost(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.deletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_getCommentsByPostID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.getCommentsByPostID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.getCommentsByPostID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_createComment(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.createComment(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.createComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_updateComment(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.updateComment(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.updateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBlogController_deleteComment(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		bc      *BlogController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlogController{}
			if err := bc.deleteComment(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BlogController.deleteComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
