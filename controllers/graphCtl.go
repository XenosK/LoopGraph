package controllers

import (
	"LoopGraph/dbutils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"LoopGraph/models"
	"strconv"
	"encoding/json"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{

	})
}

func totalMath(stockinfo map[string]string) string {
	var totalRate float64
	for k, v := range stockinfo {
		if k != "stock" {
			vf, _ := strconv.ParseFloat(v, 64)
			totalRate+=vf
		}
	}
	totalString := strconv.FormatFloat(float64(totalRate), 'f', 6, 64)
	return totalString
}

func GetTop10(c *gin.Context)  {
	//fmt.Println(c.Query("page"))
	//fmt.Println(c.Query("limit"))
	//realprice := []models.Realprice{}
	//dbutils.DB.Find(&realprice)
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"msg": "",
	//	"count": 20,
	//	"data": realprice,
	//
	//})

	// 存储每只股票的结果
	stockRecord := make(map[string]map[string]string)
	// 初始化持仓记录
	stockRecord["持仓股票统计"] = map[string]string{"stock":"持仓股票统计"}
	// 查询的mysql数据模型
	realprice := []models.Realprice{}
	dbutils.DB.Find(&realprice)
	// mysql中字段content需要的内容
	content := []models.Content{}
	// 遍历查处的所有结果
	for _, con := range realprice {
		// 获取每条结果的content字段，并发序列化json
		err := json.Unmarshal([]byte(con.Content), &content)
		if err!=nil{
			fmt.Println(err)
			break
		}

		today_rate := 0.0
		today_date := ""
		for i:=0; i<len(content); i++ {
			// 获取content(数组)，中的股票code，日期
			code := content[i].Code
			date := content[i].Time_key[:10]
			change_rate := strconv.FormatFloat(float64(content[i].Change_rate), 'f', 6, 64)
			// 标普500不加入持仓统计
			if code!="US..INX" {
				today_rate += content[i].Change_rate
			}
			today_date = date
			if _, ok := stockRecord[code]; ok {

			} else {
				stockRecord[code] = map[string]string{}
				stockRecord[code]["stock"] = code
			}
			stockRecord[code][date] = change_rate

		}
		stockRecord["持仓股票统计"][today_date] = strconv.FormatFloat(float64(today_rate), 'f', 6, 64)

	}
	stockRecord["US..INX"]["stock"] = "标普500指数"
	stockRecord["US..INX"]["total"] = totalMath(stockRecord["US..INX"])
	stockRecord["持仓股票统计"]["total"] = totalMath(stockRecord["持仓股票统计"])

	var stockRecords[]map[string]string
	stockRecords = append(stockRecords, stockRecord["持仓股票统计"])
	stockRecords = append(stockRecords, stockRecord["US..INX"])
	delete(stockRecord, "持仓股票统计")
	delete(stockRecord, "US..INX")

	for k, v := range stockRecord {
		stockRecord[k]["total"] = totalMath(stockRecord[k])
		stockRecords = append(stockRecords, v)
	}
	//fmt.Println(stockRecords)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "",
		"count": 200,
		"data": stockRecords,

	})

}