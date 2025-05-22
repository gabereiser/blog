package controllers

import (
	"fmt"
	"strconv"

	"github.com/gabereiser/blog/service/blog"
	"github.com/gabereiser/blog/service/middleware"
	"github.com/gofiber/fiber/v2"
)

type BlogController struct {
}

func NewBlogController() *BlogController {
	return &BlogController{}
}
func (bc *BlogController) RegisterRoutes(app *fiber.App) {
	// Post routes
	app.Get("/posts", bc.getAllPosts)
	app.Get("/posts/:id", bc.getPostByID)
	app.Post("/posts/create", middleware.Protected(), bc.createPost)
	app.Put("/posts/:id/update", middleware.Protected(), bc.updatePost)
	app.Delete("/posts/:id/delete", middleware.Protected(), bc.deletePost)
	// Comment routes
	app.Get("/posts/:id/comments", bc.getCommentsByPostID)
	app.Post("/posts/:id/comments/create", middleware.Protected(), bc.createComment)
	app.Put("/posts/:id/comments/:commentID/update", middleware.Protected(), bc.updateComment)
	app.Delete("/posts/:id/comments/:commentID/delete", middleware.Protected(), bc.deleteComment)
}

func (bc *BlogController) getAllPosts(c *fiber.Ctx) error {
	// Get the page and limit from query parameters
	page := c.Query("page", "0")
	limit := c.Query("limit", "10")
	// Convert page and limit to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid page number",
		})
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid limit number",
		})
	}
	// Fetch posts from the blog service
	posts, err := blog.GetPosts(pageInt, limitInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch posts",
		})
	}
	return c.JSON(posts)
}

func (bc *BlogController) getPostByID(c *fiber.Ctx) error {
	// Get the post ID from the URL parameters
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID is required",
		})
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid post ID",
		})
	}
	// Fetch the post from the blog service
	post, err := blog.GetPostByID(uint(postIDInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch post",
		})
	}
	return c.JSON(post)
}

func (bc *BlogController) createPost(c *fiber.Ctx) error {

	// Get the request body
	var postRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&postRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Create the post using the blog service
	post, err := blog.CreatePost(*currentUser, postRequest.Title, postRequest.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create post",
		})
	}
	return c.JSON(post)
}

func (bc *BlogController) updatePost(c *fiber.Ctx) error {
	// Get the post ID from the URL parameters
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID is required",
		})
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid post ID",
		})
	}
	// Get the request body
	var postRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&postRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	post, err := blog.GetPostByID(uint(postIDInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch post",
		})
	}

	if post.AuthorID != currentUser.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not the author of this post",
		})
	}

	post.Title = postRequest.Title
	post.Content = postRequest.Content

	post, err = blog.UpdatePost(*currentUser, post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update post",
		})
	}
	return c.JSON(post)
}

func (bc *BlogController) deletePost(c *fiber.Ctx) error {
	// Get the post ID from the URL parameters
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID is required",
		})
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid post ID",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	post, err := blog.GetPostByID(uint(postIDInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch post",
		})
	}

	if post.AuthorID != currentUser.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not the author of this post",
		})
	}

	err = blog.DeletePost(*currentUser, post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete post",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Post deleted successfully",
	})
}

func (bc *BlogController) getCommentsByPostID(c *fiber.Ctx) error {
	// Get the post ID from the URL parameters
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID is required",
		})
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid post ID",
		})
	}
	// Fetch comments for the post using the blog service
	comments, err := blog.GetCommentsByPostID(uint(postIDInt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch comments",
		})
	}
	return c.JSON(comments)
}

func (bc *BlogController) createComment(c *fiber.Ctx) error {
	// Get the post ID from the URL parameters
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID is required",
		})
	}
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid post ID",
		})
	}
	// Get the request body
	var commentRequest struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&commentRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Create the comment using the blog service
	comment, err := blog.CreateComment(uint(postIDInt), *currentUser, commentRequest.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create comment",
		})
	}
	return c.JSON(comment)
}

func (bc *BlogController) updateComment(c *fiber.Ctx) error {
	// Get the post ID and comment ID from the URL parameters
	postID := c.Params("id")
	commentID := c.Params("commentID")
	if postID == "" || commentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID and Comment ID are required",
		})
	}
	commentIDInt, err := strconv.Atoi(commentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid comment ID",
		})
	}
	// Get the request body
	var commentRequest struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&commentRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Update the comment using the blog service
	comment, err := blog.UpdateComment(uint(commentIDInt), *currentUser, commentRequest.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Failed to update comment %s", err.Error()),
		})
	}
	return c.JSON(comment)
}

func (bc *BlogController) deleteComment(c *fiber.Ctx) error {
	// Get the post ID and comment ID from the URL parameters
	postID := c.Params("id")
	commentID := c.Params("commentID")
	if postID == "" || commentID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Post ID and Comment ID are required",
		})
	}
	commentIDInt, err := strconv.Atoi(commentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid comment ID",
		})
	}
	currentUser := middleware.GetCurrentUser(c)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Delete the comment using the blog service
	err = blog.DeleteComment(uint(commentIDInt), *currentUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete comment",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Comment deleted successfully",
	})
}
