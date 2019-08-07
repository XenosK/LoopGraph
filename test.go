package main

import (
	"LoopGraph/dbutils"
	"LoopGraph/models"
	"fmt"
	"sort"
	"strconv"
)

func main()  {
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
		for i := 0; i < len(content); i++ {
			// 获取content(数组)，中的股票code，日期
			code := content[i].Code
			date := content[i].Time_key[:10]
			change_rate := strconv.FormatFloat(float64(content[i].Change_rate), 'f', 6, 64)
			// 标普500不加入持仓统计
			if code != "US..INX" {
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
		var a int= len(content)
		var ip *int
		ip = &a
		fmt.Println(ip)
		stockRecord["持仓股票统计"][today_date] = strconv.FormatFloat(float64(today_rate/3), 'f', 6, 64)

	}
	var dateList []string

	for k,_ := range stockRecord["US..INX"]{
		if k!="stock"{
			dateList = append(dateList, k)
		}
	}
	dateList = append(dateList, "2019-08-03")
	fmt.Println(dateList)
	sort.Strings(dateList)
	fmt.Println(dateList)

}
