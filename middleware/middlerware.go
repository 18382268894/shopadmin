/**
*FileName: middleware
*Create on 2018/12/6 上午4:11
*Create by mok
*/

package middleware

import (
	"github.com/gin-gonic/gin"
	"shopadmin/pkg/session"
	"net/http"
	"shopadmin/utils"
)

//登录验证中间件
func IsLogin(c *gin.Context){

	id,_ := c.Cookie("login_session")
	err := utils.CheckParams([]*string{&id},utils.WithNotEmpty())
	if err != nil{
		c.Redirect(http.StatusSeeOther,"/login")
		c.Abort()
		return
	}
	s,err  := session.Mgr.GetSession(id)
	if err != nil{
		c.Redirect(http.StatusSeeOther,"/login")
		c.Abort()
		return
	}
	c.Set("sessionid",id)
	c.Set("uid",s.MustGet("uid"))
	c.Set("username",s.MustGet("username"))
	c.Set("email",s.MustGet("email"))
	c.Next()
}