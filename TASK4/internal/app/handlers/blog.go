package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"hyperconan.com/blog_sys/internal/dao"
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
	value, exists := c.Get("user_info")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息不存在"})
		return
	}
	userInfo, ok := value.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}

	uid, ok := userInfo["id"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户ID"})
		return
	}
	post.UserID = uint(uid.(float64)) // uid为interface{},需要进行断言 转换为 float64 再转换为 uint
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

func UpdatePost(c *gin.Context) {
	var post dao.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	userinfo, exists := c.Get("user_info")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户信息不存在"})
		return
	}
	value, ok := userinfo.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户信息"})
		return
	}
	uid, ok := value["id"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户ID"})
		return
	}
	uid = uint(uid.(float64)) // uid为interface{},需要进行断言 转换为 float64 再转换为 uint

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
