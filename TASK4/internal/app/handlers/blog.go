package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hyperconan.com/blog_sys/internal/dao"
	"hyperconan.com/blog_sys/tools"
)

// 博客模块 实现文章的创建、读取、更新、删除功能、发表评论和评论读取

func CreatePost(c *gin.Context) {
	var post dao.Post
	// 从 JSON 绑定文章内容
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断从 Context 中获取的 user_info 是否是 jwt.MapClaims 的几种方式：

	// 方式1：使用类型断言（推荐，两步完成）
	// value, exists := c.Get("user_info")
	// if !exists {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息不存在"})
	//     return
	// }
	// userInfo, ok := value.(jwt.MapClaims)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
	//     return
	// }

	// 方式2：先检查是否存在，再进行类型断言
	// value, exists := c.Get("user_info")
	// if !exists {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息不存在"})
	//     return
	// }
	// userInfo, ok := value.(jwt.MapClaims)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息类型不匹配"})
	//     return
	// }

	// 方式3：使用 switch 进行类型判断
	// value, exists := c.Get("user_info")
	// if !exists {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息不存在"})
	//     return
	// }
	// switch v := value.(type) {
	// case jwt.MapClaims:
	//     userInfo := v
	//     // 使用 userInfo
	//     _ = userInfo
	// default:
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息类型不匹配"})
	//     return
	// }

	// 实际使用：方式1（最简洁）
	// 注意：c.Get() 返回 (value, exists)，需要先接收这两个值
	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	post.UserID = uid
	if err := dao.Db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post_id": post.ID})
}

func GetAllPosts(c *gin.Context) {
	var posts []dao.Post
	err := dao.Db.Omit("Content").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

type updatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	pid := c.Param("post_id")
	if pid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		return
	}

	reqBody := updatePostReq{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pidInt, _ := strconv.Atoi(pid)
	post := dao.Post{}

	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	dao.Db.Where("id = ? and user_id = ?", pidInt, uid).First(&post)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Title = reqBody.Title
	post.Content = reqBody.Content

	if err := dao.Db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	var post dao.Post
	postID := c.Param("post_id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		return
	}
	pid, _ := strconv.Atoi(postID)
	post.ID = uint(pid)

	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} // uid为interface{},需要进行断言 转换为 float64 再转换为 uint

	if err := dao.Db.Delete(&post, "user_id = ?", uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func PostComment(c *gin.Context) {
	var comment dao.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	comment.UserID = uid
	if err := dao.Db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment posted successfully", "comment_id": comment.ID})
}

func GetCommentsByPostID(c *gin.Context) {
	var comments []dao.Comment
	postID := c.Param("post_id")
	err := dao.Db.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
