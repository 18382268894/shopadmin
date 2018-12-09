/**
*FileName: category
*Create on 2018/12/9 上午3:26
*Create by mok
*/

package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopadmin/db/category"
	"strconv"
	"shopadmin/utils"
	"shopadmin/model"
)

//分类展示
func Display(c *gin.Context){
	categories,err := category.CategoryList()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK,"categorylist",gin.H{
		"categories":categories,
	})
	return
}

//添加分类
func AddCategory(c *gin.Context){
	pidstr := c.PostForm("pid")
	title := c.PostForm("title")
	err := utils.CheckParams([]*string{&title,&pidstr},utils.WithNotEmpty(),utils.WithTram())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	pid,err := strconv.ParseUint(pidstr,10,16)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	var cate = &model.Category{
		Pid:uint16(pid),
		Title:title,
	}
	if err = category.AddCategory(cate);err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther,"/addcategory")
	return
}

//添加分类页面展示
func AddCategoryDisplay(c *gin.Context){
	categories,err := category.CategoryList()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	/*for _,cate := range categories{
		fmt.Println(*cate)
	}*/
	c.HTML(http.StatusOK,"addcategory",gin.H{
		"categories":categories,
	})
	return
}


//删除分类
func DeleteCategory(c *gin.Context){
	cidstr := c.Query("cid")
	cid,err := strconv.ParseUint(cidstr,10,16)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	err  = category.DeleteCategory(uint16(cid))
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther,"/categorylist")
}

//编辑分类页面展示
func UpdateCategoryDisplay(c *gin.Context){
	cidstr := c.Query("cid")
	cid,err := strconv.ParseUint(cidstr,10,16)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	categories,err := category.CategoryList()
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}
	cate,_ := category.GetCategory(uint16(cid))
	c.HTML(http.StatusOK,"updatecategory",gin.H{
		"categories":categories,
		"current_category":cate,
	})
	return
}

//修改分类
func UpdateCategory(c *gin.Context){
	cidstr := c.PostForm("cid")
	pidstr := c.PostForm("pid")
	title := c.PostForm("title")
	err := utils.CheckParams([]*string{&title,&pidstr,&cidstr},utils.WithNotEmpty(),utils.WithTram())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	cid,err := strconv.ParseUint(cidstr,10,16)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	pid,err := strconv.ParseUint(pidstr,10,16)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	err = category.UpdateCategory(&model.Category{Cid:uint16(cid),Pid:uint16(pid),Title:title})
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther,"/categorylist")
	return
}