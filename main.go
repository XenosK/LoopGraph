package main

import (
	"fmt"
<<<<<<< HEAD
	"LoopGraph/routers"
)
func main()  {
	fmt.Println("开始了")
	router := routers.InitRouter()
	router.Run(":8070")

}

=======
	_ "LoopGraph/dbutils"
	"LoopGraph/routers"
)

func main()  {

	fmt.Println("开始了")
	router := routers.InitRouter()
	router.Run(":8070")
}
>>>>>>> update laui html and json
