package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hyperconan.com/blog_sys/internal/app/response"
	"hyperconan.com/blog_sys/internal/dao"
)

func UserRegister(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), err)
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "SERVER_ERROR", "Failed to hash password", err)
		return
	}
	user.Password = string(hashedPassword)

	if err := dao.Db.Create(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "DB_ERROR", "Failed to create user", err)
		return
	}

	response.JSON(c, http.StatusCreated, "CREATED", "User registered successfully", nil)
}

func UserLogin(c *gin.Context) {
	user := dao.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "BAD_REQUEST", err.Error(), err)
		return
	}
	if !user.IsValidUser() {
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid username or password", nil)
		return
	}

	// 生成 JWT
	token, err := user.GetJWToken()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "SERVER_ERROR", "Failed to generate token", err)
		return
	}
	response.Success(c, gin.H{"token": token})
}

func Test(c *gin.Context) {
	response.Success(c, gin.H{"message": "test success"})
}
