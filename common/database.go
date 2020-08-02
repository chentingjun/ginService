package common

import (
	"carlos/gin-service/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "47.113.107.36"
	port := "3306"
	database := "personal"
	username := "root"
	password := "root123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	fmt.Println(args)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	DB = db
	db.AutoMigrate(&model.User{})

	return db
}

func GetDB() *gorm.DB {
	return DB
}
