package controllers

import (
	"LoopGraph/configs"
	"LoopGraph/dbutils"
	"LoopGraph/models"
	"LoopGraph/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// func Index(c *gin.Context) {
// 	sid := []models.Strategy_details{}
// 	strategyData := []map[string]interface{}{}

// 	// 查数据库，转json（bytes）
// 	dbutils.DB.Find(&sid)
// 	jsons, errs := json.Marshal(sid) //转换成JSON返回的是byte[]
// 	if errs != nil {
// 		fmt.Println(errs.Error())
// 	}

// 	// json 转map
// 	errs1 := json.Unmarshal(jsons, &strategyData)
// 	if errs1 != nil {
// 		fmt.Println(errs1.Error())
// 	}

// 	// 渲染html页面
// 	//c.String(200,"eweqe")
// 	c.HTML(http.StatusOK, "index.html", gin.H{
// 		"dates_rlong":  utils.DateList_rlong,
// 		"dates_flong":  utils.DateList_flong,
// 		"dates_slong":  utils.DateList_slong,
// 		"dates_rshort": utils.DateList_rshort,
// 		"dates_fshort": utils.DateList_fshort,
// 		"dates_sshort": utils.DateList_sshort,
// 		//"strategyData" :strategyData,
// 	})

// }

func Index(c *gin.Context) {
	Categorys := []models.Strategy_details{}
	strategyData := []map[string]interface{}{}
	// strategyData := []interface

	// 查数据库，转json（bytes）
	dbutils.DB.Where("sid > ?", 0).Group("category").Find(&Categorys)
	// dbutils.DB.Table("strategy_details").Where("sid > ?", 0).Select("strategy_details.category").Group("category").Rows()
	jsons, errs := json.Marshal(Categorys) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	// json 转map
	errs1 := json.Unmarshal(jsons, &strategyData)
	if errs1 != nil {
		fmt.Println(errs1.Error())
	}

	// var categorySlice []interface{}
	// for _, x := range strategyData {
	// 	categorySlice = append(categorySlice, x["category"])
	// }

	c.HTML(http.StatusOK, "index.html", gin.H{
		"categorySlice": strategyData,
	})

}

func GetTop10(c *gin.Context) {
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
	case "rlong":
		dbutils.DB.Find(&rlong)
		j, _ := json.Marshal(rlong)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		} // struct转json， json转map
	case "rshort":
		dbutils.DB.Find(&rshort)
		j, _ := json.Marshal(rshort)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		}
	case "slong":
		dbutils.DB.Find(&slong)
		j, _ := json.Marshal(slong)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		}
	case "sshort":
		dbutils.DB.Find(&sshort)
		j, _ := json.Marshal(sshort)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		}
	case "flong":
		dbutils.DB.Find(&flong)
		j, _ := json.Marshal(flong)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		}
	case "fshort":
		dbutils.DB.Find(&fshort)
		j, _ := json.Marshal(fshort)
		err := json.Unmarshal(j, &dbmap)
		if err != nil {
			fmt.Println(err)
		}
		//default:
		//	dbutils.DB.Find(&rlong);j, _ := json.Marshal(rlong); err := json.Unmarshal(j, &m); if err != nil{fmt.Println(err)}
	}

	//// 序列化数据模型,转化的map：m,和传入的name，调用对应的序列化逻辑
	//var stockRecords []map[string]interface{}
	stockRecords := utils.SerializeLongShort(dbmap, name)

	// 分页操作
	count := len(stockRecords)
	sp, bp := utils.Paging(count, page, limit)
	stockRecords = stockRecords[sp:bp]

	//返回json
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "ok",
		"count": count,
		"data":  stockRecords,
	})
}

func StrategyApi(c *gin.Context) {

	sid, _ := strconv.Atoi(c.Query("sid"))
	// sid := c.Query("sid")
	mongourl := configs.Config.Local["mongourl"]
	mongo, err := mgo.Dial(mongourl) // 建立连接
	if err != nil {
		fmt.Println(err)
	}
	client := mongo.DB("strategys").C("loopresult")

	strategy := models.Strategy{}
	cErr := client.Find(bson.M{"sid": int(sid)}).One(&strategy)
	if cErr != nil {
		fmt.Println(cErr)
	}
	// json转map
	jsons, _ := json.Marshal(strategy) //转换成JSON返回的是byte[]

	defer mongo.Close()

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		//"data":  string(jsons),
		"data": json.RawMessage(jsons),
	})
}

func StrategyDetails(c *gin.Context) {
	sid, _ := strconv.Atoi(c.Query("sid"))
	// category, _ := strconv.Atoi(c.Query("category"))
	//c.HTML(http.StatusOK, "strategyDetails.html", gin.H{
	//	"sid": sid,
	//})
	c.HTML(http.StatusOK, "strategyindexMon.html", gin.H{
		"sid": sid,
		// "category": category,
	})
}

func StrategyDetailsBtc(c *gin.Context) {
	code := c.Param("code")
	kline := c.Param("kline")
	cid := c.Param("cid")
	ranger := c.Param("ranger")
	// sid, _ := strconv.Atoi(c.Query("sid"))
	// category, _ := strconv.Atoi(c.Query("category"))
	//c.HTML(http.StatusOK, "strategyDetails.html", gin.H{
	//	"sid": sid,
	//})
	c.HTML(http.StatusOK, "strategyindex.html", gin.H{
		"code":   code,
		"kline":  kline,
		"sid":    cid,
		"ranger": ranger,
		// "category": category,
	})
}

func StrategyIdMap(c *gin.Context) {
	// fmt.Println(c.Query("category"))
	category := c.Query("category")
	fmt.Println(category)
	sid := []models.Strategy_details{}
	strategyData := []map[string]interface{}{}

	// 查数据库，转json（bytes）
	dbutils.DB.Where("category = ?", category).Find(&sid)
	jsons, errs := json.Marshal(sid) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	// json 转map
	errs1 := json.Unmarshal(jsons, &strategyData)
	if errs1 != nil {
		fmt.Println(errs1.Error())
	}

	// 渲染
	c.HTML(http.StatusOK, "loopBack", gin.H{
		//"strategyData": json.RawMessage(jsons), 	//byte[]转换成string 输出
		"strategyData": strategyData,
		//"strategyData": map[string]string{"sid":"1"}, 	//byte[]转换成string 输出
	})
}

func RealTimeLOOP(c *gin.Context) {
	code := c.Query("code")
	kline := c.Query("kline")
	cid := c.Query("cid") // cid：long，short
	ranger := c.Query("ranger")
	var url = "https://eniac.loopbook.cn/backtrader/airun/" + code + "/" + kline + "/" + cid + "/" + ranger
	// var url = "http://eniac.loopbook.cn:31778/backtrader/airun/OKEX.btc/finance/" + cid
	// var url = "http://127.0.0.1:7777/backtrader/airun/OKEX.btc/finance/" + cid
	var loopjson []byte
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	loopjson, _ = ioutil.ReadAll(response.Body)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		//"data":  string(jsons),
		"data": json.RawMessage(loopjson),
	})
	// var loopMap map[string]interface{}
	// errs := json.Unmarshal(loopjson, &loopMap)
	// if errs != nil {
	// 	fmt.Println(errs)
	// }
	// c.JSON(http.StatusOK, gin.H, loopMap)
}
