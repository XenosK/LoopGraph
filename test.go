package main

// import (
// 	"LoopGraph/dbutils"
// 	"LoopGraph/models"
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	Categorys := []models.Strategy_details{}
// 	strategyData := []map[string]interface{}{}
// 	// strategyData := []interface

// 	// 查数据库，转json（bytes）
// 	dbutils.DB.Where("sid > ?", 0).Group("category").Find(&Categorys)
// 	// dbutils.DB.Table("strategy_details").Where("sid > ?", 0).Select("strategy_details.category").Group("category").Rows()
// 	jsons, errs := json.Marshal(Categorys) //转换成JSON返回的是byte[]
// 	if errs != nil {
// 		fmt.Println(errs.Error())
// 	}

// 	// json 转map
// 	errs1 := json.Unmarshal(jsons, &strategyData)
// 	if errs1 != nil {
// 		fmt.Println(errs1.Error())
// 	}

// 	var categorySlice []interface{}

// 	for _, x := range strategyData {
// 		categorySlice = append(categorySlice, x["category"])
// 	}

// 	fmt.Println(categorySlice)
// }
