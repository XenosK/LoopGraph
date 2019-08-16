package controllers

import (
	"LoopGraph/dbutils"
	"LoopGraph/models"
	"LoopGraph/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func Index(c *gin.Context) {
	// 根据登陆的用户，得到指定信息
	//user := c.MustGet(gin.AuthUserKey).(string)
	//if secret, ok :=secrets[user]; ok {
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"user":secret,
	//	})
	//}
	////else{
	////	c.JSON(http.StatusOK, gin.H{
	////		"user": user, "secret": "NO SECRET :(",
	////	})
	////}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"dates" : utils.DateList,
	})

}



func GetTop10(c *gin.Context)  {
	// 获取api请求参数
	//user := c.MustGet(gin.AuthUserKey).(string)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	// 查询的mysql数据模型
	realprice := []models.Realprice{}
	//dbutils.DB.Find(&realprice).Offset(offset_s).Limit(limit_s)
	dbutils.DB.Find(&realprice)

	// 序列化数据模型
	stockRecords := utils.SerializeTop10(realprice)
	// 分页操作
	count := len(stockRecords)
	small_limit := (page-1)*limit
	big_limit := page*limit
	if small_limit>count{
		small_limit = 0
		big_limit = 0
	} else if big_limit>count{
		big_limit = count
	}
	stockRecords = stockRecords[small_limit:big_limit]

	//返回json
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": count,
		"data":  stockRecords,
	})

}
