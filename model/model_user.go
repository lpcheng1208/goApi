// 自动生成模板Message
package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TUserInfo struct {
	Rid            int32
	NickName       string `json:"nick_name" structs:"nick_name"`
	Avatar         string
	Gender         int32
	Birthday       string
	Height         int32
	Weight         int32
	Install        string
	LoginType      int32
	Sub            string
	UserId         string
	Email          string
	Ccid           string
	PromoteUser    string
	IsCustomerCare int32
	Status         int32
	Ctime          time.Time // 创建时间
	Mtime          time.Time // 更新时间

}

type TUserInfoNull struct {
	Rid            sql.NullInt32
	NickName       sql.NullString `json:"nick_name" structs:"nick_name"`
	Avatar         sql.NullString
	Gender         sql.NullInt32
	Birthday       sql.NullString
	Height         sql.NullInt32
	Weight         sql.NullInt32
	Install        sql.NullString
	LoginType      sql.NullInt32
	Sub            sql.NullString
	UserId         sql.NullString
	Email          sql.NullString
	Ccid           sql.NullString
	PromoteUser    sql.NullString
	IsCustomerCare sql.NullInt32
	Status         sql.NullInt32
	Ctime          time.Time // 创建时间
	Mtime          time.Time // 更新时间

}

type TUserMoney struct {
	Id           int32
	UserId       string
	Money        int64
	SumInto      int64
	SumOut       int64
	TotalRevenue int64
	Ctime        Time // 创建时间
	Mtime        Time // 更新时间

}

type TUserMoneyNull struct {
	Id           sql.NullInt32
	UserId       sql.NullString
	Money        sql.NullInt64
	SumInto      sql.NullInt64
	SumOut       sql.NullInt64
	TotalRevenue sql.NullInt64
	Ctime        Time // 创建时间
	Mtime        Time // 更新时间

}

type UserMoneyRecord struct {
	UserId       string
	Change       int64
	ActType      int32
	BeforeChange int64
	AfterChange  int64
	Action       int32
	ActDesc      string          // 创建时间
	Remark       json.RawMessage // 更新时间
	Ctime        Time            // 创建时间
	Mtime        Time            // 更新时间
}
