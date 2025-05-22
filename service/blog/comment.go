package blog

import (
	"errors"
	"time"

	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/data/models"
)

func GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := database.DB().Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetCommentByID(id uint) (models.Comment, error) {
	var comment models.Comment
	err := database.DB().First(&comment, id).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func GetCommentsByAuthor(userID uint, page int, limit int) ([]models.Comment, error) {
	var comments []models.Comment
	err := database.DB().Where("author_id = ?", userID).Offset((page - 1) * limit).Limit(limit).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CreateComment(postID uint, author models.User, content string) (models.Comment, error) {
	var post models.Post
	err := database.DB().First(&post, postID).Error
	if err != nil {
		return models.Comment{}, err
	}

	comment := models.Comment{
		Content:   content,
		PostID:    post.ID,
		Post:      post,
		AuthorID:  author.ID,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = database.DB().Create(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func UpdateComment(id uint, author models.User, content string) (models.Comment, error) {
	var comment models.Comment
	err := database.DB().First(&comment, id).Error
	if err != nil {
		return comment, err
	}
	if comment.AuthorID != author.ID {
		return comment, errors.New("you are not the author of this comment")
	}
	comment.Content = content
	err = database.DB().Save(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func DeleteComment(id uint, author models.User) error {
	var comment models.Comment
	err := database.DB().First(&comment, id).Error
	if err != nil {
		return err
	}

	if comment.AuthorID != author.ID {
		return errors.New("you are not the author of this comment")
	}
	err = database.DB().Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
