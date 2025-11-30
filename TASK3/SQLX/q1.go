package main

import (
	"fmt"
	"math/rand"

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

func createEmployees() {
	// 有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	schema := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50),
		department VARCHAR(50),
		salary FLOAT
	);
	`
	db.MustExec(schema)
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

func addEmployee(emp Employee) {
	db.MustExec("INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)",
		emp.Name, emp.Department, emp.Salary)
}

func addMultiEmployee() {
	departs := []string{"技术部", "运营部", "财务部"}
	for i := 1; i < 101; i++ {
		departIdx := rand.Intn(3)
		emp := Employee{
			ID:         i,
			Name:       fmt.Sprintf("employee-%d", i),
			Department: departs[departIdx],
			Salary:     rand.Float32() * 10000,
		}
		addEmployee(emp)
	}
}

func getEmployeesFromTech() []Employee {
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	employees := []Employee{}
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部") // 获取多个用select
	if err != nil {
		fmt.Println("select employees from tech department failed, err:", err)
	}
	for _, emp := range employees {
		fmt.Println("ID:", emp.ID, "Name:", emp.Name, "Department:", emp.Department, "Salary:", emp.Salary)
	}
	return employees
}

func getTopSalaryEmployee() Employee {
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	emp := Employee{}
	err := db.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1") // 获取单个用Get
	if err != nil {
		fmt.Println("select top salary employee failed, err:", err)
	}
	return emp
}

func main() {
	initDB()
	//createEmployees()
	//addMultiEmployee()
	getEmployeesFromTech()
	top_emp := getTopSalaryEmployee()
	fmt.Println("Top Salary Employee:", top_emp.Name, "Salary:", top_emp.Salary, "Department:", top_emp.Department, "ID:", top_emp.ID)
}
