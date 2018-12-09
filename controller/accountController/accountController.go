/**
*FileName: controller
*Create on 2018/12/5 上午5:21
*Create by mok
*/

package accountController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"shopadmin/db/account"
	"shopadmin/pkg/session"
	"fmt"
	"crypto/md5"
	"shopadmin/utils"
	"time"
	"html/template"
)

func Display(c *gin.Context){
	c.HTML(http.StatusOK,"login",nil)
}

//登录验证
func Login(c *gin.Context){
	//参数检查
	email := c.PostForm("email")
	password := c.PostForm("password")
	//检测邮箱密码是否为空
	err := utils.CheckParams([]*string{&email,&password},utils.WithTram(),utils.WithNotEmpty())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	//检查邮箱格式是否正确
	err = utils.CheckParams([]*string{&email},utils.WithEmailRule())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	//登录检查和数据更新
	loginIP,_:= utils.IPToInt(c.ClientIP())
	password = fmt.Sprintf("%x",md5.Sum([]byte(password)))
	u,err := account.CheckLogin(email,password,int64(loginIP))
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":err.Error(),
		})
		log.Fatalln(err.Error())
		return
	}
	//保存登录信息到session
	var s session.Session
	if s,err = session.Mgr.CreateSession();err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	s.Set("username",u.Username)
	s.Set("email",u.Email)
	s.Set("uid",u.UID)
	c.SetCookie("login_session",s.ID(),3600,"/","",false,true)
	c.Redirect(http.StatusSeeOther,"/index")
}

//退出
func Logout(c * gin.Context){
	id := c.MustGet("sessionid")
	session.Mgr.DeleteSession(id.(string))
	time.Sleep(3*time.Second)
	c.Redirect(http.StatusSeeOther,"/login")
	return
}

//找回密码界面
func SeekPasswordDisplay(c *gin.Context){
	c.HTML(http.StatusOK,"seek",nil)
}

//找回密码发送邮件
func SeekPasswordEmail(c *gin.Context){
	email := c.PostForm("email")
	//验证参数是否合法
	err := utils.CheckParams([]*string{&email},utils.WithTram(),utils.WithNotEmpty(),utils.WithEmailRule())
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	//查询数据库中是否存在该email的用户
	if err = account.CheckIsExist(email);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}

	//生成token用于邮件验证链接
	var data = make(map[string]interface{})
	data["email"] = email
	tokenss,err := utils.NewToken(data)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"找回密码失败，请稍后重试",
		})
		log.Println(err.Error())
		return
	}

	//存在该用户就发送邮件
	tp,err := template.ParseFiles("./views/account/seekemail.html")
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"找回密码失败，请稍后重试",
		})
		log.Println(err.Error())
		return
	}
	var conf = utils.EmailConfig{
		From:"电子商城",
		To:[]string{email},
		Subject:"找回密码",
		Tp:tp,
		Messages:map[string]interface{}{
			"token":tokenss,
		},
	}
	err = utils.SendEmail(&conf)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"发送邮件失败，请稍后重试",
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"我们已经给你的邮箱发送了邮件，请尽快进行点击验证",
	})
	return
}



