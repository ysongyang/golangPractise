package controller

import (
	"github.com/gin-gonic/gin"
	"golangPractise/mallApi/dao"
	"golangPractise/mallApi/model"
	"golangPractise/mallApi/utils/response"
)

func InsertNewUser(context *gin.Context) {
	utilGin := response.Gin{Ctx: context}
	var user model.Users
	err := context.ShouldBind(&user)
	if err != nil {
		utilGin.Response(1, err.Error(), "")
		return
	}
	// 写入数据库
	res, err := dao.InsertUser(&user)
	if err != nil {
		utilGin.Response(1, err.Error(), "")
		return
	}
	utilGin.Response(0, "success", res)
}
