package jdft

import "github.com/gin-gonic/gin"

var StatusCodeMap map[int]string
var ResponderList []Responder

func init(){
	ResponderList = []Responder{
		new(StringResponder),
	}
	StatusCodeMap = map[int]string{
		1:"成功",
		-400: "服务器错误",
	}
}

type Responder interface {
	RespondTo()gin.HandlerFunc
}

type StringResponder func(ctx *gin.Context)(int,string)

func (s StringResponder)RespondTo()gin.HandlerFunc{
	return func(context *gin.Context) {
		status,data := s(context)
		if status > 0{
			context.JSON(200,gin.H{
				"success":true,
				"status":status,
				"data":data,
			})
		}
	}
}
//传过来应该是一个 func(ctx *gin.Context)(int,string)这样的类型
func ConvertResponder(handler interface{})gin.HandlerFunc{
	return nil
}