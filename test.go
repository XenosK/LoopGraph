package main

//import (
//	"encoding/json"
//	"fmt"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type Student struct {
//	Name   string  `bson: "name"`
//	Age    int     `bson: "age"`
//	Sid    string  `bson: "sid"`
//	Status int     `bson: "status"`
//}
//
//type Per struct {
//	Per   []Student
//}
//
////type Strategy struct {
////	Records      	map[string]interface{}	`json:"records"`
////	Sqn   			map[string]interface{}	`json:"sqn"`
////	TimeReturn  	map[string]interface{}	`json:"timeReturn"`
////	//AnnualReturn    map[string]interface{}	`json:"annualReturn"`
////	AnnualReturn    interface{}	`json:"annualReturn,omitempty"`
////	SharpeRatio     map[string]interface{}	`json:"sharpeRatio"`
////	DrawDown   		map[string]interface{}	`json:"drawDown"`
////	TradeAnalyzer   map[string]interface{}	`json:"tradeAnalyzer"`
////
////	Cashvalue 		float64	`json:"cashvalue"`
////	Portvalue 		float64	`json:"portvalue"`
////	Pnl				float64	`json:"pnl"`
////	Yield			float32	`json:"yield"`
////	Use_time_es		string	`json:"use_time_es"`
////	Use_time_loop	string	`json:"use_time_loop"`
////	Create_time		string	`json:"create_time"`
////	Sid				float32	`json:"strategyId"`
////}
//
//type Strategy struct {
//	Records      	map[string]interface{}	`bson:"records"`
//	Sqn   			map[string]interface{}	`bson:"sqn"`
//	TimeReturn  	map[string]interface{}	`bson:"timeReturn"`
//	AnnualReturn    map[string]interface{}	`bson:"annualReturn"`
//	//AnnualReturn    interface{}	`bson:"annualReturn,omitempty"`
//	SharpeRatio     map[string]interface{}	`bson:"sharpeRatio"`
//	DrawDown   		map[string]interface{}	`bson:"drawDown"`
//	TradeAnalyzer   map[string]interface{}	`bson:"tradeAnalyzer"`
//
//	Cashvalue 		float64	`bson:"cashvalue"`
//	Portvalue 		float64	`bson:"portvalue"`
//	Pnl				float64	`bson:"pnl"`
//	Yield			float32	`bson:"yield"`
//	Use_time_es		string	`bson:"use_time_es"`
//	Use_time_loop	string	`bson:"use_time_loop"`
//	Create_time		string	`bson:"create_time"`
//	Sid				int64 `bson:"sid,omitempty"`
//}
//
//
//func main() {
//
//	mongo, err := mgo.Dial("mongodb://root:dds-uf66@192.168.80.200:3717,192.168.80.201:3717/admin?replicaSet=mgset-10450573") // 建立连接
//	if err!= nil{
//		fmt.Println(err)
//	}
//	//client := mongo.DB("loopLog").C("Demo")
//	client := mongo.DB("strategys").C("loopresult")
//
//	strategy := Strategy{}
//	cErr := client.Find(bson.M{"sid": 33}).One(&strategy)
//	if cErr != nil {
//		fmt.Println(cErr)
//	}
//
//	jsons, _ := json.Marshal(strategy) //转换成JSON返回的是byte[]
//	fmt.Println(strategy)
//	fmt.Println(string(jsons))
//
//	defer mongo.Close()
//}