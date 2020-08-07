package utils

import (
	"fmt"
	"goApi/config"
	"strconv"
	"strings"
	"unsafe"
)

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToInt64(parseString string) (result int64, err error) {
	result, err = strconv.ParseInt(parseString, 10, 64)
	return
}

func StringToInt(parseString string) (result int, err error) {
	result, err = strconv.Atoi(parseString)
	return
}

func IntToString(intNum int) (result string) {
	result = strconv.Itoa(intNum)
	return
}

func Int64ToString(intNum int64) (result string) {
	result = strconv.FormatInt(intNum, 10)
	return
}

func GetCdnUrl(originUrl string) (cdnUrl string) {
	b := strings.HasPrefix(originUrl, "http")
	if b {
		cdnUrl = originUrl
		return
	} else {
		return config.CdnPrefix + strings.Trim(originUrl, "/")
	}
}

func GetMessageIndex(sendFrom, sendTo string) (msgIndex string) {
	if sendFrom > sendTo {
		msgIndex = fmt.Sprintf("%s#%s", sendFrom, sendTo)
	} else {
		msgIndex = fmt.Sprintf("%s#%s", sendTo, sendFrom)
	}
	return
}
