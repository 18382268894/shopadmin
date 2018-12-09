/**
*FileName: db
*Create on 2018/12/7 下午10:01
*Create by mok
*/

package db


type errParser interface {
	error
	Parse(error)bool
}

