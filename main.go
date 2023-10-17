package main

import (
	"Vue-Gin-BlogWeb/model"
	"Vue-Gin-BlogWeb/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
