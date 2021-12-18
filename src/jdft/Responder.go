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
		new(ModelResponder),
		new(SliceResponder),
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

type ModelResponder func(ctx *gin.Context) (int, JModel)

type SliceResponder func(ctx *gin.Context) (int, []interface{})

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
				"data":    data,
			})
			return
		}
		context.JSON(500, CanntFindStatusJSON)
	}
}

func (s SliceResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := s(context)
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

func (m ModelResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		code, data := m(context)
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
