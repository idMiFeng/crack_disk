package dao

import (
	"crackpassword/model"
	"crackpassword/utils"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var err error

func InitDb() {

	dns := "sqlite3.db" // 数据库文件名为 sqlite3.db，存放在当前目录下

	utils.DB, err = gorm.Open(sqlite.Open(dns), &gorm.Config{
		// 关闭外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		panic("无法连接到数据库")
	}

	// 迁移数据表结构
	_ = utils.DB.AutoMigrate(
		&model.PasswordEntry{},
	)

}
