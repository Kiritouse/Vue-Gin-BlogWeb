package v1

import (
	"Vue-Gin-BlogWeb/model"
	"Vue-Gin-BlogWeb/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

var code int

// 添加用户
func AddUser(c *gin.Context) {

	var data model.User
	_ = c.ShouldBindJSON(&data) //将Json数据与结构体绑定
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED //抛出错误
	}
	//接收到的东西返回到前端去
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})

}

//查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {

}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DeleteUser(c *gin.Context) {

}
