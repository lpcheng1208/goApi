//
package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goApi/model"
	"goApi/model/model_name"
)

type TUserInfoModel struct {
	E
	DB *sql.DB
	Tx *sql.Tx
}

// 获取所有的表字段
func (m *TUserInfoModel) getColumns() string {
	return " `rid`,`nick_name`,`avatar`,`gender`,`birthday`,`height`,`weight`,`install`, `login_type`, `sub`, `userid`,`email`,`ccid`,`promote_user`,`is_customer_care`,`status`,`ctime`,`mtime` "
}

// 获取多行数据.
func (m *TUserInfoModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*model.TUserInfo, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer query.Close()
	for query.Next() {
		row := model.TUserInfoNull{}
		err = query.Scan(
			&row.Rid,       //
			&row.NickName,  //
			&row.Avatar,    //
			&row.Gender,    //
			&row.Birthday,  //
			&row.Height,    //
			&row.Weight,    //
			&row.Install,   //
			&row.LoginType, //
			&row.Sub,       //
			&row.UserId,    //
			&row.Email,     //
			&row.Ccid,      //
			&row.PromoteUser,
			&row.IsCustomerCare,
			&row.Status, //
			&row.Ctime,  // 创建时间
			&row.Mtime,  // 更新时间
		)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		rowsResult = append(rowsResult, &model.TUserInfo{
			Rid:            row.Rid.Int32,       //
			NickName:       row.NickName.String, //
			Avatar:         row.Avatar.String,   //
			Gender:         row.Gender.Int32,    //
			Birthday:       row.Birthday.String, //
			Height:         row.Height.Int32,    //
			Weight:         row.Weight.Int32,    //
			Install:        row.Install.String,  //
			LoginType:      row.Weight.Int32,    //
			Sub:            row.Sub.String,      //
			UserId:         row.UserId.String,   //
			Email:          row.Email.String,    //
			Ccid:           row.Ccid.String,     //
			PromoteUser:    row.PromoteUser.String,
			IsCustomerCare: row.IsCustomerCare.Int32,
			Status:         row.Status.Int32, //
			Ctime:          row.Ctime,        // 创建时间
			Mtime:          row.Mtime,        // 更新时间
		})
	}
	return
}

// 获取单行数据
func (m *TUserInfoModel) getRow(sqlText string, params ...interface{}) (rowResult *model.TUserInfo, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := model.TUserInfoNull{}
	err = query.Scan(
		&row.Rid,       //
		&row.NickName,  //
		&row.Avatar,    //
		&row.Gender,    //
		&row.Birthday,  //
		&row.Height,    //
		&row.Weight,    //
		&row.Install,   //
		&row.LoginType, //
		&row.Sub,       //
		&row.UserId,    //
		&row.Email,     //
		&row.Ccid,      //
		&row.PromoteUser,
		&row.IsCustomerCare,
		&row.Status, //
		&row.Ctime,  // 创建时间
		&row.Mtime,  // 更新时间
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	rowResult = &model.TUserInfo{
		Rid:            row.Rid.Int32,       //
		NickName:       row.NickName.String, //
		Avatar:         row.Avatar.String,   //
		Gender:         row.Gender.Int32,    //
		Birthday:       row.Birthday.String, //
		Height:         row.Height.Int32,    //
		Weight:         row.Weight.Int32,    //
		Install:        row.Install.String,  //
		LoginType:      row.Weight.Int32,    //
		Sub:            row.Sub.String,      //
		UserId:         row.UserId.String,   //
		Email:          row.Email.String,    //
		Ccid:           row.Ccid.String,     //
		PromoteUser:    row.PromoteUser.String,
		IsCustomerCare: row.IsCustomerCare.Int32,
		Status:         row.Status.Int32, //
		Ctime:          row.Ctime,        // 创建时间
		Mtime:          row.Mtime,        // 更新时间
	}
	return
}

// _更新数据
func (m *TUserInfoModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	b = affectCount > 0
	return
}

// _更新数据
func (m *TUserInfoModel) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	b = affectCount > 0
	return
}

// 新增信息
func (m *TUserInfoModel) Create(value *model.TUserInfo) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + model_name.TABLE_T_USER_INFO + " (`nick_name`,`avatar`,`gender`,`birthday`,`Height`,`weight`,`install`, `login_type`, `sub`, `userid`,`status`,`ctime`,`mtime`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.NickName,  //
		value.Avatar,    //
		value.Gender,    //
		value.Birthday,  //
		value.Height,    //
		value.Weight,    //
		value.Install,   //
		value.LoginType, //
		value.Sub,       //
		value.UserId,    //
		value.Status,    //
		value.Ctime,     // 创建时间
		value.Mtime,     // 更新时间
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	return
}

// 新增信息 tx
func (m *TUserInfoModel) CreateUserTx(value *model.TUserInfo) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + model_name.TABLE_T_USER_INFO + " (`nick_name`,`avatar`,`gender`,`birthday`,`height`,`weight`,`install`, `login_type`, `sub`, `userid`,`email`, `ccid`,`promote_user`,`status`,`ctime`,`mtime`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.NickName,    //
		value.Avatar,      //
		value.Gender,      //
		value.Birthday,    //
		value.Height,      //
		value.Weight,      //
		value.Install,     //
		value.LoginType,   //
		value.Sub,         //
		value.UserId,      //
		value.Email,       //
		value.Ccid,        //
		value.PromoteUser, //
		value.Status,      //
		value.Ctime,       // 创建时间
		value.Mtime,       // 更新时间
	)
	if err != nil {
		err = m.E.Stack(err)
		m.Tx.Rollback()
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		m.Tx.Rollback()
		return
	}
	return
}

// 新增用户钱包信息 tx
func (m *TUserInfoModel) CreateUserMoneyTx(userid string) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + "t_user_money" + " (`userid`, `money`, `sum_into`, `sum_out`, `total_revenue`, `ctime`, `mtime`) VALUES (?,0,0,0,0,NOW(),Now())"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		userid,
	)
	if err != nil {
		err = m.E.Stack(err)
		m.Tx.Rollback()
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		m.Tx.Rollback()
		return
	}
	return
}

// 更新数据
func (m *TUserInfoModel) Update(value *model.TUserInfo) (b bool, err error) {
	const sqlText = "UPDATE " + model_name.TABLE_T_USER_INFO + " SET `nick_name`=?,`avatar`=?,`gender`=?,`birthday`=?,`Height`=?,`weight`=? WHERE `rid` = ?"
	params := make([]interface{}, 0)
	params = append(params, value.NickName)
	params = append(params, value.Avatar)
	params = append(params, value.Gender)
	params = append(params, value.Birthday)
	params = append(params, value.Height)
	params = append(params, value.Weight)
	params = append(params, value.Rid)

	return m.Save(sqlText, params...)
}

// 更新数据 tx
func (m *TUserInfoModel) UpdateTx(value *model.TUserInfo) (b bool, err error) {
	const sqlText = "UPDATE " + model_name.TABLE_T_USER_INFO + " SET `nick_name`=?,`avatar`=?,`gender`=?,`birthday`=?,`Height`=?,`weight`=?,`status`=?,`ctime`=?,`mtime`=? WHERE `rid` = ?"
	params := make([]interface{}, 0)
	params = append(params, value.NickName)
	params = append(params, value.Avatar)
	params = append(params, value.Gender)
	params = append(params, value.Birthday)
	params = append(params, value.Height)
	params = append(params, value.Weight)
	params = append(params, value.Status)
	params = append(params, value.Ctime)
	params = append(params, value.Mtime)
	params = append(params, value.Rid)

	return m.SaveTx(sqlText, params...)
}

// 查询多行数据
func (m *TUserInfoModel) Find() (resList []*model.TUserInfo, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + model_name.TABLE_T_USER_INFO
	resList, err = m.getRows(sqlText)
	return
}

// 获取最后一行数据
func (m *TUserInfoModel) Last() (result *model.TUserInfo, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + model_name.TABLE_T_USER_INFO + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	return
}

// 单列数据
func (m *TUserInfoModel) Pluck(id int64) (result map[int64]interface{}, err error) {
	const sqlText = "SELECT `rid`, `nick_name` FROM " + model_name.TABLE_T_USER_INFO + " where `rid` = ?"
	rows, err := m.DB.Query(sqlText, id)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	result = make(map[int64]interface{})
	var (
		_id  int64
		_val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 单列数据 by 支持切片传入
// Get column data
func (m *TUserInfoModel) Plucks(ids []int64) (result map[int64]interface{}, err error) {
	result = make(map[int64]interface{})
	if len(ids) == 0 {
		return
	}
	sqlText := "SELECT `rid`, `nick_name` FROM " + model_name.TABLE_T_USER_INFO + " where " +
		"`rid` in (" + RepeatQuestionMark(len(ids)) + ")"
	params := make([]interface{}, len(ids))
	for idx, id := range ids {
		params[idx] = id
	}
	rows, err := m.DB.Query(sqlText, params...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	var (
		_id  int64
		_val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 获取单个数据
// Get one data
func (m *TUserInfoModel) One(id int64) (result int64, err error) {
	sqlText := "SELECT `rid` FROM " + model_name.TABLE_T_USER_INFO + " where `rid`=?"
	err = m.DB.QueryRow(sqlText, id).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return
}

// 获取行数
// Get line count
func (m *TUserInfoModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + model_name.TABLE_T_USER_INFO
	err = m.DB.QueryRow(sqlText).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return
}

// 判断数据是否存在
// Check the data is have?
func (m *TUserInfoModel) Has(id int64) (b bool, err error) {
	sqlText := "SELECT `rid` FROM " + model_name.TABLE_T_USER_INFO + " where `rid` = ?"
	var count int64
	err = m.DB.QueryRow(sqlText, id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return count > 0, nil
}

// 根据UserId 获取User
func (m *TUserInfoModel) GetUserInfoByLoginId(userid string, loginType int32) (result *model.TUserInfo, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + model_name.TABLE_T_USER + " where login_type=? and userid=? limit 1"
	result, err = m.getRow(sqlText, loginType, userid)
	return
}

// 根据UserId 获取User
func (m *TUserInfoModel) GetUserMoney(userid string) (result *model.TUserMoney, err error) {
	sqlText := "SELECT `userid`, `money`, `sum_into`, `sum_out`, `total_revenue`, `ctime`, `mtime` " + "FROM t_user_money " + " where userid=? limit 1"
	result, err = m.getMoneyRow(sqlText, userid)
	return
}

// 获取单行数据
func (m *TUserInfoModel) getMoneyRow(sqlText string, params ...interface{}) (rowResult *model.TUserMoney, err error) {
	query := m.DB.QueryRow(sqlText, params...)
	row := model.TUserMoneyNull{}
	err = query.Scan(
		&row.UserId,
		&row.Money,
		&row.SumInto,
		&row.SumOut,
		&row.TotalRevenue,
		&row.Ctime, // 创建时间
		&row.Mtime, // 更新时间
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	rowResult = &model.TUserMoney{
		Id:           0,
		UserId:       row.UserId.String,
		Money:        row.Money.Int64,
		SumInto:      row.SumInto.Int64,
		SumOut:       row.SumOut.Int64,
		TotalRevenue: row.TotalRevenue.Int64,
		Ctime:        row.Ctime,
		Mtime:        row.Mtime,
	}
	return
}

// 更新数据
func (m *TUserInfoModel) UpdateUserMoney(money int64, userId string, doAct int) (b bool, err error) {
	var sqlText = "UPDATE t_user_money SET "
	params := make([]interface{}, 0)
	if doAct == 1 {
		sqlText += " `money`=money + ? ,`sum_into`=sum_into + ?, `total_revenue`= total_revenue + ?"
		params = append(params, money)
		params = append(params, money)
		params = append(params, money)
	} else {
		sqlText += " `money`=money - ? ,`sum_out`=sum_out + ?"
		params = append(params, money)
		params = append(params, money)
	}

	sqlText += ",`mtime`=Now() WHERE `userid` = ?"
	params = append(params, userId)

	return m.Save(sqlText, params...)
}

// 新增用户钱包信息 tx
func (m *TUserInfoModel) CreateUserMoneyRecord(moneyRecord *model.UserMoneyRecord) (lastId int64, err error) {
	const sqlText = "INSERT INTO " + "user_money_record" + " (`userid`, `change`, `act_type`, `before_change`, `after_change`, `action`, `act_desc`, `remark`,`ctime`, `mtime`) VALUES (?,?,?,?,?,?,?,?,NOW(),Now())"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		moneyRecord.UserId,
		moneyRecord.Change,
		moneyRecord.ActType,
		moneyRecord.BeforeChange,
		moneyRecord.AfterChange,
		moneyRecord.Action,
		moneyRecord.ActDesc,
		moneyRecord.Remark,
	)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	return
}

// 通过主键id 获取 用户信息
func (m *TUserInfoModel) GetUserInfoByRid(id int64) (result *model.TUserInfo, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + model_name.TABLE_T_USER_INFO + " WHERE `rid` = ? LIMIT 1"
	result, err = m.getRow(sqlText, id)
	return
}

// 根据UserId 获取User
func (m *TUserInfoModel) GetUserInfoByUserId(userid string) (result *model.TUserInfo, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + model_name.TABLE_T_USER + " where userid=? limit 1"
	result, err = m.getRow(sqlText, userid)
	return
}

// 根据时间检查当天是否完成任务
func (m *TUserInfoModel) GetMoneyRecordByType(userid string, actType int32, startTime string) (b bool, err error) {
	sqlText := "SELECT id FROM user_money_record" + " where userid=? and act_type=? and ctime>=?"
	var count int64
	err = m.DB.QueryRow(sqlText, userid, actType, startTime).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		err = m.E.Stack(err)
		return
	}
	return count > 0, nil
}

// 更新数据
func (m *TUserInfoModel) UpdateUserPromoteAndCcid(InviteId string, userId string, Ccid string) (b bool, err error) {
	const sqlText = "UPDATE " + model_name.TABLE_T_USER_INFO + " SET `promote_user`=?, `ccid`=?, `mtime`=NOW() WHERE `userid` = ?"
	params := make([]interface{}, 0)
	params = append(params, InviteId)
	params = append(params, Ccid)
	params = append(params, userId)

	return m.Save(sqlText, params...)
}


// 查询多行数据
func (m *TUserInfoModel) QueryAllCustomer() (resList []string, err error) {
	sqlText := "SELECT userid FROM " + model_name.TABLE_T_USER_INFO + " Where is_customer_care=1"
	query, err := m.DB.Query(sqlText)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer query.Close()

	for query.Next(){
		var userId string
		query.Scan(&userId)
		resList = append(resList, userId)
	}
	return
}