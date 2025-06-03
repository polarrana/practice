package models

import "gorm.io/gorm"

// Post 文章模型
type Post struct {
	gorm.Model
	Title    string `gorm:"not null;size:50"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Comments []Comment
}
