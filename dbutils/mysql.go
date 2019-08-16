package dbutils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"LoopGraph/configs"
)


var DB *gorm.DB

func InitMysql() {
	conf:=configs.Config{}.Local
	if conf != nil{
		fmt.Println(conf)
	}

	var err error
	DB, err = gorm.Open("mysql", "root:123456@/loop?charset=utf8&parseTime=True&loc=Local")
	//DB, err = gorm.Open("mysql", conf["url"])
	fmt.Println("链接")
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
