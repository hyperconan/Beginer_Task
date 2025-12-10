package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hyperconan.com/blog_sys/internal/pkg/logger"
)

// JSON 返回统一的 JSON 响应格式
func JSON(c *gin.Context, status int, code string, message string, data any) {
	if status >= http.StatusBadRequest {
		logger.S.Warnw("request failed", "status", status, "code", code, "msg", message, "path", c.FullPath())
	}
	resp := gin.H{
		"code":    code,
		"message": message,
	}
	if data != nil {
		resp["data"] = data
	}
	c.JSON(status, resp)
}

// Error 统一错误响应，记录日志
func Error(c *gin.Context, status int, code string, message string, err error) {
	if err != nil {
		logger.S.Errorw("request error", "status", status, "code", code, "msg", message, "err", err.Error(), "path", c.FullPath())
	} else {
		logger.S.Warnw("request error", "status", status, "code", code, "msg", message, "path", c.FullPath())
	}
	JSON(c, status, code, message, nil)
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	JSON(c, http.StatusOK, "OK", "success", data)
}
