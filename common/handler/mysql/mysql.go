package mysql

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sharp/common/consts"
	"sharp/common/handler/conf"
	"sharp/common/handler/log"
	"time"
)

var (
	MysqlClient *gorm.DB
)

func Init() {
	host := conf.Viper.GetString("mysql.host")
	dsn := conf.Viper.GetString("mysql.dsn")
	dsn = fmt.Sprintf(dsn, host)

	logMap := map[string]interface{}{
		"host": host,
	}

	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logMap[consts.LogErrMsg] = err.Error()
		log.Logger.Fatalf(consts.DLTagInitMysqlFailed, log.BuildLogByMap(context.Background(), logMap))
		os.Exit(1)
	}

	sqlDB, err := dbClient.DB()
	if err != nil {
		logMap[consts.LogErrMsg] = err.Error()
		log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(context.Background(), logMap))
	}

	sqlDB.SetConnMaxLifetime(time.Duration(conf.Viper.GetInt("mysql.conn_max_lifetime")) * time.Second)
	sqlDB.SetMaxIdleConns(conf.Viper.GetInt("mysql.max_idel_conns"))
	sqlDB.SetMaxOpenConns(conf.Viper.GetInt("mysql.max_open_conns"))

	MysqlClient = dbClient
	log.Logger.Infof(consts.DLTagInitMysqlSuccess, log.BuildLogByMap(context.Background(), logMap))
}
