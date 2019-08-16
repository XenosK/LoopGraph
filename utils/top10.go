package utils

import (
	"LoopGraph/models"
	"fmt"
	"sort"
	"encoding/json"
)

// 日期存放结果
var DateList []string


// 总收益
func TotalMath(stockinfo map[string]interface{}) float64 {
	var totalRate float64 =1
	for k, v := range stockinfo {
		if k != "stock" {

			if vf, ok := v.(float64);ok{
				totalRate=(1+vf/100)*totalRate
			}

		}
	}
	totalString := float64(totalRate-1)*100
	return totalString
}

//// 总收益
//func TotalMath(stockinfo map[string]string) string {
//	var totalRate float64 =1
//	for k, v := range stockinfo {
//		if k != "stock" {
//
//			vf, _ := strconv.ParseFloat(v, 64)
//			totalRate=(1+vf/100)*totalRate
//		}
//	}
//	totalString := strconv.FormatFloat(float64(totalRate-1)*100, 'f', 6, 64)
//	return totalString
//}

/**
根据value排序
 */
func SortMap(stock_mp map[string]map[string]interface{}) []string {
	//var new_stock_list = make([]string, 0)
	var Stock_list = make([]string, 0)

	for stock_key, _ := range stock_mp {
		Stock_list = append(Stock_list, stock_key)
	}
	sort.Strings(Stock_list)
	//for _, v := range stock_list {
	//	new_stock_list = append(new_stock_list, stock_mp[v])
	//}
	return Stock_list
}


// 日期去重
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


// format mysql数据
func SerializeTop10(realprice []models.Realprice) []map[string]interface{}{
	// 存储每只股票的结果
	stockRecord := make(map[string]map[string]interface{})
	// 初始化持仓记录
	stockRecord["持仓股票统计"] = map[string]interface{}{"stock": "持仓股票统计"}

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
			//change_rate := strconv.FormatFloat(float64(content[i].Change_rate), 'f', 6, 64)
			change_rate := float64(content[i].Change_rate)
			// 标普500不加入持仓统计
			if code != "US..INX" {
				today_rate += content[i].Change_rate
				stock_num+=1
			}
			today_date = date
			if _, ok := stockRecord[code]; ok {

			} else {
				stockRecord[code] = make(map[string]interface{})
				stockRecord[code]["stock"] = code
			}
			stockRecord[code][date] = change_rate

		}
		stockRecord["持仓股票统计"][today_date] = float64(today_rate)/float64(stock_num)

	}

	//var DateList []string
	// 清空日期列表，并重新放入最新的列表
	DateList = DateList[0:0]
	for k,_ := range stockRecord["US..INX"]{
		if k!="stock"{
			DateList = append(DateList, k)
		}
	}
	// 去冲日期列表
	DateList = RemoveRepByLoop(DateList)
	// 排序日期列表
	//sort.Strings(DateList)
	// 倒叙
	sort.Stable(sort.Reverse(sort.StringSlice(DateList)))


	stockRecord["US..INX"]["stock"] = "标普500指数"
	stockRecord["US..INX"]["total"] = TotalMath(stockRecord["US..INX"])
	stockRecord["持仓股票统计"]["total"] = TotalMath(stockRecord["持仓股票统计"])

	// 最终返回结果，先将持仓股票和标普500放入，排列在前
	var stockRecords []map[string]interface{}
	stockRecords = append(stockRecords, stockRecord["持仓股票统计"])
	stockRecords = append(stockRecords, stockRecord["US..INX"])
	delete(stockRecord, "持仓股票统计")
	delete(stockRecord, "US..INX")

	// 对剩余股票排序后，写入最总结果的list
	sortListKey := SortMap(stockRecord)
	// v，股票名称，即stockRecord中排序后的key
	for _, v := range sortListKey{
		// 计算总涨跌幅，存入total字段
		stockRecord[v]["total"] = TotalMath(stockRecord[v])
		stockRecords = append(stockRecords, stockRecord[v])
	}

	return stockRecords
}

