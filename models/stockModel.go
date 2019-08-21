package models


import "time"

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id" structs:"id"`
	CreatedAt time.Time  `json:"createdAt" structs:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" structs:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt" structs:"deletedAt"`
}


type Article struct {
	Model

	AuthorID     uint64    `json:"authorID" structs:"authorID"`
	Title        string    `gorm:"size:128" json:"title" structs:"title"`
	Abstract     string    `gorm:"type:mediumtext" json:"abstract" structs:"abstract"`
	Tags         string    `gorm:"type:text" json:"tags" structs:"tags"`
	Content      string    `gorm:"type:mediumtext" json:"content" structs:"content"`
	Path         string    `sql:"index" gorm:"size:255" json:"path" structs:"path"`
	Status       int       `sql:"index" json:"status" structs:"status"`
	Topped       bool      `json:"topped" structs:"topped"`
	Commentable  bool      `json:"commentable" structs:"commentable"`
	ViewCount    int       `json:"viewCount" structs:"viewCount"`
	CommentCount int       `json:"commentCount" structs:"commentCount"`
	IP           string    `gorm:"size:128" json:"ip" structs:"ip"`
	UserAgent    string    `gorm:"size:255" json:"userAgent" structs:"userAgent"`
	PushedAt     time.Time `json:"pushedAt" structs:"pushedAt"`

	BlogID uint64 `sql:"index" json:"blogID" structs:"blogID"`
}

type ApiData struct {

	Code     uint64    `json:"id" structs:"id"`
	Msg      string    `json:"msg" structs:"msg"`
	Count    uint64    `json:"id" structs:"id"`
	Data     uint64    `json:"id" structs:"id"`
}


// 做多top9bar10
type LongShort struct {
	//Model

	Content      string    `json:"content" structs:"content"`
	Id           string    `json:"id" structs:"id"`
	Uptime       time.Time `json:"uptime" structs:"uptime"`
	Create_time  string `json:"create_time" structs:"create_time"`

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

	Code         string    `json:"code" structs:"code"`
	Change_rate  float64   `json:"change_rate" structs:"change_rate"`
	Time_key     string    `json:"time_key" structs:"time_key"`

}


type Top10 struct {
	//Model
	Code         int16                 `json:"code" structs:"code"`
	Count        int32                 `json:"count" structs:"count"`
	Msg   		 string                `json:"msg" structs:"msg"`
	Data         []map[string]string   `json:"Data" structs:"Data"`
}


