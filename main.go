package main

import (
	_ "LoopGraph/configs"
	_ "LoopGraph/dbutils"
	"LoopGraph/routers"
	"fmt"
)

func main()  {

	fmt.Println("开始了")
	router := routers.InitRouter()
	router.Run(":8070")
}
