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
		log.FatalMap(context.Background(), consts.DLTagInitMysqlFailed, logMap)
		os.Exit(1)
	}

	sqlDB, err := dbClient.DB()
	if err != nil {
		logMap[consts.LogErrMsg] = err.Error()
		log.FatalMap(context.Background(), consts.DLTagCommonErrorInfo, logMap)
		os.Exit(1)
	}

	sqlDB.SetConnMaxLifetime(time.Duration(conf.Viper.GetInt("mysql.conn_max_lifetime")) * time.Second)
	sqlDB.SetMaxIdleConns(conf.Viper.GetInt("mysql.max_idel_conns"))
	sqlDB.SetMaxOpenConns(conf.Viper.GetInt("mysql.max_open_conns"))

	MysqlClient = dbClient
	log.InfoMap(context.Background(), consts.DLTagInitMysqlSuccess, logMap)
}
