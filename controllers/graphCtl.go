package controllers

import (
	"LoopGraph/dbutils"
	"LoopGraph/models"
	"LoopGraph/utils"
	"encoding/json"
	"fmt"
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
	// 渲染html页面
	c.HTML(http.StatusOK, "index.html", gin.H{
		"dates_rlong" : utils.DateList_rlong,
		"dates_flong" : utils.DateList_flong,
		"dates_slong" : utils.DateList_slong,
		"dates_rshort" : utils.DateList_rshort,
		"dates_fshort" : utils.DateList_fshort,
		"dates_sshort" : utils.DateList_sshort,
	})

}



func GetTop10(c *gin.Context)  {
	// 获取api请求参数
	//user := c.MustGet(gin.AuthUserKey).(string)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	name := c.Param("name")

	// db模型
	rlong := []models.Realprice{}
	rshort := []models.Class_2_short_realprice{}
	slong := []models.Sliding_window_20_long_realprice{}
	sshort := []models.Sliding_window_20_short_realprice{}
	flong := []models.Finance_long_realprice{}
	fshort := []models.Finance_short_realprice{}

	// 根据传入的name，查询对应的db，并把结果转成字典作为参数：m
	dbmap := []map[string]interface{}{}
	switch name {
		//case "rlong": dbutils.DB.Find(&rlong);stockRecords = utils.SerializeTop10(rlong)
		case "rlong": dbutils.DB.Find(&rlong);j, _ := json.Marshal(rlong); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}// struct转json， json转map
		case "rshort": dbutils.DB.Find(&rshort);j, _ := json.Marshal(rshort); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}
		case "slong": dbutils.DB.Find(&slong);j, _ := json.Marshal(slong); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}
		case "sshort": dbutils.DB.Find(&sshort);j, _ := json.Marshal(sshort); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}
		case "flong": dbutils.DB.Find(&flong);j, _ := json.Marshal(flong); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}
		case "fshort": dbutils.DB.Find(&fshort);j, _ := json.Marshal(fshort); err := json.Unmarshal(j, &dbmap); if err != nil{fmt.Println(err)}
		//default:
		//	dbutils.DB.Find(&rlong);j, _ := json.Marshal(rlong); err := json.Unmarshal(j, &m); if err != nil{fmt.Println(err)}
	}

	//// 序列化数据模型,转化的map：m,和传入的name，调用对应的序列化逻辑
	//var stockRecords []map[string]interface{}
	stockRecords := utils.SerializeLongShort(dbmap, name)

	// 分页操作
	count := len(stockRecords)
	sp,bp := utils.Paging(count, page, limit)
	stockRecords = stockRecords[sp:bp]

	//返回json
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "ok",
		"count": count,
		"data":  stockRecords,
	})
}