package user_service

import (
	"context"
	"errors"
	"sharp/common/consts"
	"sharp/common/dto"
	"sharp/common/handler/log"
	"sharp/common/util"
	"sharp/dao"
)

func Login(ctx context.Context, req dto.LoginAndRegisterReq) (bool, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "login",
		consts.LogParams: req,
	}

	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(ctx, logMap))
		}
	}()

	queryParams := map[string]interface{}{
		"nick_name": req.NickName,
		"passwd":    util.MD5(req.PassWord),
	}

	_, err = dao.GetUserInfo(ctx, queryParams)

	if errors.Is(err, consts.ErrLoginError) {
		err = nil
		return false, nil
	}

	return true, err
}

func Register(ctx context.Context, req dto.LoginAndRegisterReq) (bool, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "register",
		consts.LogParams: req,
	}

	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(ctx, logMap))
		}
	}()

	userBaseInsert := dao.UserBaseInsert{
		NickName: req.NickName,
		Passwd:   util.MD5(req.PassWord),
	}

	err = dao.InsertUserInfo(ctx, userBaseInsert)
	if err != nil {
		return false, nil
	}

	return true, err
}
