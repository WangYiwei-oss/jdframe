package jdft

import (
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"reflect"
	"strconv"
)

var StatusCodeMap map[int]string
var ResponderList []Responder
var CanntFindStatusJSON gin.H

func init() {
	ResponderList = []Responder{
		new(SingleResponder),
		new(StringResponder),
		new(JsonResponder),
		new(IntResponder),
		new(Float32Responder),
		new(Float64Responder),
	}
	StatusCodeMap = make(map[int]string)
	for k, v := range configparser.GlobalSettings["STATUS_CODE"].(map[string]interface{}) {
		ik, _ := strconv.Atoi(k)
		StatusCodeMap[ik] = v.(string)
	}

	CanntFindStatusJSON = gin.H{
		"success": false,
		"status":  "Server error, Cannt find StatusCode",
		"data":    "",
	}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

type SingleResponder func(ctx *gin.Context) int

type StringResponder func(ctx *gin.Context) (int, string)

type IntResponder func(ctx *gin.Context) (int, int)

type Float32Responder func(ctx *gin.Context) (int, float32)

type Float64Responder func(ctx *gin.Context) (int, float64)

type Json interface{}
type JsonResponder func(ctx *gin.Context) (int, Json)

func (s SingleResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := s(context)
		retcode := 200
		success := true
		if code < 0 {
			retcode = 200
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			context.JSON(retcode, gin.H{
				"success": success,
				"status":  stat,
				"data":    "",
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (s StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := s(context)
		success := true
		if code < 0 {
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			context.JSON(200, gin.H{
				"success": success,
				"status":  stat,
				"data":    data,
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (i IntResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := i(context)
		success := true
		if code < 0 {
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			context.JSON(200, gin.H{
				"success": success,
				"status":  stat,
				"data":    data,
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (f Float32Responder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := f(context)
		success := true
		if code < 0 {
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			context.JSON(200, gin.H{
				"success": success,
				"status":  stat,
				"data":    data,
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (f Float64Responder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := f(context)
		success := true
		if code < 0 {
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			context.JSON(200, gin.H{
				"success": success,
				"status":  stat,
				"data":    data,
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (j JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := j(context)
		retcode := 200
		success := true
		if code < 0 {
			retcode = 200
			success = false
		}
		if stat, ok := StatusCodeMap[code]; ok {
			ret, err := json.Marshal(data)
			if err != nil {
				panic("解析model错误")
			}
			context.JSON(retcode, gin.H{
				"success": success,
				"status":  stat,
				"data":    string(ret),
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

//TODO 关键函数
//传过来应该是一个 func(ctx *gin.Context)(int,string)这样的类型
func ConvertResponder(handler interface{}) gin.HandlerFunc {
	v_handler := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		v_r := reflect.ValueOf(r).Elem()
		if v_handler.Type().ConvertibleTo(v_r.Type()) {
			v_r.Set(v_handler)
			return v_r.Interface().(Responder).RespondTo() //TODO
		}
	}
	return nil
}
