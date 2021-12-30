package configinjector

import (
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"log"
	"reflect"
)

type presetBeans map[string]interface{}

var ConfigInjector presetBeans

func RegisterBean(name string, bean interface{}) {
	ConfigInjector[name] = bean
}

func _findBean(name string, beanNames []interface{}) bool {
	for _, beanName := range beanNames {
		if name == beanName.(string) {
			return true
		}
	}
	return false
}

func GetBeans() []interface{} {
	ret := make([]interface{}, 0)
	beanNames, _ := configparser.GlobalSettings["BEANS"].([]interface{})
	for presetName, presetFunc := range ConfigInjector {
		if _findBean(presetName, beanNames) {
			presetFuncV := reflect.ValueOf(presetFunc)
			if presetFuncV.Kind() == reflect.Func {
				parms := make([]reflect.Value, 0)
				ret = append(ret, presetFuncV.Call(parms))
			} else {
				log.Fatalln("ConfigInjector: 必须传一个func对象")
			}
		}
	}
	return ret
}

func init() {
	ConfigInjector = make(map[string]interface{})
}
