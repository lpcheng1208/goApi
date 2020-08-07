package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goApi/global"
	"io/ioutil"
)

/**
api访问日志中间件
*/

func ParseParamMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := global.LoggerWithContext(ctx)

		var mapResult map[string]interface{}

		// 获取非 GET 的参数
		if ctx.Request.Method != "GET" {
			// 直接获取 body 内容
			dataJson, _ := ctx.GetRawData()
			// 取出来之后需要重新赋值回去，不然处理函数取不到值
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(dataJson))

			_ = json.Unmarshal(dataJson, &mapResult)
		}

		// 获取 Get 的参数
		Params := ctx.Request.URL.Query()

		if mapResult == nil {
			mapResult = make(map[string]interface{})
		}

		for k, v := range Params {
			value := v[0]
			_, ok := mapResult[k]
			// 如果post参数同名，以最终post 上传的 参数为准， 不进行覆盖
			if !ok {
				mapResult[k] = value
			}
		}

		ctx.Set("Param", mapResult)
		log.Info("获取到的参数", zap.Any("request path", ctx.Request.URL.Path), zap.Any("Request Param", mapResult))
		ctx.Next()
	}

}
