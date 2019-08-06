package routers

import (
	"LoopGraph/controllers"
	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.Index)
	router.GET("/top10", controllers.GetTop10)

	return router
}
