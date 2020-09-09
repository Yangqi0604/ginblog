package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *gorm.DB
var err error
func InitDb()  {
	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
	if err !=nil {
		fmt.Printf("数据库连接失败，请检查参数：",err)
	}
	//禁用默认表名的复数形式
	db.SingularTable(true)

	db.AutoMigrate(&User{},&Article{},&Category{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}