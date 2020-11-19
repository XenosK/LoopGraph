package routers

import (
	"LoopGraph/controllers"
	"github.com/gin-gonic/gin"
)



func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")

	// html
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"loop":    "LOOP2themoon",
	}))

	authorized.GET("/", controllers.Index)
	authorized.GET("/top10/:name", controllers.GetTop10)
	authorized.GET("/detailslist", controllers.StrategyDetails)
	//authorized.GET("/api/strategy", controllers.StrategyApi)

	////api
	api := router.Group("/api",)
	api.GET("/strategy", controllers.StrategyApi)
	////authorized.GET("/top10/:name", controllers.GetTop10)



	return router
}

