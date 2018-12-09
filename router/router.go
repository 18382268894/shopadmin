/**
*FileName: router
*Create on 2018/12/6 上午4:19
*Create by mok
*/

package router

import (
	"github.com/gin-gonic/gin"
	"shopadmin/controller/accountController"
	"shopadmin/middleware"
	"shopadmin/controller/index"
	"shopadmin/controller/managerController"
	"shopadmin/controller/category"
	"shopadmin/controller/productController"
)

func LoadRouter(r *gin.Engine){
	r.LoadHTMLGlob("./views/**/*")
	r.Static("/static","./public")
	r.POST("/login", accountController.Login)
	r.GET("/login", accountController.Display)
	r.GET("/seek", accountController.SeekPasswordDisplay)
	r.POST("/seek", accountController.SeekPasswordEmail)


	r.Use(middleware.IsLogin)
	r.GET("/logout", accountController.Logout)
	r.GET("/index",index.Index)

	r.GET("/managers", managerController.ManagerList)
	r.GET("/addmanager", managerController.AddManagerDisPlay)
	r.POST("/addmanager", managerController.AddMnager)
	r.GET("/deletemanager", managerController.DeleteManager)

	r.GET("/categorylist",category.Display)
	r.GET("/addcategory",category.AddCategoryDisplay)
	r.POST("/addcategory",category.AddCategory)
	r.GET("/deletecategory",category.DeleteCategory)
	r.GET("/updatecategory",category.UpdateCategoryDisplay)
	r.POST("/updatecategory",category.UpdateCategory)

	r.GET("/addproduct",productController.CreateProductDisplay)
}