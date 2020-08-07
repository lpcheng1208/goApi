package request

import "github.com/shopspring/decimal"

type UpdateInfoReq struct {
	NickName string          `json:"nick_name" form:"nick_name"`
	Avatar   string          `json:"avatar" form:"avatar"`
	Gender   int32           `json:"gender" form:"gender"`
	Birthday string          `json:"birthday" form:"birthday"`
	Height   int32           `json:"height" form:"height"`
	Weight   int32           `json:"weight" form:"weight"`
	Money    decimal.Decimal `json:"money" form:"weight"`
	Userid   string          `json:"userid" form:"weight"`
	Email    string          `json:"email" form:"email"`
	Rid      int32           `json:"rid" form:"rid"`
}

type LoginReq struct {
	Ver       int    `json:"ver"`
	LoginType int32  `json:"login_type"`
	Channel   string `json:"channel"`
	Deviceid  string `json:"deviceId"`
	Token     string `json:"token"`
	Install   string `json:"install"`
	Name      string `json:"nick_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type UpdateMoneyReq struct {
	Action  int             `json:"action" form:"action"`
	Money   int64 `json:"money" form:"money"`
	Userid  string          `json:"userid" form:"userid"`
	Rid     int64           `json:"rid" form:"rid"`
	ActDesc string          `json:"act_desc" form:"act_desc"`
	ActType int32           `json:"act_type" form:"act_type"`
}
