/**
*FileName: model
*Create on 2018/12/5 上午6:16
*Create by mok
*/

package model


type User struct {
	UID uint8 `db:"uid",json:"uid"`
	Email string `db:"email",json:"email"`
	Username string `db:"username",json:"username"`
	Password string `db:"password",json:"password"`
	LoginIP int64 `db:"login_ip",json:"login_ip"`
	LoginTime int64 `db:"login_time",json:"login_time"`
	CreateTime int64 `db:"create_time",json:"create_time"`
	ModifyTime int64 `db:"modify_time".json:"modify_time"`
}