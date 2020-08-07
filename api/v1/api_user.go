package v1

import (
	"bytes"
	"database/sql"
	"github.com/asmcos/requests"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
	"github.com/zheng-ji/goSnowFlake"
	"go.uber.org/zap"
	"goApi/config"
	"goApi/global"
	"goApi/global/response"
	"goApi/middleware"
	"goApi/model"
	"goApi/model/request"
	modelRes "goApi/model/response"
	"goApi/service"
	"goApi/utils"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func parseIP(ip string) (isp string, country string, ok bool) {
	reqUrl := config.JSonAPI + ip
	req := requests.Requests()
	req.SetTimeout(time.Second * 5)
	res, err := req.Post(reqUrl)
	if err != nil {
		return
	}
	jsonStr := res.Text()
	if !strings.Contains(jsonStr, "success") {
		return
	}
	ipResult := gojsonq.New().FromString(jsonStr)
	isp = ipResult.Copy().Find("isp").(string)
	country = ipResult.Copy().Find("country").(string)
	countryCode := ipResult.Copy().Find("countryCode").(string)
	countryCode = strings.ToLower(countryCode)
	country = countryCode + "_" + country
	ok = true
	return
}

// @Tags User
// @Summary 用户登录
// @accept application/json
// @Produce application/json
// @Param data body request.LoginReq true "用户登录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/user/login [post]
func UserLogin(c *gin.Context) {
	var ReqLogin request.LoginReq

	// 因为要改变L的值 所以必须传入地址
	_ = c.ShouldBindJSON(&ReqLogin)
	UserVerify := utils.Rules{
		"LoginType": {utils.NotEmpty()},
		"Token":     {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(ReqLogin, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}

	loginToken := ReqLogin.Token
	LoginType := ReqLogin.LoginType
	Install := ReqLogin.Install
	Avatar := ReqLogin.Avatar
	Email := ReqLogin.Email
	Name := ReqLogin.Name

	log := global.LoggerWithContext(c.Copy())

	var (
		subId string
	)

	// 选择 登录的方式
	switch LoginType {
	case 1:
		subId = googleTokenCheck(loginToken, log)
	case 2:
		subId = faceBookTokenCheck(loginToken)
	default:
		subId = googleTokenCheck(loginToken, log)
	}

	if subId == "" {
		response.FailWithMessage("login err, uid is null", c)
		return
	}
	// 采用 login type + subId 的 md5 值作为userid
	UserId := utils.MD5V([]byte(subId + utils.IntToString(int(LoginType))))

	Db := service.TUserInfoModel{
		DB: global.MysqlDb,
	}

	// 根据 userid md5 和login type 查询用户是否存在
	result, err := Db.GetUserInfoByUserId(UserId)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Info("查询用户出现错误,请查看", zap.Any("err", err))
			response.Fail(c)
			return
		} else {
			// err 等于 sql 内查询不到数据 创建一个空的 userInfo
			result = &model.TUserInfo{}
		}
	}
	var isRegister bool
	// 通过自定义time 进行时间的格式化
	now := model.Time(time.Now())
	// 用户不存在 执行注册逻辑
	if result.Rid == 0 {

		// 数据不存在走注册的逻辑
		result.LoginType = LoginType
		result.Sub = subId
		result.UserId = UserId
		result.NickName = Name
		result.Avatar = Avatar
		result.Status = 1
		result.Gender = 1
		result.Install = Install
		result.Email = Email
		result.Ctime = now.Time()
		result.Mtime = now.Time()
		// 开启事务
		tx, err := global.MysqlDb.Begin()
		if err != nil {
			response.FailWithMessage("开启事务失败", c)
			return
		}
		Db.Tx = tx

		ok, insertId := creteUser(Db, result, UserId)
		if !ok {
			response.FailWithMessage("注册失败", c)
			return
		}
		Db.Tx.Commit()

		result.Rid = int32(insertId)
		go DoRegisterTask(result, log)
		isRegister = true
		log.Info("用户创建成功", zap.Any("insertId", insertId))
		// 用户 存在 执行登录逻辑
	} else {
		isRegister = false
		// 数据存在走登录的逻辑
		log.Info("用户登录成功，用户信息", zap.Any("userInfo", result))
	}
	j := &middleware.JWT{
		SigningKey: []byte(global.SERVER_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UserId:  UserId,
		UserRid: int64(result.Rid),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "goApi",                    // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("jwt 签发失败", c)
		return
	}
	structData := modelRes.ToUserDto(*result)
	mapData := structs.New(structData).Map()
	mapData["token"] = token
	mapData["isRegister"] = isRegister
	global.REDIS_CONN.Set("userToken:"+UserId, token, time.Second*60*60*24*7)
	response.OkWithData(gin.H{"isRegister": isRegister, "token": token, "userid": UserId, "nick_name": Name}, c)
}

func DoRegisterTask(userResult *model.TUserInfo, log zap.Logger) {
	install := userResult.Install
	userId := userResult.UserId
	global.LOGGER.Info("DoRegisterTask", zap.Any("userId", userId), zap.Any("install", install))
	//b := strings.Contains(install, "invite")
	//if b {
	//	splitResult := strings.Split(install, "=")
	//	inviteUserId := splitResult[len(splitResult)-1]
	//	Db := service.TTaskConfModel{
	//		TUserInfoModel: service.TUserInfoModel{
	//			DB: global.MysqlDb,
	//		},
	//	}
	//
	//	// 邀请任务的任务id
	//	tid := "2"
	//	taskResult, err := Db.FindTaskInfoById(tid)
	//	if err != nil {
	//		log.Error("获取任务失败", zap.Any("err", err))
	//		return
	//	}
	//	Coins := taskResult.Coins
	//	result, _ := Db.GetUserMoney(inviteUserId)
	//	desc := fmt.Sprintf("%s, get money: %d coins", install, Coins)
	//
	//	allCus, _ := Db.QueryAllCustomer()
	//	var ccId string
	//	if len(allCus) >0{
	//		rand.Seed(time.Now().Unix())
	//		ccId = allCus[rand.Intn(len(allCus))]
	//	}
	//	_, err = Db.UpdateUserPromoteAndCcid(inviteUserId, userId, ccId)
	//
	//	if err != nil{
	//		log.Debug("更新推荐人，客服信息错误", zap.Any("err", err))
	//	}
	//
	//	_, err = ChangeMoneyTools(Coins, inviteUserId, 1, taskResult.TaskType, desc, result)
	//	updateResult := "update ok"
	//	if err != nil {
	//		updateResult = "update err"
	//		log.Info("更新邀请人金币失败", zap.Any("err", err))
	//	}
	//	log.Info("获取到邀请人的钱包数据-->更新结果", zap.String("inviteUserId", inviteUserId), zap.Any("updateResult", updateResult))
	//
	//}

}

func creteUser(Db service.TUserInfoModel, result *model.TUserInfo, userId string) (ok bool, insertId int64) {
	lastId, err := Db.CreateUserTx(result)
	ok = true
	if err != nil {
		ok = false
		global.LOGGER.Info("新增用户信息失败", zap.Error(err))
	}
	_, err = Db.CreateUserMoneyTx(userId)
	if err != nil {
		ok = false
		global.LOGGER.Info("新增用户钱包失败", zap.Error(err))
	}
	if !ok {
		Db.Tx.Rollback()
	}
	return ok, lastId
}

// @Tags User
// @Summary  update user info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateInfoReq true "更新 user info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/user/update [post]
func UpdateUserInfo(c *gin.Context) {
	var L request.UpdateInfoReq
	_ = c.ShouldBindJSON(&L)
	nickName := L.NickName
	Gender := L.Gender
	Avatar := L.Avatar
	Weight := L.Weight
	Height := L.Height
	Email := L.Email
	Birthday := L.Birthday

	claims, _ := c.Get("claims")
	claimsInfo := claims.(*request.CustomClaims)
	rid := claimsInfo.UserRid

	UserModel := service.TUserInfoModel{DB: global.MysqlDb}
	user, err := UserModel.GetUserInfoByRid(rid)
	if err != nil {
		if err != sql.ErrNoRows {
			response.Fail(c)
			return
		} else {
			response.FailWithDetailed(13, gin.H{}, "user not exists", c)
			return
		}
	}
	if nickName != "" {
		user.NickName = nickName
	}
	if Gender != 0 {
		user.Gender = Gender
	}
	if Avatar != "" {
		user.Avatar = utils.GetCdnUrl(Avatar)
	}
	if Weight != 0 {
		user.Weight = Weight
	}
	if Height != 0 {
		user.Height = Height
	}
	if Birthday != "" {
		re := regexp.MustCompile(config.BirthdayRegexp)
		birthdayBool := re.MatchString(Birthday)
		if !birthdayBool {
			response.FailWithMessage("Birthday not match", c)
			return
		}
		user.Birthday = Birthday
	}
	if Email != "" {
		re := regexp.MustCompile(config.EmailRegexp)
		emailBool := re.MatchString(Email)
		if !emailBool {
			response.FailWithMessage("Email not match", c)
			return
		}
		user.Email = Email
	}

	b, err := UserModel.Update(user)
	if err != nil {
		response.FailWithMessage("update err", c)
		return
	}
	log := global.LoggerWithContext(c)
	log.Info("update result", zap.Any("bool", b))
	response.OkWithData(modelRes.ToUserDto(*user), c)

}

// @Tags User
// @Summary  get user info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param userid path string false "查询用户的userid"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/user/info [get]
func GetUserInfo(c *gin.Context) {
	log := global.LoggerWithContext(c.Copy())
	claims, _ := c.Get("claims")
	claimsInfo := claims.(*request.CustomClaims)
	UserModel := service.TUserInfoModel{DB: global.MysqlDb}

	var queryId string

	// 如果从url 获取到userid 否则采用自己的userId
	aUserId := c.Query("userid")
	if aUserId == "" {
		queryId = claimsInfo.UserId
	} else {
		queryId = aUserId
	}

	// 根据 queryId 获取 用户
	user, err := UserModel.GetUserInfoByUserId(queryId)
	if err != nil {
		if err != sql.ErrNoRows {
			response.Fail(c)
			return
		} else {
			response.FailWithDetailed(13, gin.H{}, "user not exists", c)
			return
		}
	}

	result, err := UserModel.GetUserMoney(queryId)
	if err != nil {
		log.Error("获取用户的金币失败", zap.Any("err", err))
		return
	}

	mapData := modelRes.ToUserDtoMap(*user)

	//Db := service.UsergoApiRecordModel{
	//	TUserInfoModel: service.TUserInfoModel{
	//		DB: global.MysqlDb,
	//	},
	//}
	//
	//lastId, _ := Db.GetLastClientId(user.UserId)
	//
	//list, _ := Db.QueryAllCustomer()

	mapData["money"] = result.Money
	//mapData["latest_client_id"] = lastId
	//mapData["list"] = list
	response.OkWithData(mapData, c)

}

func googleTokenCheck(loginToken string, log zap.Logger) (uid string) {
	reqUrl := config.GoogleLoginURL + loginToken
	req := requests.Requests()
	req.SetTimeout(time.Second * 5)
	res, err := req.Post(reqUrl)
	if err != nil {
		return
	}
	googleJson := res.Text()

	jsonResult := gojsonq.New().FromString(googleJson)
	log.Sugar().Debug("google result:", jsonResult.Get())
	sub := jsonResult.Copy().Find("sub")
	if sub == nil {
		uid = ""
	} else {
		uid = sub.(string)
	}
	return
}

func faceBookTokenCheck(loginToken string) (uid string) {
	loginTokenUrl, _ := url.Parse(config.FBAppID + "|" + config.FBAppKEY)
	fbToken := loginTokenUrl.String()
	reqUrl := config.FaceBookLoginURL + fbToken + "&input_token=" + loginToken

	req := requests.Requests()
	req.SetTimeout(time.Second * 5)
	res, err := req.Get(reqUrl)
	global.LOGGER.Info("fb login", zap.Any("reqUrl", reqUrl), zap.Any("res", res))
	if err != nil {
		return
	}
	googleJson := res.Text()
	jsonResult := gojsonq.New().FromString(googleJson)
	sub := jsonResult.Copy().Find("data.user_id")
	if sub == nil {
		uid = ""
	} else {
		global.LOGGER.Info("获取到到fb token", zap.Any("sub", sub))
		uid = sub.(string)
	}
	return
}

// @Tags User
// @Summary  update user money
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateMoneyReq true "更新 user info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/user/moneyChange [post]
func UpdateUserMoney(c *gin.Context) {
	log := global.LoggerWithContext(c)

	var L request.UpdateMoneyReq
	_ = c.ShouldBindJSON(&L)
	Rid := L.Rid
	Money := L.Money
	Action := L.Action
	ActDesc := L.ActDesc
	ActType := L.ActType
	ReqVerify := utils.Rules{
		"Money":  {utils.NotEmpty(), utils.Gt("0")},
		"Rid":    {utils.NotEmpty()},
		"Action": {utils.Gt("0")},
	}
	reqVerifyErr := utils.Verify(L, ReqVerify)
	if reqVerifyErr != nil {
		response.FailWithMessage(reqVerifyErr.Error(), c)
		return
	}

	Db := service.TUserInfoModel{DB: global.MysqlDb}


	claims, _ := c.Get("claims")
	claimsInfo := claims.(*request.CustomClaims)
	opRid := claimsInfo.UserRid

	// 通过 jwt 携带的信息查询登录者的信息
	opUser, err := Db.GetUserInfoByRid(opRid)
	if err != nil {
		if err != sql.ErrNoRows {
			response.Fail(c)
			return
		} else {
			response.FailWithDetailed(13, gin.H{}, "user not exists", c)
			return
		}
	}
	isCustomerCare := opUser.IsCustomerCare
	if isCustomerCare != 1{
		response.FailWithMessage("role err", c)
		return
	}

	// 通过参数查找用户信息
	user, err := Db.GetUserInfoByRid(Rid)
	if err != nil {
		if err != sql.ErrNoRows {
			response.Fail(c)
			return
		} else {
			response.FailWithDetailed(13, gin.H{}, "user not exists", c)
			return
		}
	}

	userId := user.UserId

	// 获取当前用户有多少金币
	result, err := Db.GetUserMoney(userId)
	if err != nil {
		return
	}

	// 减操作 最好先做一下判断再执行 sql，减少数据库的io操作
	oldMoney := result.Money
	if Action == 2 && oldMoney <= Money {
		response.FailWithMessage("update err, please check user money", c)
		return
	}

	AfterChange, err := ChangeMoneyTools(Money, userId, Action, ActType, ActDesc, result)
	if err != nil {
		log.Error("更新用户金币失败", zap.String("err", err.Error()))
		response.FailWithMessage("update error", c)
		return
	}

	data := modelRes.ToUserDtoMap(*user)
	data["money"] = AfterChange
	response.OkWithData(data, c)

}

func ChangeMoneyTools(money int64, userId string, doAct int, ActType int32, ActDesc string, result *model.TUserMoney) (AfterChange int64, err error) {
	log := global.LOGGER
	Db := service.TUserInfoModel{DB: global.MysqlDb}
	// 更新 金币
	b, err := Db.UpdateUserMoney(money, userId, doAct)
	if err != nil || b == false {
		return 0, err
	}

	// decimal 的加减操作
	if doAct == 1 {
		AfterChange = result.Money + money
	} else {
		AfterChange = result.Money - money
	}

	// 构建 UserMoneyRecord 数据
	logDetail := model.UserMoneyRecord{
		UserId:       userId,
		Change:       money,
		BeforeChange: result.Money,
		AfterChange:  AfterChange,
		Action:       int32(doAct),
		ActDesc:      ActDesc,
		ActType:      ActType,
		Remark:       []byte("{}"),
	}
	// 创建用户的金币流水记录
	_, errRecord := Db.CreateUserMoneyRecord(&logDetail)
	if errRecord != nil {
		log.Error("创建记录失败", zap.Any("err", errRecord.Error()))
	}
	return AfterChange, nil
}

// @Tags User
// @Summary  测试方法
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateMoneyReq true "测试方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/user/test [post]
func TestMethod(c *gin.Context) {

	log := global.LoggerWithContext(c)

	// 直接获取 body 内容
	dataJson, err := c.GetRawData()
	if err != nil {

	}
	//ipResult := gojsonq.New().FromString(string(dataJson))
	//testList := ipResult.Copy().Find("list")
	//log.Info("测试请求信息", zap.Any("json", string(dataJson)), zap.Any("testList", testList))
	//很关键
	//把读过的字节流重新放到body
	//不然下面到处理器读取不到数据
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(dataJson))

	id, _ := global.UniqId.NextId()
	t, ts, workerId, seq := goSnowFlake.ParseId(id)
	log.Info("parseId", zap.Any("t", t), zap.Any("ts", ts), zap.Any("workerId", workerId), zap.Any("seq", seq))
	response.OkWithData(gin.H{"id": id}, c)
}

// @Tags User
// @Summary 苹果登陆
// @Produce  application/json
// @Param data body request.AppleLoginCodeAndId true "苹果登陆"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /appLogin/AppleLoginCode [post]
func AppleLogin(c *gin.Context) {
	//一个请求信息的结构体，接收authorizationCode和userIdentifier
	var appleReq request.AppleLoginCodeAndId
	_ = c.ShouldBindJSON(&appleReq)
	//生成client_secret，参数：苹果账户的KeyId,TeamId, ClientID, KeySecret
	clientSecret := utils.GetAppleSecret(config.AppleLoginKeyId, config.AppleLoginTeamId, config.AppleLoginClientID, config.AppleLoginKeySecret)

	//获取用户信息，参数：ClientID, 上面的clientSecret, authorizationCode
	data, err := utils.GetAppleLoginData(config.AppleLoginClientID, clientSecret, appleReq.AuthorizationCode)

	//检查用户信息，参数：上面的data，客户端传过来的userIdentifier，上面的err
	ok := utils.CheckAppleID(data, appleReq.UserIdentifier, err)
	if !ok {
		response.FailWithMessage("APPLE登陆失败[error 1]，请重试或使用其他方式登陆", c)
		return
	}

	//你的登陆业务

}
