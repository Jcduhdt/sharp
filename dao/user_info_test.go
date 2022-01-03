package dao

import (
	"context"
	"sharp/common/handler/conf"
	"sharp/common/handler/env"
	"sharp/common/handler/log"
	"sharp/common/handler/mysql"
	"sharp/common/handler/redis"
	"sharp/common/util"
	"testing"
)

var (
	ub  = UserBase{}
	ctx = context.Background()
)

func init(){
	env.Init()
	conf.InitConf(conf.FindConfigDir())
	log.Init()
	mysql.Init()
	redis.Init()
}

func TestUserBase_GetUserInfo(t *testing.T) {
	queryParams := map[string]interface{}{
		"nick_name": "Jcduhdt",
		"passwd":   util.MD5("123456789"),
	}
	userInfo, err := GetUserInfo(ctx, queryParams)
	t.Logf("res = %+v, err = %+v", userInfo, err)
}

func TestUserBase_InsertUserInfo(t *testing.T) {
	userBaseInsert := UserBaseInsert{
		UserName:  "ives zhang",
		NickName:  "Jcduhdt",
		Passwd:    util.MD5("123456789"),
		Gender:    1,
		Signature: "你是我唯一想要的了解",
	}
	err := InsertUserInfo(ctx,userBaseInsert)
	t.Logf("err = %+v", err)
}