package orms

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	//存储用户信息，包括 id 、 username 、 password 、 email 等字段。Username添加索引
	gorm.Model
	Username string `gorm:"index:idx_username"`
	Password string
	Email    string
	Posts    []Post `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) IsValidUser() bool {
	var hashPwd string
	
	// 验证用户账号密码
	Db.Where("username = ? AND password = ?", u.Username, u.Password).First(&u)
	return u.ID != 0
}

func (u *User) GetJWToken() string {
	// 生成JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"sub":      fmt.Sprintf("%s", u.Password[:5]),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte("hyperconan"))
	if err != nil {
		fmt.Println("generate token failed, err:", err)
	}
	return token
}
