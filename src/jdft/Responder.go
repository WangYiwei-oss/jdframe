package jdft

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var StatusCodeMap map[int]string
var ResponderList []Responder
var CanntFindStatusJSON gin.H
func init(){
	ResponderList = []Responder{
		new(StringResponder),
	}
	StatusCodeMap = map[int]string{
		1:"成功",
		-400: "服务器错误",
	}
	CanntFindStatusJSON = gin.H{
		"success": false,
		"status":  "Server error, Cannt find StatusCode",
		"data":    "",
	}
}

type Responder interface {
	RespondTo()gin.HandlerFunc
}

type StringResponder func(ctx *gin.Context)(int,string)

func (s StringResponder)RespondTo()gin.HandlerFunc{
	return func(context *gin.Context) {
		code,data := s(context)
		if code > 0{
			if stat,ok:=StatusCodeMap[code];ok {
				context.JSON(200, gin.H{
					"success": true,
					"status":  stat,
					"data":    data,
				})
				return
			}
		}else{
			if stat,ok:=StatusCodeMap[code];ok{
				context.JSON(400,gin.H{
					"success":false,
					"status":stat,
					"data":data,
				})
				return
			}
		}
		context.JSON(500,CanntFindStatusJSON)
	}
}



//传过来应该是一个 func(ctx *gin.Context)(int,string)这样的类型
func ConvertResponder(handler interface{})(gin.HandlerFunc){
	v_handler := reflect.ValueOf(handler)
	for _,r := range ResponderList{
		v_r := reflect.ValueOf(r).Elem()
		if v_handler.Type().ConvertibleTo(v_r.Type()){
			v_r.Set(v_handler)
			return v_r.Interface().(Responder).RespondTo()	//TODO
		}
	}
	return nil
}