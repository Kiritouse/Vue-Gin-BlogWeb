package model

import (
	"Vue-Gin-BlogWeb/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
) //这里是引入mysql的驱动，但是我们没有直接使用，所以前面加了一个下划线

// 数据库的入口文件
var db *gorm.DB //调用gorm的DB方法,指向gorm.DB结构体，也就是指向数据库
var err error

func InitDb() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)

	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}

	//自动迁移, AutoMigrate 会创建表、缺失的外键、约束、列和索引。 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。 出于保护您数据的目的，它 不会 删除未使用的列
	err := db.AutoMigrate(&User{}, &Article{}, &Category{})
	if err != nil {
		fmt.Println("自动迁移失败，请检查参数:", err)
		return
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	//或许这里应该有close函数？但是似乎新版gorm里面没有了，到时候去查询一下

}
