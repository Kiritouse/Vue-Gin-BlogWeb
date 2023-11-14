package model

import (
	"Vue-Gin-BlogWeb/utils/errmsg"
	"errors"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"` //在json里绑定name
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`  //在json里绑定id
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate) //查询数据库中是否存在这个用户
	if cate.ID > 0 {
		//fmt.Println("users.ID", users.ID)
		return errmsg.ERROR_CATENAME_USED //3001
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCate(data *Category) int {
	err = db.Create(&data).Error //返回一个DB类型的对象去访问里面的Error对象
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类下的所有文章

// 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error //分页查询
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return cate
}

// 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error //Updates 更新
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error //删除数据

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
