package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func PrintAndDie(err error) {
	fmt.Printf("init failed,err=%+v", err)
	os.Exit(-1)
}

func MD5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
