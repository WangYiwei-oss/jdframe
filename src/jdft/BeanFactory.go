package jdft

import (
	"log"
	"reflect"
)

type BeanFactory struct {
	Beans []interface{}
}

func NewBeanFactory() *BeanFactory {
	return &BeanFactory{}
}

// addBean 将bean注册到bean factory
func (b *BeanFactory) addBean(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr { //必须是指针类型
		log.Fatalln("bean_factory required ptr object")
	}
	if t.Elem().Kind() != reflect.Struct { //如果不是一个结构体指针的话不予注册
		return
	}
	b.Beans = append(b.Beans, bean) //把bean注册到注册表中去
	b.inject(bean)                  //
}

func (b *BeanFactory) lookupBean(p reflect.Type) interface{} {
	for _, bean := range b.Beans {
		if p == reflect.TypeOf(bean) {
			return bean
		}
	}
	return nil
}

// inject 将结构体的带有inject:"-"标签的，进行依赖注入
func (b *BeanFactory) inject(obj interface{}) {
	if obj == nil {
		return
	}
	v_obj := reflect.ValueOf(obj)
	if v_obj.Kind() == reflect.Ptr { //确保传入对象而不是指针
		v_obj = v_obj.Elem()
	}
	for i := 0; i < v_obj.NumField(); i++ { //对结构体每一个成员进行遍历
		objField := v_obj.Type().Field(i)
		if v_obj.Field(i).Kind() != reflect.Ptr || !v_obj.Field(i).IsNil() { //字段必须是空的指针
			continue
		}
		if v_obj.Field(i).CanSet() && objField.Tag.Get("inject") == "-" { //必须带inject="-"注解
			if bean := b.lookupBean(objField.Type); bean != nil { //去注册表中寻找，如果找到了就进行依赖注入
				v_obj.Field(i).Set(reflect.ValueOf(bean))
				b.inject(bean) //处理循环注入
			}
		}
	}
}
