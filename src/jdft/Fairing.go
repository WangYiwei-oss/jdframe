package jdft

import "github.com/gin-gonic/gin"

//用于规范中间件代码的接口。作为第二种加入中间件的方式

type Fairing interface {
	OnRequest(ctx *gin.Context) error                   //执行控制器方法前，修改如头信息，判断参数等等
	OnResponse(result interface{}) (interface{}, error) //执行控制器方法后，可以修改返回值内容
}
