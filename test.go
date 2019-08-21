package main
//
//import (
//	"LoopGraph/dbutils"
//	"encoding/json"
//	"fmt"
//	"sort"
//)
//
////type name struct {
////	rlong []models.Realprice
////	rshort []models.Class_2_short_realprice
////
////}
//
//type name struct {
//	a string
//	b string
//
//}
//
//func ceshi (args ...interface{}) {
//
//	x := name{"12","er"}
//	a := args[0]
//	dbutils.DB.Find(&a)
//
//	fmt.Println(x)
//	jsons, errs := json.Marshal(x)
//	fmt.Println(errs)
//	fmt.Println(jsons)
//	fmt.Println(string(jsons))
//}
//
//var DateList_rlong []string
//var DateList_flong []string
//var DateList_slong []string
//var DateList_rshort []string
//var DateList_fshort []string
//var DateList_sshort []string
//
//
//func DataStatistics(name string, inx map[string]interface{} ){
//	// 清空日期列表，并重新放入最新的列表
//	dateList := []string{}
//	for k,_ := range inx{
//		if k!="stock"{
//			dateList = append(dateList, k)
//		}
//	}
//	// 去重日期列表
//	//dateList = RemoveRepByLoop(dateList)
//	// 倒叙排序日期列表
//	//sort.Strings(DateList)
//	sort.Stable(sort.Reverse(sort.StringSlice(dateList)))
//
//	switch name {
//	case "rlong": DateList_rlong = dateList
//	case "rshort": DateList_rshort = dateList
//	case "slong": DateList_slong = dateList
//	case "sshort": DateList_sshort = dateList
//	case "flong": DateList_flong = dateList
//	case "fshort": DateList_fshort = dateList
//	}
//
//}
//
//func main() {
//	//m := []map[string]interface{}{}
//	//mm := []map[string]interface{}{}
//	//
//	//rlong := []models.Realprice{}
//	//dbutils.DB.Find(&rlong)
//	////for con, _ := range rlong{
//	////	j, _ := json.Marshal(con)
//	////	json.Unmarshal(j, &m)
//	////	fmt.Println(m)
//	////}
//	//j, _ := json.Marshal(rlong)
//	//a :=json.Unmarshal(j, &m)
//	//fmt.Println(a)
//	//b := m[0]["content"]
//	////var data []byte = []byte(b)
//	////fmt.Println(m[0]["content"])
//	//fmt.Println(reflect.TypeOf(b))
//	//result, ok := b.(string)
//	//if !ok {
//	//	fmt.Println("not ok")
//	//}else {
//	//	data := []byte(result)
//	//	c :=json.Unmarshal(data, &mm)
//	//	fmt.Println(c)
//	//	fmt.Println(mm)
//	//}
//
//	st := map[string]interface{}{"2019-08-02":-0.7282739927037866,"2019-08-05":-2.977782000298592,"stock":"标普500指数","total":-3.684369581134783}
//	//DataStatistics(DateList_rlong, st)
//	DataStatistics("rlong", st)
//	fmt.Println(DateList_rlong)
//
//
//
//}