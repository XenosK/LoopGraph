package routers

import (
	"LoopGraph/controllers"
	"github.com/gin-gonic/gin"
)



func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
<<<<<<< HEAD
		"root":    "root",
=======
		"loop":    "LOOP2themoon",
>>>>>>> update laui html and json
	}))

	authorized.GET("/", controllers.Index)
	authorized.GET("/top10", controllers.GetTop10)
<<<<<<< HEAD
	//authorized.GET("/demo", controllers.Demo)
=======
>>>>>>> update laui html and json

	//router.GET("/", controllers.Index)
	//router.GET("/top10", controllers.GetTop10)

	return router
}
<<<<<<< HEAD
=======

>>>>>>> update laui html and json
