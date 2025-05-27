package main

import "gorm.io/gorm"

type Student struct {
	ID    uint
	Name  string `gorm:"size:10"`
	Age   uint
	Grade string `gorm:"size:5"`
}

type Account struct {
	ID      uint
	Balance uint
}
type Transaction struct {
	ID          uint
	FromAccount uint
	ToAccount   uint
	Amount      uint
}

// employees 表，包含字段 id 、 name 、 department 、 salary 。
type Employee struct {
	ID         uint
	Name       string `gorm:"size:10"`
	Department string `gorm:"size:10"`
	Salary     uint
}

// books 表，包含字段 id 、 title 、 author 、 price 。
type Book struct {
	ID     uint
	Title  string `gorm:"size:10"`
	Author string `gorm:"size:10"`
	Price  uint
}

// User 用户模型
type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	PostsCount int    `gorm:"default:0"`           // 新增：用户文章数量统计
	Posts      []Post `gorm:"foreignKey:AuthorID"` // 一对多关系: 用户拥有多篇文章
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title         string    `gorm:"type:varchar(200);not null"`
	Content       string    `gorm:"type:text;not null"`
	Published     bool      `gorm:"default:false"`
	CommentsNum   int       `gorm:"default:0"`                      // 新增：评论数量统计
	CommentStatus string    `gorm:"type:varchar(20);default:'无评论'"` // 新增：评论状态
	AuthorID      uint      `gorm:"index;not null"`                 // 外键，指向用户
	Author        User      `gorm:"foreignKey:AuthorID"`            // 反向引用
	Comments      []Comment `gorm:"foreignKey:PostID"`              // 一对多关系: 文章有多条评论
}

// 在Post模型中添加AfterCreate钩子
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的PostsCount
	result := tx.Model(&User{}).Where("id = ?", p.AuthorID).
		Update("posts_count", gorm.Expr("posts_count + 1"))

	if result.Error != nil {
		return result.Error
	}

	// 如果更新行数为0，说明用户不存在
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"index;not null"`    // 外键，指向文章
	Post    Post   `gorm:"foreignKey:PostID"` // 反向引用
	UserID  uint   `gorm:"index;not null"`    // 外键，指向评论者
	User    User   `gorm:"foreignKey:UserID"` // 评论者信息
}

// 在Comment模型中添加AfterDelete钩子
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 查询当前文章的剩余评论数量
	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	// 更新文章的评论状态
	updateData := map[string]interface{}{
		"comments_num": count,
	}

	// 如果评论数为0，更新状态为"无评论"
	if count == 0 {
		updateData["comment_status"] = "无评论"
	}

	// 执行更新
	if err := tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}
