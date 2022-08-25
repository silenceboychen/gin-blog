package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	//将[]byte转成16进制
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
