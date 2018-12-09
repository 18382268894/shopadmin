/**
*FileName: model
*Create on 2018/12/9 上午3:31
*Create by mok
*/

package model

import "time"

var Categories = &[]*Category{}

type Category struct {
	Cid uint16  `gorm:"primary_key;type:mediumint(11) unsigned AUTO_INCREMENT comment '分类ID';"`
	Pid uint16  `gorm:"type:mediumint(11) unsigned;not null;"`
	Title string `gorm:"type:varchar(30);not null;unique"`
	CreateTime time.Time `gorm:"type:timestamp;not	null;default:CURRENT_TIMESTAMP;"`
	Level uint16 `gorm:"-"`
}

func (c *Category)TableName()string{
	return "shop_category"
}
