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
		"loop":    "LOOP2themoon",
	}))

	authorized.GET("/", controllers.Index)
	authorized.GET("/top10", controllers.GetTop10)

	//router.GET("/", controllers.Index)
	//router.GET("/top10", controllers.GetTop10)

	return router
}

