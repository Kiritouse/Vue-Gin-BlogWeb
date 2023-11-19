package model

import (
	"Vue-Gin-BlogWeb/utils/errmsg"
	"errors"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //Category int
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` //描述
	Content string `gorm:"type:longtext" json:"content"`  //文章的主体内容
	Img     string `gorm:"type:varchar(100)" json:"img"`  //文章的图片
}

// 新增分类
func CreateArt(data *Article) int {
	err = db.Create(&data).Error //返回一个DB类型的对象去访问里面的Error对象
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的所有文章

//todo 查询单个文章

// todo 查询文章列表
func GetArt(pageSize int, pageNum int) []Category {
	var cate []Category

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error //分页查询
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return cate
}

// 编辑文章信息
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error //Updates 更新
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error //删除数据

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
