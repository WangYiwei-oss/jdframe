package jdft

import (
	"reflect"
)

type BeanFactory struct {
	Beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	return &BeanFactory{}
}

func (b *BeanFactory) addBean(bean interface{}) {
	b.Beans = append(b.Beans, bean)
	b.inject(bean) //bean自己也可以注入
}

func (b *BeanFactory) lookupBean(p reflect.Type) interface{} {
	for _, bean := range b.Beans {
		if p == reflect.TypeOf(bean) {
			return bean
		}
	}
	return nil
}

func (b *BeanFactory) inject(obj interface{}) {
	v_obj := reflect.ValueOf(obj).Elem()
	t_obj := v_obj.Type()
	for i := 0; i < v_obj.NumField(); i++ {
		objfield := v_obj.Field(i)
		if objfield.Kind() != reflect.Ptr || !objfield.IsNil() { //字段必须是空的指针
			continue
		}
		inject_value := t_obj.Field(i).Tag.Get("inject")
		if inject_value == "" {
			continue
		}
		if inject_value == "-" {
			if bean := b.lookupBean(objfield.Type()); bean != nil {
				//objfield.Set(reflect.ValueOf(bean))
				objfield.Set(reflect.New(objfield.Type().Elem())) //指针的指针设置New自己
				objfield.Elem().Set(reflect.ValueOf(bean).Elem()) //而值直接设为豌豆荚中的对应元素
			}
		}
	}
}
