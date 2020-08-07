package config

const (
	GoogleLoginURL   = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="
	FaceBookLoginURL = "https://graph.facebook.com/debug_token?access_token="
	JSonAPI          = "http://ip-api.com/json/"
	FBAppID          = "8a1345b284d2351362cb953a20d0e6ef"
	FBAppKEY         = "1710734882415042"

	PageSize = 20

	// apple 相关常量定义
	AppleLoginKeyId     = ""
	AppleLoginTeamId    = ""
	AppleLoginClientID  = ""
	AppleLoginKeySecret = ""

	SendMsgToClientDev = "http://127.0.0.1:8600/goim/send_client"
	SendMsgToClientProd = "https://api.wetene.com/goim/send_client"
	TimeFormat      = "2006-01-02T15:04:05.000Z0700"

	CdnPrefix = "https://d3hrn6f85i87xg.cloudfront.net/"

	EmailRegexp    = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	BirthdayRegexp = "^((((19|20)\\d{2})-(0?[13-9]|1[012])-(0?[1-9]|[12]\\d|30))|(((19|20)\\d{2})-(0?[13578]|1[02])-31)|(((19|20)\\d{2})-0?2-(0?[1-9]|1\\d|2[0-8]))|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))-0?2-29))$"
)
