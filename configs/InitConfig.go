package configs

import (
	"fmt"
	"github.com/dingdayu/golangtools/config"
)

// 配置,相当于定义了一个类型，如string，所以下面一定要配置这个类型的变量，才可以访问使用
type mysqlConfig struct {
	Local  map[string]string
	Line   map[string]string
}

var Config mysqlConfig

// return 方法
//func GetConfig(m string) string{
//	//cfg, err := goconfig.LoadConfigFile("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml")
//	var conf Config
//	err := config.New("configs/yaml/mysql.yaml", &conf)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	var mysql_url string
//	if m=="line" {
//		mysql_url =  conf.Line["url"]
//	}else{
//		mysql_url =  conf.Local["url"]
//	}
//	fmt.Println(mysql_url)
//	//return mysql_url
//}


// 全局变量
func InitConfig() {
	err := config.New("configs/yaml/mysql.yaml", &Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	InitConfig()
}
