package configs
//package configs

import (
	"fmt"
	"github.com/dingdayu/golangtools/config"
)

// 配置
type Config struct {
	Local  map[string]string
	Line   map[string]string
}


func init() {
	//cfg, err := goconfig.LoadConfigFile("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml")
	var conf Config
	err := config.New("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}
