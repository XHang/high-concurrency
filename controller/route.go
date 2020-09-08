package controller

import (
	"github.com/gin-gonic/gin"
)

//Engine 实例
var Engine *gin.Engine

//Init 注册路由
func Init() {
	r := gin.Default()
	r.GET("/grab/mask", instance.grabMask)
	Engine = r
}
