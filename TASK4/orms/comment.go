package orms

import (
	"gorm.io/gorm"
)

type Comment struct {
	//存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
	gorm.Model
	Content string
	UserID  uint `gorm:"index:idx_uid"`
	PostID  uint `gorm:"index:idx_pid"`
}
