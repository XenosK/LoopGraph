package main

import (
	"fmt"
	"LoopGraph/routers"
)
func main()  {
	fmt.Println("开始了")
	router := routers.InitRouter()
	router.Run(":8070")

}
