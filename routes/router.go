package routes

import (
	"Vue-Gin-BlogWeb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由的入口文件
func InitRouter() { //如果函数名称第一个字母为大写的话就代表是共有的方法
	//如果是小写的话，就代表是私有的方法
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello www.topgoer.com",
			})
		})
	}
	r.Run(utils.HttpPort)

}
