package routes

import (
	v1 "Vue-Gin-BlogWeb/api/v1"
	"Vue-Gin-BlogWeb/utils"
	"github.com/gin-gonic/gin"
)

// 路由的入口文件
func InitRouter() { //如果函数名称第一个字母为大写的话就代表是共有的方法
	//如果是小写的话，就代表是私有的方法
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块的路由接口,定义了POST请求和GET请求,PUT请求和DELETE请求
		router.POST("user/add", v1.AddUser)      //向服务器提交数据
		router.GET("users", v1.GetUsers)         //查询用户列表
		router.PUT("user/:id", v1.EditUser)      //编辑用户
		router.DELETE("user/:id", v1.DeleteUser) //
		//分类模块的路由接口

		//文章模块的路由接口
		router.POST("category/add", v1.AddCategory)  //向服务器提交数据
		router.GET("category", v1.GetCate)           //查询分类列表
		router.PUT("category/:id", v1.EditCate)      //编辑分类
		router.DELETE("category/:id", v1.DeleteCate) //
	}
	r.Run(utils.HttpPort)

}
