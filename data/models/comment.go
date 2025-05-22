package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `json:"-"`
	AuthorID  uint      `gorm:"not null" json:"author_id"`
	Author    User      `json:"-"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
