package dao

import "gorm.io/gorm"

type Post struct {
	//存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
	gorm.Model
	Title    string    `json:"title" gorm:"index:idx_title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"user_id" gorm:"index:idx_uid"`
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"`
}

func (post *Post) getComments() error {
	return Db.Where("post_id = ?", post.ID).Find(&post.Comments).Error
}

func getAllPosts() ([]Post, error) {
	var posts []Post
	err := Db.Omit("Content").Find(&posts).Error
	return posts, err
}
