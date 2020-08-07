package initialize

import (
	"github.com/zheng-ji/goSnowFlake"
	"goApi/global"
	"os"
)

func InitIw(workId int64) {
	iw, err := goSnowFlake.NewIdWorker(workId)
	if err != nil {
		os.Exit(0)
	}
	global.UniqId = iw

}
