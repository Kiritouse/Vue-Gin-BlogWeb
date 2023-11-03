package model

import (
	"Vue-Gin-BlogWeb/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users) //查询数据库中是否存在这个用户
	if users.ID > 0 {
		fmt.Println("users.ID", users.ID)
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	err = db.Create(&data).Error //返回一个DB类型的对象去访问里面的Error对象
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
