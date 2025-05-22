package database

import "github.com/gabereiser/blog/data/models"

func Migrate() {
	// Perform the database migration here
	// This is where you would typically call db.AutoMigrate() for your models
	// For example:
	// db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
