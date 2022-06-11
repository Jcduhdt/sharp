package conf

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"

	"sharp/common/handler/env"
)

var (
	Viper = viper.New()
)

const (
	ConfigName     = "app"
	IncludeConfigs = "include.configs"
)

func InitConf(confPath string) {
	path := confPath + env.GetEnv() + "/"
	Viper.AddConfigPath(path)

	Viper.SetConfigName(ConfigName)
	err := Viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("fatal error config file: %v", err)
	}

	files := Viper.GetStringSlice(IncludeConfigs)
	for _, file := range files {
		Viper.SetConfigName(file)
		err = Viper.MergeInConfig()
		if err != nil {
			fmt.Errorf("fatal error config file: %+v,err=%+v", file, err)
		}
	}
}

func FindConfigDir() string {
	_, thisPath, _, _ := runtime.Caller(0)
	testDir := path.Dir(thisPath)
	handlerDir := path.Dir(testDir)
	commonDir := path.Dir(handlerDir)
	projectDir := path.Dir(commonDir)
	return projectDir + "/conf/"
}
