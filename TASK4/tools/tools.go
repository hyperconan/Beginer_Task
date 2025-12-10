package tools

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdFromContext(c *gin.Context) (uint, error) {
	value, exists := c.Get("user_info")
	if !exists {
		return 0, errors.New("用户信息不存在")
	}
	userInfo, ok := value.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("无法获取用户信息")
	}

	uid, ok := userInfo["id"]
	if !ok {
		return 0, errors.New("无法获取用户ID")
	}
	switch v := uid.(type) {
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	case uint:
		return v, nil
	default:
		return 0, errors.New("用户ID类型不正确")
	}
}
