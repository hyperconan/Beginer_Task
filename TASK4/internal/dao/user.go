package dao

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	//存储用户信息，包括 id 、 username 、 password 、 email 等字段。Username添加索引
	gorm.Model
	Username string `json:"username" gorm:"index:idx_username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Posts    []Post `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) IsValidUser() bool {
	// 验证用户账号密码
	var query_user User
	Db.Where("username = ?", u.Username).First(&query_user)
	res := bcrypt.CompareHashAndPassword([]byte(query_user.Password), []byte(u.Password))
	if res == nil {
		*u = query_user
		return true
	}
	return false
}

func (u *User) GetJWToken() (string, error) {
	// 生成JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte("hyperconan"))
	return token, err
}
