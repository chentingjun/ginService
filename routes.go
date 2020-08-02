package main

import (
	UserController "carlos/gin-service/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register", UserController.Register)
	r.POST("/api/auth/login", UserController.Login)
	return r
}
