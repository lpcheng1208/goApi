package utils

import (
	"goApi/global"
	"os"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    checkFileIsExist
// @description   查看文件是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    bool             bool
// 判断文件是否存在  存在返回 true 不存在返回false

func CheckFileIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// @title    createDir
// @description   批量创建文件夹
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.OPLOGGER.Debug("create directory ", v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.OPLOGGER.Error("create directory", v, " error:", err)
			}
		}
	}
	return err
}
