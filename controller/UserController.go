package UserController

import (
	"carlos/gin-service/common"
	"carlos/gin-service/model"
	"carlos/gin-service/utils"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if IsTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "用户信息已存在"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	// 创建用户
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功"})
}


// IsTelephoneExist 判断手机号是否存在
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

