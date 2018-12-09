/**
*FileName: manager
*Create on 2018/12/7 下午8:17
*Create by mok
*/

package managerController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopadmin/db/account"
	"log"
	"fmt"
	"shopadmin/model"
	"shopadmin/utils"
	"crypto/md5"
	"strconv"
)


//获取管理员列表
func ManagerList(c *gin.Context){
	users,err := account.GetAll()
	var data  = make(map[string]interface{})
	data["error"] = "获取管理员列表失败"
	if err != nil{
		c.HTML(http.StatusInternalServerError,"managers",data)
		log.Println(err.Error())
		return
	}
	fmt.Println(users[0])
	data["error"] = ""
	data["users"] = users
	c.HTML(http.StatusOK,"managers",data)
	return
}

//添加管理员页面展示
func AddManagerDisPlay(c *gin.Context){
	c.HTML(http.StatusOK,"addmanager",nil)
}

//添加管理员
func AddMnager(c *gin.Context){
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := fmt.Sprintf("%x",md5.Sum([]byte(c.PostForm("password"))))
	//参数检查
	err := utils.CheckParams([]*string{&email,&username,&password},utils.WithTram(),utils.WithNotEmpty())
	err = utils.CheckParams([]*string{&email},utils.WithEmailRule())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	var u  = &model.User{}
	u.Email = email
	u.Password = password
	u.Username = username
		err = account.Insert(u)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"新增用户失败,请稍后再试",
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"创建管理员成功",
	})
	return
}


//删除管理员
func DeleteManager(c *gin.Context){
	idstring := c.Query("uid")
	err := utils.CheckParams([]*string{&idstring},utils.WithTram(),utils.WithNotEmpty())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	UID,err := strconv.ParseUint(idstring,10,8)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"参数格式不正确",
		})
		log.Println(err.Error())
		return
	}
	err = account.DeleteManager(uint8(UID))
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"删除管理员失败，请稍后重试",
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"删除管理员成功",
	})
	return
}