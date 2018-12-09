/**
*FileName: productDB
*Create on 2018/12/9 下午11:26
*Create by mok
*/

package productDB

import (
	"shopadmin/model"
	"shopadmin/db"
)

/*
type Product struct {
	ProductId  uint32 `gorm:"primary_key;type:int(11) unsigned auto_increment comment '商品id'"`
	Cid uint16 `gorm:"type:mediumint(11) unsigned not null comment '商品分类id'"`
	Category Category `gorm:"ForeignKey:Cid"`
	Name string `gorm:"type:varchar(30) comment '商品名称' not null"`
	Summary string `gorm:"type:varchar(255) comment '商品摘要' not null"`
	Num  uint32  `gorm:"type:int unsigned comment '商品库存' not null; default:0"`
	Price float64 `gorm:"type:decimal(10,2) unsigned comment '商品价格' not null"`
	Pics string		`gorm:"type:varchar(500) comment '商品图片存放位置'; default: ''"`
	CreateTime time.Time `gorm:"type:timestamp not null;default:current_timestamp"`
	ModifyTime time.Time `gorm:"type:timestamp on update current_timestamp;omitempty"`
}
*/

//获取单个商品信息
func GetProduct(productid uint32)(*model.Product,error){
	var product = new(model.Product)
	err := db.GDB.Model(&model.Product{}).Find(product).Error
	return product,err
}


//新增商品
func InsertProduct(product *model.Product)(error){
	return db.GDB.Model(&model.Product{}).Create(product).Error
}