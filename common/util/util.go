package util

import (
	"fmt"
	"os"
)

func PrintAndDie(err error)  {
	fmt.Printf("init failed,err=%+v",err)
	os.Exit(-1)
}