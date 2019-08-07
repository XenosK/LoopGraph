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
		"root":    "root",
	}))

	authorized.GET("/", controllers.Index)
	authorized.GET("/top10", controllers.GetTop10)
	//authorized.GET("/demo", controllers.Demo)

	//router.GET("/", controllers.Index)
	//router.GET("/top10", controllers.GetTop10)

	return router
}
