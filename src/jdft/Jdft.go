package jdft

import (
	"github.com/gin-gonic/gin"
)

type Jdft struct {
	*gin.Engine
	Router *gin.RouterGroup
	beanfactory *BeanFactory
}

func NewJdft() *Jdft {
	return &Jdft{Engine: gin.New(),beanfactory: NewBeanFactory()}
}

func (j *Jdft)Handle(httpMethod, relativePath string,handler interface{})*Jdft{
	if handlerfunc := ConvertResponder(handler);handlerfunc!=nil{
		j.Router.Handle(httpMethod,relativePath,handlerfunc)
	}
	return j
}

func (j *Jdft)Launch(){
	j.Run(":8081")
}

func (j *Jdft)Mount(version string,controllers ...JdController)*Jdft{
	j.Router = j.Group(version)
	for _,controller := range controllers{
		controller.Build(j)
		j.beanfactory.inject(controller)
	}
	return j
}

func (j *Jdft)Beans(beans ...interface{})*Jdft{
	for _,bean := range beans{
		j.beanfactory.addBean(bean)
	}
	return j
}
