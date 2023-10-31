package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
}
