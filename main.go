package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User 用户信息 struct
type User struct {
	gorm.Model
	Name      string `gorm:"type: varchar(20); not null"`
	Telephone string `gorm:"type: varchar(11); not null unique"`
	Password  string `gorm:"size: 255; not null"`
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "手机号必须11位"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "密码不能少于6位"})
			return
		}
		// 如果名称没有传，给一个随机的字符中
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		// 判断手机号是否存在
		if IsTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "用户信息已存在"})
			return
		}
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		// 创建用户
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功"})

	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}

// RandomString 生成随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfhjklqwertyuiopzxcvnmASDFJKLQWERTYUIOZXVNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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

	db.AutoMigrate(&User{})

	return db
}

// IsTelephoneExist 判断手机号是否存在
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
