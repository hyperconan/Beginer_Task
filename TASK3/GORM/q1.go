package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init_db() {
	dsn := "root:123456@tcp(localhost:13306)/metanode?charset=utf8&parseTime=True&loc=Local"
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open db failed, err:", err)
	}

	db = instance
}

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Posts []Post `gorm:"foreignKey:UserID;references:ID"` //foreignkey是外键Post的UserID，references是引用的主键,User表的ID
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"column:title"`
	UserID   uint      `gorm:"column:user_id"`
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"`
}

type Comment struct {
	gorm.Model
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	PostID  uint   `gorm:"column:post_id"`
}

func createTables() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func insertData() {
	tx := db.Begin()
	for i := 0; i < 5; i++ {
		user := User{
			Name: fmt.Sprintf("user-%d", i),
		}
		tx.Create(&user)
		for j := i + 1; j < (i+1)*5; j++ {
			post := Post{
				Title:  fmt.Sprintf("post-%d-%d", i, j),
				UserID: user.ID,
			}
			tx.Create(&post)
			for k := j + 1; k < (j+1)*5; k++ {
				comment := Comment{
					Title:   fmt.Sprintf("comment-%d-%d-%d", i, j, k),
					Content: fmt.Sprintf("content-%d-%d-%d", i, j, k),
					PostID:  post.ID,
				}
				tx.Create(&comment)
			}
		}
	}
	tx.Commit()
}

func queryPostAndCommentViaUser(username string) {
	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//posts := []Post{}
	//db.
	//	Joins("LEFT JOIN users ON posts.user_id = users.id").
	//	Preload("Comments").
	//	Where("users.name = ?", username).
	//	Find(&posts)

	//posts := []Post{}
	user := User{}
	db.
		Preload("Posts.Comments").
		Where("name = ?", username).
		Find(&user)

	for _, post := range user.Posts {
		for _, comment := range post.Comments {
			fmt.Printf("Post Title: %s, Comment Title: %s, Comment Content: %s \n", post.Title, comment.Title, comment.Content)
		}
	}
}

//func getMaxCommentCountPost(){
//	post := Post{}
//	/*
//	select count(1) comment_num,posts.id from posts.id = comments.post_id group by posts.id order by comment_num desc
//	 */
//	db.
//}

func main() {
	init_db()
	//createTables()
	//insertData()
	queryPostAndCommentViaUser("user-0")

}
