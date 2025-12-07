package modules

import (
	"hyperconan.com/blog_sys/orms"
)

// 博客模块 实现文章的创建、读取、更新、删除功能、发表评论和评论读取

func CreatePost(post *orms.Post) error {
	return orms.Db.Create(post).Error
}

func GetAllPosts() ([]orms.Post, error) {
	var posts []orms.Post
	err := orms.Db.Omit("Content").Find(&posts).Error
	return posts, err
}

func UpdatePost(post *orms.Post) error {
	return orms.Db.Save(post).Error
}

func DeletePost(post *orms.Post) error {
	return orms.Db.Delete(post, "id = ? and user_id = ?", post.ID, post.UserID).Error
}

func postComment(comment *orms.Comment) error {
	return orms.Db.Create(comment).Error
}

func getCommentsByPostID(postID uint) ([]orms.Comment, error) {
	var comments []orms.Comment
	err := orms.Db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
