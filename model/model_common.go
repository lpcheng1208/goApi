package model

type FaceBookToken struct {
	Options struct {
		Googleplay struct {
			AppName string `json:"app_name" bson:"app_name"`
			AppKey  string `json:"app_key" bson:"app_key"`
			AppID   string `json:"app_id" bson:"app_id"`
		} `bson:"googleplay"`
		Purechatlite struct {
			AppName string `bson:"app_name"`
			AppKey  string `bson:"app_key"`
			AppID   string `bson:"app_id"`
		} `bson:"purechatlite"`
		Webcam struct {
			AppName string `bson:"app_name"`
			AppKey  string `bson:"app_key"`
			AppID   string `bson:"app_id"`
		} `bson:"webcam"`
		Lila struct {
			AppName string `bson:"app_name"`
			AppKey  string `bson:"app_key"`
			AppID   string `bson:"app_id"`
		} `bson:"Lila"`
		Talkcool struct {
			AppName string `bson:"app_name"`
			AppKey  string `bson:"app_key"`
			AppID   string `bson:"app_id"`
		} `bson:"talkcool"`
		Chatoo struct {
			AppName string `bson:"app_name"`
			AppKey  string `bson:"app_key"`
			AppID   string `bson:"app_id"`
		} `bson:"chatoo"`
	} `bson:"options"`
	// 不是bson.ObjectId 不需要 omitempty
	ID string `bson:"_id"`
}

type PaginationResult struct {
	List     interface{} `json:"list"`
	NextPage bool        `json:"next_page"`
}
