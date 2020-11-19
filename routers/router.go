package routers

import (
	"LoopGraph/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.LoadHTMLFiles("views/*")
	router.LoadHTMLGlob("views/*")
	router.Static("/static", "./static")

	// html
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"loop": "LOOP2themoon",
	}))
	// 每日数据
	authorized.GET("/", controllers.Index)
	//mysql top10 每日
	authorized.GET("/top10/:name", controllers.GetTop10)
	// 走势图
	authorized.GET("/details", controllers.StrategyDetails)
	authorized.GET("/details/:code/:kline/:cid/:ranger", controllers.StrategyDetailsBtc)
	authorized.GET("/strategys", controllers.StrategyIdMap) // 直接走实时接口回测
	//authorized.GET("/api/strategy", controllers.StrategyApi)

	////api
	api := router.Group("/api")
	// api.GET("/strategy", controllers.StrategyApi) // 走mongo查询
	api.GET("/loop", controllers.RealTimeLOOP) // 直接走实时接口回测

	return router
}
