package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hyperconan.com/blog_sys/internal/app/response"
	"hyperconan.com/blog_sys/internal/dao"
	"hyperconan.com/blog_sys/tools"
)

// 博客模块 实现文章的创建、读取、更新、删除功能、发表评论和评论读取

func CreatePost(c *gin.Context) {
	var post dao.Post
	// 从 JSON 绑定文章内容
	if err := c.ShouldBindJSON(&post); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), err)
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
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "用户认证失败", err)
		return
	}
	post.UserID = uid
	if err := dao.Db.Create(&post).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "创建文章失败", err)
		return
	}

	response.Success(c, gin.H{"post_id": post.ID})
}

func GetAllPosts(c *gin.Context) {
	var posts []dao.Post
	err := dao.Db.Omit("Content").Find(&posts).Error
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "获取文章列表失败", err)
		return
	}
	response.Success(c, gin.H{"posts": posts})
}

type updatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	pid := c.Param("post_id")
	if pid == "" {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "post_id is required", nil)
		return
	}

	reqBody := updatePostReq{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), err)
		return
	}
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "post_id must be number", err)
		return
	}
	post := dao.Post{}

	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "用户认证失败", err)
		return
	}
	if err := dao.Db.Where("id = ? and user_id = ?", pidInt, uid).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "文章不存在", err)
		} else {
			response.Error(c, http.StatusInternalServerError, "DB_ERROR", "查询文章失败", err)
		}
		return
	}

	post.Title = reqBody.Title
	post.Content = reqBody.Content

	if err := dao.Db.Save(&post).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "更新文章失败", err)
		return
	}

	response.Success(c, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	postID := c.Param("post_id")
	if postID == "" {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "post_id is required", nil)
		return
	}
	pid, err := strconv.Atoi(postID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", "post_id must be number", err)
		return
	}

	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "用户认证失败", err)
		return
	} // uid为interface{},需要进行断言 转换为 float64 再转换为 uint

	delPost := dao.Post{}
	if err := dao.Db.First(&delPost, "id = ? and user_id = ?", pid, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "文章不存在", err)
		} else {
			response.Error(c, http.StatusInternalServerError, "DB_ERROR", "查询文章失败", err)
		}
		return
	}
	if err := dao.Db.Delete(&delPost).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "删除文章失败", err)
		return
	}

	response.Success(c, gin.H{"message": "Post deleted successfully"})
}

func PostComment(c *gin.Context) {
	var comment dao.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), err)
		return
	}
	uid, err := tools.GetUserIdFromContext(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "用户认证失败", err)
		return
	}
	comment.UserID = uid
	if err := dao.Db.Create(&comment).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "创建评论失败", err)
		return
	}

	response.Success(c, gin.H{"comment_id": comment.ID})
}

func GetCommentsByPostID(c *gin.Context) {
	var comments []dao.Comment
	postID := c.Param("post_id")
	err := dao.Db.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "获取评论失败", err)
		return
	}
	response.Success(c, gin.H{"comments": comments})
}
