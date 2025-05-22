package blog

import (
	"errors"
	"time"

	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/data/models"
)

func GetPosts(page int, limit int) ([]models.Post, error) {
	// This function retrieves posts from the database with pagination
	// It should return a slice of Post structs and an error if any
	var posts []models.Post
	err := database.DB().Offset((page - 1) * limit).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostsByAuthor(userID uint, page int, limit int) ([]models.Post, error) {
	// This function retrieves posts by a specific user from the database with pagination
	// It should return a slice of Post structs and an error if any
	var posts []models.Post
	err := database.DB().Where("user_id = ?", userID).Offset((page - 1) * limit).Limit(limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostByID(id uint) (models.Post, error) {
	// This function retrieves a post by its ID from the database
	// It should return a Post struct and an error if any
	var post models.Post
	err := database.DB().First(&post, id).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func CreatePost(author models.User, title string, contents string) (models.Post, error) {
	// This function creates a new post in the database
	// It should return the created Post struct and an error if any
	post := models.Post{
		Title:     title,
		Content:   contents,
		AuthorID:  author.ID,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := database.DB().Create(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func UpdatePost(author models.User, post models.Post) (models.Post, error) {
	// This function updates an existing post in the database
	// It should return the updated Post struct and an error if any
	if post.AuthorID != author.ID {
		return post, errors.New("you are not the author of this post")
	}
	err := database.DB().Save(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func DeletePost(author models.User, post models.Post) error {
	// This function deletes a post by its ID from the database
	// It should return an error if any
	if author.ID != post.AuthorID {
		return errors.New("you are not the author of this post")
	}
	err := database.DB().Delete(&models.Post{}, post.ID).Error
	if err != nil {
		return err
	}
	return nil
}
