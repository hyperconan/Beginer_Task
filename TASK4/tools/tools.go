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
	return uint(uid.(float64)), nil
}
