/**
*FileName: shopadmin
*Create on 2018/12/5 上午4:29
*Create by mok
*/

package main

import (
	"github.com/gin-gonic/gin"
	"shopadmin/db"
	"shopadmin/pkg/session"
	"shopadmin/router"
)

func main(){
	err := db.Init()
	if err != nil{
		panic(err)
	}
	session.Init()
	r := gin.New()
	router.LoadRouter(r)
	r.Run(":9090")
}