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
	Name    string `gorm:"column:name"`
	Posts   []Post `gorm:"foreignKey:UserID;references:ID"` //foreignkey是外键Post的UserID，references是引用的主键,User表的ID
	PostNum int    `gorm:column:post_num`
}

type Post struct {
	gorm.Model
	Title            string    `gorm:"column:title"`
	UserID           uint      `gorm:"column:user_id"`
	Comments         []Comment `gorm:"foreignKey:PostID;references:ID"`
	HasCommentStatus string    `gorm:"column:has_comment_status"`
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 在文章创建时自动更新用户的文章数量统计字段。 User.PostNum为文章数
	user := User{}
	tx.First(&user, "id = ?", post.UserID)
	user.PostNum++
	tx.Save(&user)
	return
}

type Comment struct {
	gorm.Model
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
	PostID  uint   `gorm:"column:post_id"`
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	tx.First(&post, "id = ?", comment.PostID)

	commentCount := int64(0)
	tx.Select("count(1)").Where("post_id = ?", post.ID).Count(&commentCount)
	post.HasCommentStatus = func(num int64) string {
		if num > 0 {
			return "有结论"
		}
		return "无结论"
	}(commentCount)
	tx.Save(&post)
	return
}

func (comment *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 在文章创建时自动更新用户的文章数量统计字段。 User.PostNum为文章数
	post := Post{}
	tx.First(&post, "id = ?", comment.PostID)
	post.HasCommentStatus = "有评论"
	tx.Save(&post)
	return
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

func getMaxCommentCountPost() {
	post := Post{}
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	db.
		Table("posts").
		Joins("INNER JOIN (select * from (select count(1) comment_num,post_id from comments group by post_id)tmp order by comment_num desc limit 1) top_tb on top_tb.post_id = posts.id").
		Find(&post)
	fmt.Println(post)
}

func main() {
	init_db()
	/*题目1：模型定义
		q1假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	要求 ：
	使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	编写Go代码，使用Gorm创建这些模型对应的数据库表。
	*/
	createTables()
	insertData()

	/*
		题目2：关联查询
		基于上述博客系统的模型定义。
		要求 ：
		编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
		编写Go代码，使用Gorm查询评论数量最多的文章信息。
	*/
	//queryPostAndCommentViaUser("user-0")
	//getMaxCommentCountPost()

	/*
			题目3：钩子函数
		继续使用博客系统的模型。
		要求 ：
		为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
		为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	*/

}
