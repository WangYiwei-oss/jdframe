package jdft

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

var onceCron sync.Once
var taskCron *cron.Cron

func getCronTask() *cron.Cron{	//创建定时任务的单例模式
	onceCron.Do(func(){
		taskCron=cron.New(cron.WithSeconds())
	})
	return taskCron
}

type Jdft struct {
	*gin.Engine
	Router *gin.RouterGroup
	beanfactory *BeanFactory
}

func NewJdft() *Jdft {
	getCronTask().Start()
	return &Jdft{Engine: gin.New(),beanfactory: NewBeanFactory()}
}

func (j *Jdft)Handle(httpMethod, relativePath string,handler interface{})*Jdft{
	if handlerfunc := ConvertResponder(handler);handlerfunc!=nil{
		j.Router.Handle(httpMethod,relativePath,handlerfunc)
	}
	return j
}

func (j *Jdft)Launch(){
	ip,_ := GlobalSettings["IP"].(string)
	err := j.Run(ip)
	if err!=nil{
		log.Fatalln("服务器启动失败",err)
	}
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

//增加定时任务
func (j *Jdft)CronTask(expr string,f func())*Jdft{
	_,err := getCronTask().AddFunc(expr,f)
	if err!=nil{
		log.Fatalln("[error] 启动定时任务失败：",err)
	}
	return j
}