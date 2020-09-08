package controller

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"spike/service"
	"strconv"
)

type splke struct{}

//instance 实例
var instance splke

// GrabMask 抢口罩
func (*splke) grabMask(ctx *gin.Context) {
	uidStr := ctx.Request.URL.Query().Get("userID")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		writeError(ctx, err)
		return
	}
	err = service.Mask.Buy(uid)
	if err != nil {
		writeError(ctx, err)
		return
	}
	writeResponse(ctx, "抢购完成啦")
	return
}
func writeResponse(ctx *gin.Context, msg string) {
	_, err := ctx.Writer.WriteString(msg)
	if err != nil {
		logger.Errorf("渲染响应失败，因为【%v】", err)
	}
}
func writeError(ctx *gin.Context, serviceErr error) {
	ctx.Writer.WriteHeader(500)
	_, err := ctx.Writer.WriteString(serviceErr.Error())
	if err != nil {
		logger.Errorf("渲染响应失败，因为【%v】", err)
	}
}
