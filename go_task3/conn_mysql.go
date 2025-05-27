package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

var db *gorm.DB

// 在init函数中初始化数据库连接
func init() {
	fmt.Println("init")
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gorm"
	timeout := "10s"

	// dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 连接MySQL,获取通用数据库对象db
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println("连接数据库成功,DB=", DB)
	db = DB
	//创建students表
	//err = db.AutoMigrate(&Student{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Student表创建成功！")
}

func main() {
	//task1()
	//err := task2()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//task3()
	//task4()
	//task5()
	//task6()
	//task7()
}

/*
题目1：基本CRUD操作,
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。,
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。,
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。,
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func task1() {
	//db.Exec('INSERT INTO students (`name`, age, grade) VALUES (\"张三\", 20, \"三年级\")')
	//students := []Student{}
	//db.Raw("SELECT * FROM students WHERE age > 18").Scan(&students)
	//fmt.Println(students)
	//db.Exec("UPDATE students SET grade = \"四年级\" WHERE `name` = \"张三\"")
	//db.Exec("DELETE FROM students WHERE age < 15")
}

/*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func task2() error {
	var fromID, toID, amount uint = 1, 2, 100
	// 开始事务（注意：GORM 的 Begin 不需要显式错误检查）
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 统一错误处理（包括 panic 和普通错误）
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务回滚（panic）: %v", r)
		}
	}()

	// 1. 查询并锁定转出账户（FOR UPDATE）
	var fromAccount Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&fromAccount, fromID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %v", err)
	}

	// 2. 检查余额
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("余额不足")
	}

	// 3. 扣减转出账户余额（使用 Update 避免重复查询）
	if err := tx.Model(&Account{}).
		Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣减余额失败: %v", err)
	}

	// 4. 增加转入账户余额（直接更新，无需查询）
	if err := tx.Model(&Account{}).
		Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加余额失败: %v", err)
	}

	// 5. 记录交易
	transaction := Transaction{
		FromAccount: fromID,
		ToAccount:   toID,
		Amount:      amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil

}

// 编写Go代码，查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
// 编写Go代码，查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func task3() {
	employees := []Employee{}
	if err := db.Where("department = ?", "技术部").
		Find(&employees).Scan(&employees).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(employees)
	employee := Employee{}
	if err := db.Order("salary desc").
		First(&employee).Scan(&employee).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(employee)
}

func task4() {
	// 查询价格大于50元的书籍
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ?"
	err := db.Raw(query, 50).Scan(&books).Error
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	// 打印结果
	fmt.Println("价格大于50元的书籍：")
	for _, book := range books {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %d\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}

// 模型定义
func task5() {
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}

var mostCommentedPost struct {
	Post
	CommentCount int
	AuthorName   string
}

// 关联查询
func task6() {
	//查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Posts.Comments").First(&user, 4)
	fmt.Println(user)
	//查询评论最多的文章。
	db.Raw(`SELECT p.*, u.username as author_name, COUNT(c.id) as comment_count
		FROM posts p LEFT JOIN comments c ON p.id = c.post_id JOIN users u ON p.author_id = u.id
		GROUP BY p.id ORDER BY comment_count DESC LIMIT 1`).Scan(&mostCommentedPost)

	fmt.Printf("评论最多的文章: %s (作者: %s, 评论数: %d)\n",
		mostCommentedPost.Title, mostCommentedPost.AuthorName, mostCommentedPost.CommentCount)
}

// 钩子函数
func task7() {
	// 创建测试用户
	user := User{
		Username: "testUser",
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(&user)

	// 创建文章（会自动触发AfterCreate钩子）
	post := Post{
		Title:    "测试文章",
		Content:  "这是一篇测试文章的内容",
		AuthorID: user.ID,
	}
	db.Create(&post)

	// 创建评论
	comment := Comment{
		Content: "这是一条测试评论",
		PostID:  post.ID,
		UserID:  user.ID,
	}
	db.Create(&comment)

	// 删除评论（会自动触发AfterDelete钩子）
	db.Delete(&comment)

	// 验证结果
	var updatedUser User
	db.First(&updatedUser, user.ID)
	fmt.Printf("用户文章数: %d\n", updatedUser.PostsCount) // 应为1

	var updatedPost Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("文章评论数: %d, 状态: %s\n",
		updatedPost.CommentsNum, updatedPost.CommentStatus) // 应为0和"无评论"
}
