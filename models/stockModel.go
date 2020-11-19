package models

import "time"

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" structs:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" structs:"deletedAt"`
}

// 做多top9bar10
type LongShort struct {
	//Model

	Content     string    `json:"content" structs:"content"`
	Id          string    `json:"id" structs:"id"`
	Uptime      time.Time `json:"uptime" structs:"uptime"`
	Create_time string    `json:"create_time" structs:"create_time"`
}

// 做多top9bar10
type Realprice struct {
	//Model
	LongShort
}

// 滑动窗口做多
type Sliding_window_20_long_realprice struct {
	//Model
	LongShort
}

//滑动窗口做空
type Sliding_window_20_short_realprice struct {
	//Model
	LongShort
}

//财务做多
type Finance_long_realprice struct {
	//Model
	LongShort
}

//财务做空
type Finance_short_realprice struct {
	//Model
	LongShort
}

// 二分类做空
type Class_2_short_realprice struct {
	//Model
	LongShort
}

//[{'code': 'US.DLR', 'time_key': '2019-08-02 00:00:00', 'open': 115.55999755859375, 'close': 117.21, 'high': 117.93000030517578, 'low': 115.37999725341797, 'pe_ratio': 0.0, 'turnover_rate': 0.4459999967366457, 'volume': 929010, 'turnover': 108765000.153, 'change_rate': 1.454168028154467, 'last_close': 115.52999877929688}

type Content struct {
	//Model

	Code        string  `json:"code" structs:"code"`
	Change_rate float64 `json:"change_rate" structs:"change_rate"`
	Time_key    string  `json:"time_key" structs:"time_key"`
}

type Strategy struct {
	//Records      	map[string]interface{}	`bson:"records,omitempty" json:"records"`
	Sqn                 map[string]interface{} `bson:"sqn,omitempty" json:"sqn"`
	TimeReturn          map[string]interface{} `bson:"timeReturn,omitempty" json:"timeReturn"`
	AnnualReturn        map[string]interface{} `bson:"annualReturn,omitempty" json:"annualReturn"`
	SharpeRatio         map[string]interface{} `bson:"sharpeRatio,omitempty" json:"sharpeRatio"`
	DrawDown            map[string]interface{} `bson:"drawDown,omitempty" json:"drawDown"`
	TradeAnalyzer       map[string]interface{} `bson:"tradeAnalyzer,omitempty" json:"tradeAnalyzer"`
	TransactionRecordVO []interface{}          `bson:"transactionRecordVO,omitempty" json:"transactionRecordVO"`

	StartDay      string  `bson:"startDay,omitempty" json:"startDay"`
	EndDay        string  `bson:"endDay,omitempty" json:"endDay"`
	StartCash     float64 `bson:"startCash,omitempty" json:"startCash"`
	Cashvalue     float64 `bson:"cashvalue,omitempty" json:"cashvalue"`
	Portvalue     float64 `bson:"portvalue,omitempty" json:"portvalue"`
	Pnl           float64 `bson:"pnl,omitempty" json:"pnl"`
	Yield         float64 `bson:"yield,omitempty" json:"yield"`
	Won           float64 `bson:"won,omitempty" json:"won"`
	Use_time_es   string  `bson:"use_time_es,omitempty" json:"use_time_es"`
	Use_time_loop string  `bson:"use_time_loop,omitempty" json:"use_time_loop"`
	Create_time   string  `bson:"create_time,omitempty" json:"create_time"`
	Sid           int64   `bson:"sid,omitempty" json:"sid"`
}

type Strategy_details struct {
	Name     string `json:"name" structs:"name"`
	Sid      int64  `json:"sid" structs:"sid"`
	Category string `json:"category" structs:"category"`
	Cid      string `json:"cid" structs:"cid"`
	Kline    string `json:"kline" structs:"kline"`
	Code     string `json:"code" structs:"code"`
	Ranger   string `json:"ranger" structs:"ranger"`
	Title    string `json:"title" structs:"title"`
	Describe string `json:"describe" structs:"describe"`
}
