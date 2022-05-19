package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TodoList struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB
var err error

func InitDb() {
	dsn := "root:输入自己的密码@tcp(127.0.0.1:3306)/自己的database?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//禁用默认表名的复数形式
	DB.SingularTable(true)
	//迁移建表
	DB.AutoMigrate(&TodoList{})
	//设置连接池中的最大闲置连接数
	DB.DB().SetMaxOpenConns(10)
}
