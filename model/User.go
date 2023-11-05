package model

import (
	"Vue-Gin-BlogWeb/utils/errmsg"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)www" json:"username"`
	Password string `gorm:"type:varchar(20) " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	fmt.Println("名字是", name)
	fmt.Println("users.Role", users.Role)
	db.Select("id").Where("username = ?", name).First(&users) //查询数据库中是否存在这个用户
	if users.ID > 0 {
		//fmt.Println("users.ID", users.ID)
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err = db.Create(&data).Error //返回一个DB类型的对象去访问里面的Error对象
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error //分页查询
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&User{}).Where("id = ?", id).Updates(maps).Error //Updates 更新
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error //删除数据
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 12, 32, 12, 32, 12, 32}
	//参数说明：password:密码，salt：掩，N：CPU和内存消耗的指数，r：块大小，p：并行化的程度，keyLen：期望的密钥长度，见文档
	HashPw, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw) //base64编码，将Hash后的密码转换成字符串
	return fpw

}
