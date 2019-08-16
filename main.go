package main

import (
	"fmt"
	_ "LoopGraph/dbutils"
	"LoopGraph/routers"
	_ "LoopGraph/configs"
)

func main()  {

	fmt.Println("开始了")
	router := routers.InitRouter()
	router.Run(":8070")
}
