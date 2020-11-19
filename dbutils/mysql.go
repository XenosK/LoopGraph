package dbutils

import (
	"LoopGraph/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var DB *gorm.DB

func InitMysql() {
	//mysqlurl := configs.GetConfig("local") //写法是启动程序时候，初始化config
	mysqlurl := configs.Config.Line["mysqlurl"]
	//mysqlurl := configs.Config.Line["url"]
	var err error
	DB, err = gorm.Open("mysql", mysqlurl)

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	// 去除表后面复数
	DB.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return "b3_pipe_" + defaultTableName
		return defaultTableName
	}


}


func init()  {
	InitMysql()
}
