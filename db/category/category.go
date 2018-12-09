/**
*FileName: category
*Create on 2018/12/9 上午3:31
*Create by mok
*/

package category

import (
	"shopadmin/model"
	"shopadmin/db"
	"errors"
	"fmt"
)




//获取category列表
func CategoryList()([]*model.Category,error){
	err := getAll()
	if err != nil{
		return nil,err
	}
	var dest = &[]*model.Category{}
	getChild(*model.Categories,dest,0,1)
	return *dest,nil
}

//获取单个category
func GetCategory(cid uint16)(*model.Category,error){
	cate := new(model.Category)
	err := db.GDB.Where("cid=?",cid).Find(cate).Error
	return cate,err
}

//添加category
func AddCategory(category *model.Category)error{
	if category == nil{
		return errors.New("category cannot be nil")
	}
	err := db.GDB.Create(category).Error
	if err != nil{
		return err
	}
	*model.Categories = append(*model.Categories,category)
	return nil
}


//删除category
func DeleteCategory(cid uint16)error{
	var deletelist = &[]*model.Category{}
	fmt.Println(cid)
	getChild(*model.Categories,deletelist,cid,0)
	*deletelist = append(*deletelist,&model.Category{Cid:cid})
	var cids = make([]uint16,0)
	for _,v := range *deletelist{
		cids = append(cids,v.Cid)
	}
	return db.GDB.Model(&model.Category{}).Where("cid in (?)",cids).Delete(&model.Category{}).Error
}


//更新category
func UpdateCategory(category *model.Category)error{
	return db.GDB.Model(&model.Category{}).Update(category).Error
}


//获取所有的category分类数据
func getAll()(error){
	err := db.GDB.Model(&model.Category{}).Find(model.Categories).Error
	return err
}

//获取pid下的子节点,并重新组合
func getChild(categories []*model.Category,dest *[]*model.Category,pid uint16,level uint16){
	for _,category := range categories {
		if category.Pid == pid {
			category.Level = level
			*dest = append(*dest, category)
			getChild(categories, dest, category.Cid, level+1)
		}
	}
}
