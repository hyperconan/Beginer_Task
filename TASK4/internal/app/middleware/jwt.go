package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"hyperconan.com/blog_sys/internal/app/response"
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
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "缺少Authorization请求头", nil)
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "Authorization格式错误，应为: Bearer <token>", nil)
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
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "无效的token: "+err.Error(), err)
			c.Abort()
			return
		}

		// 检查token是否有效且未过期
		if !token.Valid {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "token无效或已过期", nil)
			c.Abort()
			return
		}

		// 提取claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "无法解析token claims", nil)
			c.Abort()
			return
		}

		c.Set("user_info", claims)

		c.Next()
	}
}
