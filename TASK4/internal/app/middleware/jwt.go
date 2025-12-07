package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	// JWTSecret JWT签名密钥，与生成token时使用的密钥一致
	JWTSecret = "hyperconan"
)

// JWTAuthMiddleware JWT认证中间件
// 从请求头中获取token并验证其有效性和过期时间
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少Authorization请求头",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization格式错误，应为: Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 解析和验证token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWTSecret), nil
		})

		// 检查token是否有效
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 检查token是否有效且未过期
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token无效或已过期",
			})
			c.Abort()
			return
		}

		// 提取claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无法解析token claims",
			})
			c.Abort()
			return
		}

		// 将 jwt.MapClaims 转换为 map[string]any
		// 因为 jwt.MapClaims 是 map[string]any 的类型别名，需要显式转换
		userInfo := make(map[string]any)
		for k, v := range claims {
			userInfo[k] = v
		}

		c.Set("user_info", userInfo)

		c.Next()
	}
}
