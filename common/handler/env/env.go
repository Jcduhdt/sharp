package env

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sharp/common/util"
	"strings"
)

const envFile = ".deploy/service.env.txt"

var (
	cluster = map[string]string{
		"dev":  "development",
		"prod": "production",
	}

	ENV = ""
)

func Init() {
	file, err := os.Open(envFile)
	if err != nil {
		util.PrintAndDie(errors.New("environment file not exit, err = " + err.Error()))
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		util.PrintAndDie(errors.New("read environment file failed, err = " + err.Error()))
	}
	env := strings.TrimSpace(string(fileContent))

	if _, ok := cluster[env]; !ok {
		util.PrintAndDie(errors.New("invalid environment, info: " + env + err.Error()))
	}

	ENV = env
	fmt.Printf("environment is %+v\n", cluster[ENV])
}

func GetEnv() string {
	if _, ok := cluster[ENV]; !ok {
		fmt.Printf("environment invaild! env = %s\n", ENV)
	}
	return cluster[ENV]
}
