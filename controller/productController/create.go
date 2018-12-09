/**
*FileName: productController
*Create on 2018/12/10 上午1:09
*Create by mok
*/

package productController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopadmin/db/category"
)

//create展示界面
func CreateProductDisplay(c *gin.Context){
	//获取分类数据
	categories ,err :=category.CategoryList()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK,"addproduct",gin.H{
		"categories":categories,
	})
	return
}

/*//create商品
func CreateProduct(c *gin.Context){
	cidstr := c.PostForm("cid")
	name := c.PostForm("name")
	summary :=c.PostForm("summary")
	num := c.PostForm("num")


}*/