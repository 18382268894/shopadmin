/**
*FileName: controller
*Create on 2018/12/5 上午4:34
*Create by mok
*/

package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context){
	c.HTML(http.StatusOK,"index",nil)
}

