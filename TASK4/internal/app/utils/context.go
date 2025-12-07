package utils

import (
	"github.com/gin-gonic/gin"
)

// GetUserInfo 从 Context 中安全地获取 user_info
// 返回 user_info 和是否成功获取的布尔值
func GetUserInfo(c *gin.Context) (map[string]any, bool) {
	// 方式1：使用类型断言，ok 表示断言是否成功
	value, exists := c.Get("user_info")
	if !exists {
		return nil, false
	}
	userInfo, ok := value.(map[string]any)
	return userInfo, ok
}

// GetUserInfoWithCheck 从 Context 中获取 user_info 并检查类型
// 如果类型不匹配或不存在，返回错误信息
func GetUserInfoWithCheck(c *gin.Context) (map[string]any, error) {
	// 先检查值是否存在
	value, exists := c.Get("user_info")
	if !exists {
		return nil, ErrUserInfoNotFound
	}

	// 使用类型断言
	userInfo, ok := value.(map[string]any)
	if !ok {
		return nil, ErrUserInfoTypeMismatch
	}

	return userInfo, nil
}

// 错误定义
var (
	ErrUserInfoNotFound     = &ContextError{Message: "用户信息不存在"}
	ErrUserInfoTypeMismatch = &ContextError{Message: "用户信息类型不匹配"}
)

type ContextError struct {
	Message string
}

func (e *ContextError) Error() string {
	return e.Message
}
