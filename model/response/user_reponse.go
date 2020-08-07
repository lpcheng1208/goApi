package response

import (
	"github.com/fatih/structs"
	"goApi/model"
)

type UserInfoRes struct {
	Rid            int32  `structs:"rid" json:"rid" form:"rid"`
	UserId         string `structs:"userid" json:"userid" form:"userid"`
	NickName       string `structs:"nick_name" json:"nick_name" form:"nick_name"`
	Avatar         string `structs:"avatar" json:"avatar" form:"avatar"`
	Gender         int32  `structs:"gender" json:"gender" form:"gender"`
	Birthday       string `structs:"birthday" json:"birthday" form:"birthday"`
	Email          string `structs:"email" json:"email" form:"email"`
	Height         int32  `structs:"height" json:"height" form:"height"`
	Weight         int32  `structs:"weight" json:"weight" form:"weight"`
	PromoteUser    string `structs:"promote_user" json:"promote_user" form:"promote_user"`
	IsCustomerCare int32  `structs:"is_customer_care" json:"is_customer_care" form:"is_customer_care"`
}

func ToUserDto(user model.TUserInfo) UserInfoRes {
	return UserInfoRes{
		Rid:            user.Rid,
		UserId:         user.UserId,
		NickName:       user.NickName,
		Avatar:         user.Avatar,
		Gender:         user.Gender,
		Birthday:       user.Birthday,
		Height:         user.Height,
		Weight:         user.Weight,
		Email:          user.Email,
		PromoteUser:    user.PromoteUser,
		IsCustomerCare: user.IsCustomerCare,
	}
}

// 转换成 map 方便添加额外的属性
func ToUserDtoMap(user model.TUserInfo) (mapData map[string]interface{}) {
	structData := ToUserDto(user)
	mapData = structs.New(structData).Map()
	return
}
