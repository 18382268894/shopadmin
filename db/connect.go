/**
*FileName: db
*Create on 2018/12/5 上午5:56
*Create by mok
 */

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"log"
	"shopadmin/config"
	"shopadmin/model"
	"time"
)

var DB *sqlx.DB
var GDB *gorm.DB

func Init() error {
	err := initDB()
	if err != nil{
		return err
	}
	err = initGDB()
	return err
}

func Close() {
	DB.Close()
}
func initDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.Mysql_User, config.Mysql_Password, config.Mysql_Host, config.Mysql_Port, config.Mysql_DBname)
	DB, _ = sqlx.Open("mysql", dsn)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxLifetime(100 * time.Second)
	err := DB.Ping()
	if err == nil {
		log.Println("数据库连接成功")
	}
	return err
}

func initGDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", config.Mysql_User, config.Mysql_Password, config.Mysql_Host, config.Mysql_Port, config.Mysql_DBname)
	GDB, _ = gorm.Open("mysql", dsn)
	GDB.DB().SetMaxIdleConns(5)
	GDB.DB().SetMaxOpenConns(20)
	GDB.DB().SetConnMaxLifetime(100 * time.Second)
	err := GDB.DB().Ping()
	if err == nil {
		log.Println("数据库连接成功")
		err = createTables()
	}
	GDB.LogMode(true)
	return err
}

func createTables() error {
	//创建
	if !GDB.HasTable(&model.Category{}) {
		return GDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(&model.Category{}).Error
	}

	if !GDB.HasTable(&model.Product{}) {
		err := GDB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(&model.Product{}).Error
		if err == nil{
			GDB.Table("shop_product").
				AddForeignKey("cid","shop_category(cid)","restrict","cascade")
		}
		return err
	}
	return nil
}
