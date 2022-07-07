package initialize

import (
	"GinFrame/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitMysqlDB() {

	mysqlInfo := global.GinConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
		mysqlInfo.Port, mysqlInfo.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		color.Red("[InitMysqlDB]初始化失败")
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	db.SingularTable(true)
	//db.LogMode(true) // 设置为true 打印详细sql日志，默认为打印错误sql日志
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	global.GinDB = db
	color.Blue("mysql初始化完成")
}
