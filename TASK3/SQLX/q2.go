package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() {
	instance, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:13306)/metanode")
	if err != nil {
		fmt.Println("connect db failed, err:", err)
		return
	}
	db = instance
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}

func createBooks() {
	schema := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		Title VARCHAR(50),
		Author VARCHAR(50),
		Price FLOAT
	);
	`
	db.MustExec(schema)
}

type Book struct {
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func addBook(b Book) {
	db.MustExec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)",
		b.Title, b.Author, b.Price)
}

func addMultiBooks() {
	books := []Book{
		{Title: "Book-1", Author: "Author-1", Price: 100},
		{Title: "Book-2", Author: "Author-2", Price: 200},
		{Title: "Book-3", Author: "Author-3", Price: 300},
	}
	for _, b := range books {
		addBook(b)
	}
}

func getTargetPriceBooks(price float32) []Book {
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	books := []Book{}
	err := db.Select(&books, "SELECT * FROM books WHERE price >= ?", price)
	if err != nil {
		fmt.Println("select books failed, err:", err)
	}
	return books
}

func main() {
	initDB()
	//createBooks()
	//addMultiBooks()
	books := getTargetPriceBooks(50)
	for _, book := range books {
		fmt.Println("book:", book)
	}
}
