package jdft

import (
	mcasbin "github.com/WangYiwei-oss/jdframe/src/casbin"
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"github.com/WangYiwei-oss/jdframe/src/models"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

//暴露一些类型给用户
type User models.User
type UserRole models.UserRole

var CasbinEnforcer *casbin.Enforcer

var mlog *Mlogger
var onceCron sync.Once
var taskCron *cron.Cron

func init() {
	mlog = GetLogger("Global")
	CasbinEnforcer = mcasbin.E
}

func GetGlobalSettings() map[string]interface{} {
	return configparser.GlobalSettings
}

func getCronTask() *cron.Cron { //创建定时任务的单例模式
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

type Jdft struct {
	*gin.Engine
	Router      *gin.RouterGroup
	beanfactory *BeanFactory
}

func NewJdft() *Jdft {
	getCronTask().Start()
	return &Jdft{Engine: gin.New(), beanfactory: NewBeanFactory()}
}

func (j *Jdft) Handle(httpMethod, relativePath string, handler interface{}) *Jdft {
	if handlerfunc := ConvertResponder(handler); handlerfunc != nil {
		router := j.Router
		router.Handle(httpMethod, relativePath, handlerfunc)
	} else {
		log.Printf("路由%s, %s没有被解析并注入", httpMethod, relativePath)
	}
	return j
}

func (j *Jdft) Attach(f ...gin.HandlerFunc) *Jdft {
	j.Use(f...)
	return j
}

func (j *Jdft) Launch() {
	ip, _ := configparser.GlobalSettings["IP"].(string)
	err := j.Run(ip)
	if err != nil {
		log.Fatalf("服务器启动失败, %s", err)
	}
}

func (j *Jdft) Mount(version string, controllers ...JdController) *Jdft {
	j.Router = j.Group(version)
	for _, controller := range controllers {
		controller.Build(j)
		j.beanfactory.inject(controller)
	}
	return j
}

func (j *Jdft) Beans(beans ...interface{}) *Jdft {
	for _, bean := range beans {
		j.beanfactory.addBean(bean)
	}
	return j
}

//增加定时任务
func (j *Jdft) CronTask(expr string, f func()) *Jdft {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Fatalf("启动定时任务失败, %s", err)
	}
	return j
}
