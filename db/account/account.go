/**
*FileName: account
*Create on 2018/12/5 上午5:55
*Create by mok
*/

package account

import (
	"shopadmin/db"
	"shopadmin/model"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

//管理员登录
func CheckLogin(email,password string,loginIP int64)(*model.User,error){
	var u = new(model.User)
	sqlstr := "select uid,username,email from `shop_admin` where email=? and password=?"
	err := db.DB.Get(u,sqlstr,email,password)
	if err != nil{
		return nil,err
	}
	createTime := time.Now().Unix()
	sqlstr = "update `shop_admin` set login_ip=?,login_time=? where email=? and password=?"
	_,err = db.DB.Exec(sqlstr,loginIP,createTime,email,password)
	return u,err
}


//新增管理员
func Insert(user *model.User)error{
	sqlstr := "insert into `shop_admin`(email,username,password,create_time) values(?,?,?,?)"
	_,err := db.DB.Exec(sqlstr,user.Email,user.Username,user.Password,time.Now().Unix())
	return err
}


//检查用户是否存在
func CheckIsExist(email string)error{
	var count int
	sqlstr := "select count(*) from `shop_admin` where email=?"
	if err :=db.DB.Get(&count,sqlstr,email);err != nil{
		return errors.New("查询用户失败")
	}
	if count < 1{
		return errors.New("该用户不存在")
	}
	return nil
}


//修改账号数据比如密码用户名
func Update(uid int,field,data string)error{
	sqlstr := fmt.Sprintf("update `shop_admin` set %s=? where uid=?",field)
	_,err := db.DB.Exec(sqlstr,data,uid)
	return err
}


//获取所有的账户信息
func GetAll()([]*model.User,error){
	var users  = make([]*model.User,0)
	sqlstr := "select uid,email,username,login_ip,login_time,create_time,modify_time from shop_admin"
	err := db.DB.Select(&users,sqlstr)
	return users,err
}


//删除管理员
func DeleteManager(UID uint8)error{
	sqlstr := "delete from `shop_admin` where uid=?"
	_,err := db.DB.Exec(sqlstr,UID)
	return err
}