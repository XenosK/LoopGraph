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


<<<<<<< HEAD
func init() {
	//cfg, err := goconfig.LoadConfigFile("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml")
	var conf Config
	err := config.New("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml", &conf)
=======
func initConfig() {
	//cfg, err := goconfig.LoadConfigFile("/usr/local/gopath/src/LoopGraph/configs/yaml/mysql.yaml")
	var conf Config
	err := config.New("yaml/mysql.yaml", &conf)
>>>>>>> update laui html and json
	if err != nil {
		fmt.Println(err.Error())
	}
}
