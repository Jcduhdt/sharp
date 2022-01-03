package dao

import (
	"context"
	"sharp/common/consts"
	"sharp/common/handler/log"
	"sharp/common/handler/mysql"
	"time"
)

type UserBase struct {
	Id         int64     `json:"id"`
	UserName   string    `json:"user_name"`
	NickName   string    `json:"nick_name"`
	Passwd     string    `json:"passwd"`
	Gender     int       `json:"gender"`
	Birthday   time.Time `json:"birthday"`
	Mobile     int       `json:"mobile"`
	Face       string    `json:"face"`
	Email      string    `json:"email"`
	Signature  string    `json:"signature"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}

type UserBaseInsert struct {
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Passwd    string `json:"passwd"`
	Gender    int    `json:"gender"`
	Mobile    int    `json:"mobile"`
	Face      string `json:"face"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
}

func GetUserInfo(ctx context.Context, queryParams map[string]interface{}) (UserBase, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "get_user_info",
		consts.LogParams: queryParams,
	}

	var err error
	var res UserBase
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagGetMysqlFailed, log.BuildLogByMap(ctx, logMap))
			return
		}
		log.Logger.Infof(consts.DLTagGetMysqlSuccess, log.BuildLogByMap(ctx, logMap))
	}()

	tableName := "user_base"

	start := time.Now()
	db := mysql.MysqlClient.Table(tableName).Where(queryParams)
	err = db.Find(&res).Error
	logMap["cost"] = time.Since(start)
	logMap["res"] = res

	if err == nil && db.RowsAffected == 0 {
		return res, consts.ErrLoginError
	}

	return res, nil
}

func InsertUserInfo(ctx context.Context, ubi UserBaseInsert) (err error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "insert_user_info",
		consts.LogParams: ubi,
	}

	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagGetMysqlFailed, log.BuildLogByMap(ctx, logMap))
			return
		}
		log.Logger.Infof(consts.DLTagGetMysqlSuccess, log.BuildLogByMap(ctx, logMap))
	}()

	tableName := "user_base"

	start := time.Now()
	db := mysql.MysqlClient.Table(tableName).Create(&ubi)
	err = db.Error
	logMap["cost"] = time.Since(start)

	return
}
