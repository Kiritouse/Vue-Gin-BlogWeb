package model

import "github.com/jinzhu/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //Category int
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` //描述
	Content string `gorm:"type:longtext" json:"content"`  //文章的主体内容
	Img     string `gorm:"type:varchar(100)" json:"img"`  //文章的图片
}
