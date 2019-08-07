package controllers

import (
	"LoopGraph/dbutils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"LoopGraph/models"
	"sort"
	"strconv"
	"encoding/json"
)


//var secrets = gin.H{
//	"loop":    gin.H{"email": "foo@bar.com", "phone": "123433"},
//}


var DateList []string

func Index(c *gin.Context) {
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
		"dates" : DateList,
	})

}

func totalMath(stockinfo map[string]string) string {
	var totalRate float64 =1
	for k, v := range stockinfo {
		if k != "stock" {

			vf, _ := strconv.ParseFloat(v, 64)
			totalRate=(1+vf/100)*totalRate
		}
	}
	totalString := strconv.FormatFloat(float64(totalRate-1)*100, 'f', 6, 64)
	return totalString
}

func RemoveRepByLoop(slc []string) []string {
	result := []string{}  // 存放结果
	for i := range slc{
		flag := true
		for j := range result{
			if slc[i] == result[j] {
				flag = false  // 存在重复元素，标识为false
				break
			}
		}
		if flag {  // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
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
	// 获取登陆的用户
	//user := c.MustGet(gin.AuthUserKey).(string)


	// 存储每只股票的结果
	stockRecord := make(map[string]map[string]string)
	// 初始化持仓记录
	stockRecord["持仓股票统计"] = map[string]string{"stock": "持仓股票统计"}
	// 查询的mysql数据模型
	realprice := []models.Realprice{}
	dbutils.DB.Find(&realprice)
	// mysql中字段content需要的内容
	content := []models.Content{}
	// 遍历查处的所有结果
	for _, con := range realprice {
		// 获取每条结果的content字段，并发序列化json
		err := json.Unmarshal([]byte(con.Content), &content)
		if err != nil {
			fmt.Println(err)
			break
		}

		today_rate := 0.0
		today_date := ""
		stock_num := 0
		for i := 0; i < len(content); i++ {
			// 获取content(数组)，中的股票code，日期
			code := content[i].Code
			date := content[i].Time_key[:10]
			change_rate := strconv.FormatFloat(float64(content[i].Change_rate), 'f', 6, 64)
			// 标普500不加入持仓统计
			if code != "US..INX" {
				today_rate += content[i].Change_rate
				stock_num+=1
			}
			today_date = date
			if _, ok := stockRecord[code]; ok {

			} else {
				stockRecord[code] = map[string]string{}
				stockRecord[code]["stock"] = code
			}
			stockRecord[code][date] = change_rate

		}
		stockRecord["持仓股票统计"][today_date] = strconv.FormatFloat(float64(today_rate)/float64(stock_num), 'f', 6, 64)

	}
	//var DateList []string
	DateList = DateList[0:0]

	for k,_ := range stockRecord["US..INX"]{
		if k!="stock"{
			DateList = append(DateList, k)
		}
	}
	DateList = RemoveRepByLoop(DateList)
	sort.Strings(DateList)

	stockRecord["US..INX"]["stock"] = "标普500指数"
	stockRecord["US..INX"]["total"] = totalMath(stockRecord["US..INX"])
	stockRecord["持仓股票统计"]["total"] = totalMath(stockRecord["持仓股票统计"])

	var stockRecords []map[string]string
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
		"code":  0,
		"msg":   "",
		"count": 200,
		"data":  stockRecords,
	})

}

//func Demo(c *gin.Context) {
//	c.HTML(http.StatusOK, "demo.html", gin.H{
//		"slice": DateList,
//		//"slice": []string{"1","2"},
//	})
//}